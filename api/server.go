package api

import (
	"context"
	"net"
	"strconv"
	"time"

	"code.vegaprotocol.io/data-node/accounts"
	"code.vegaprotocol.io/data-node/assets"
	"code.vegaprotocol.io/data-node/candles"
	"code.vegaprotocol.io/data-node/checkpoint"
	"code.vegaprotocol.io/data-node/contextutil"
	"code.vegaprotocol.io/data-node/delegations"
	"code.vegaprotocol.io/data-node/epochs"
	"code.vegaprotocol.io/data-node/fee"
	"code.vegaprotocol.io/data-node/governance"
	"code.vegaprotocol.io/data-node/liquidity"
	"code.vegaprotocol.io/data-node/logging"
	"code.vegaprotocol.io/data-node/netparams"
	"code.vegaprotocol.io/data-node/nodes"
	"code.vegaprotocol.io/data-node/notary"
	"code.vegaprotocol.io/data-node/oracles"
	"code.vegaprotocol.io/data-node/orders"
	"code.vegaprotocol.io/data-node/parties"
	"code.vegaprotocol.io/data-node/plugins"
	"code.vegaprotocol.io/data-node/risk"
	"code.vegaprotocol.io/data-node/staking"
	"code.vegaprotocol.io/data-node/stats"
	"code.vegaprotocol.io/data-node/subscribers"
	"code.vegaprotocol.io/data-node/trades"
	"code.vegaprotocol.io/data-node/transfers"
	"code.vegaprotocol.io/data-node/vegatime"
	"golang.org/x/sync/errgroup"

	protoapi "code.vegaprotocol.io/protos/data-node/api/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

// GRPCServer represent the grpc api provided by the vega node
type GRPCServer struct {
	Config

	log                      *logging.Logger
	srv                      *grpc.Server
	stats                    *stats.Stats
	vegaTradingServiceClient TradingServiceClient

	accountsService         *accounts.Svc
	candleService           *candles.Svc
	marketService           MarketService
	orderService            *orders.Svc
	liquidityService        *liquidity.Svc
	partyService            *parties.Svc
	timeService             *vegatime.Svc
	tradeService            *trades.Svc
	transferResponseService *transfers.Svc
	riskService             *risk.Svc
	governanceService       *governance.Svc
	notaryService           *notary.Svc
	assetService            *assets.Svc
	feeService              *fee.Svc
	eventService            *subscribers.Service
	withdrawalService       *plugins.Withdrawal
	depositService          *plugins.Deposit
	netParamsService        *netparams.Service
	oracleService           *oracles.Service
	stakingService          *staking.Service
	tradingProxySvc         *tradingProxyService
	tradingDataService      *tradingDataService
	nodeService             *nodes.Service
	epochService            *epochs.Service
	delegationService       *delegations.Service
	rewardsService          *subscribers.RewardCounters
	checkpointSvc           *checkpoint.Svc

	marketDepthService *subscribers.MarketDepthBuilder

	// used in order to gracefully close streams
	ctx   context.Context
	cfunc context.CancelFunc
}

// NewGRPCServer create a new instance of the GPRC api for the vega node
func NewGRPCServer(
	log *logging.Logger,
	config Config,
	stats *stats.Stats,
	tradingServiceClient TradingServiceClient,
	timeService *vegatime.Svc,
	marketService MarketService,
	partyService *parties.Svc,
	orderService *orders.Svc,
	liquidityService *liquidity.Svc,
	tradeService *trades.Svc,
	candleService *candles.Svc,
	accountsService *accounts.Svc,
	transferResponseService *transfers.Svc,
	riskService *risk.Svc,
	governanceService *governance.Svc,
	notaryService *notary.Svc,
	assetService *assets.Svc,
	feeService *fee.Svc,
	eventService *subscribers.Service,
	oracleService *oracles.Service,
	withdrawalService *plugins.Withdrawal,
	depositService *plugins.Deposit,
	marketDepthService *subscribers.MarketDepthBuilder,
	netParamsService *netparams.Service,
	nodeService *nodes.Service,
	epochService *epochs.Service,
	delegationService *delegations.Service,
	rewardsService *subscribers.RewardCounters,
	stakingService *staking.Service,
	checkpointSvc *checkpoint.Svc,
) *GRPCServer {
	// setup logger
	log = log.Named(namedLogger)
	log.SetLevel(config.Level.Get())
	ctx, cfunc := context.WithCancel(context.Background())

	return &GRPCServer{
		log:                      log,
		Config:                   config,
		stats:                    stats,
		vegaTradingServiceClient: tradingServiceClient,
		orderService:             orderService,
		liquidityService:         liquidityService,
		tradeService:             tradeService,
		candleService:            candleService,
		timeService:              timeService,
		marketService:            marketService,
		partyService:             partyService,
		accountsService:          accountsService,
		transferResponseService:  transferResponseService,
		riskService:              riskService,
		governanceService:        governanceService,
		notaryService:            notaryService,
		assetService:             assetService,
		feeService:               feeService,
		eventService:             eventService,
		withdrawalService:        withdrawalService,
		depositService:           depositService,
		marketDepthService:       marketDepthService,
		netParamsService:         netParamsService,
		oracleService:            oracleService,
		nodeService:              nodeService,
		epochService:             epochService,
		delegationService:        delegationService,
		rewardsService:           rewardsService,
		stakingService:           stakingService,
		checkpointSvc:            checkpointSvc,
		ctx:                      ctx,
		cfunc:                    cfunc,
	}
}

// ReloadConf update the internal configuration of the GRPC server
func (g *GRPCServer) ReloadConf(cfg Config) {
	g.log.Info("reloading configuration")
	if g.log.GetLevel() != cfg.Level.Get() {
		g.log.Info("updating log level",
			logging.String("old", g.log.GetLevel().String()),
			logging.String("new", cfg.Level.String()),
		)
		g.log.SetLevel(cfg.Level.Get())
	}

	// TODO(): not updating the the actual server for now, may need to look at this later
	// e.g restart the http server on another port or whatever
	g.Config = cfg
}

func remoteAddrInterceptor(log *logging.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		// first check if the request is forwarded from our restproxy
		// get the metadata
		var ip string
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			forwardedFor, ok := md["x-forwarded-for"]
			if ok && len(forwardedFor) > 0 {
				log.Debug("grpc request x-forwarded-for",
					logging.String("method", info.FullMethod),
					logging.String("remote-ip-addr", forwardedFor[0]),
				)
				ip = forwardedFor[0]
			}
		}

		// if the request is not forwarded let's get it from the peer infos
		if len(ip) <= 0 {
			p, ok := peer.FromContext(ctx)
			if ok && p != nil {
				log.Debug("grpc peer client request",
					logging.String("method", info.FullMethod),
					logging.String("remote-ip-addr", p.Addr.String()))
				ip = p.Addr.String()
			}
		}

		ctx = contextutil.WithRemoteIPAddr(ctx, ip)

		// Calls the handler
		h, err := handler(ctx, req)

		log.Debug("Invoked RPC call",
			logging.String("method", info.FullMethod),
			logging.Error(err),
		)

		return h, err
	}
}

func (g *GRPCServer) getTCPListener() (net.Listener, error) {
	ip := g.IP
	port := strconv.Itoa(g.Port)

	g.log.Info("Starting gRPC based API", logging.String("addr", ip), logging.String("port", port))

	tpcLis, err := net.Listen("tcp", net.JoinHostPort(ip, port))
	if err != nil {
		return nil, err
	}

	return tpcLis, nil
}

// Start start the grpc server.
// Uses default TCP listener if no provided.
func (g *GRPCServer) Start(ctx context.Context, lis net.Listener) error {
	if lis == nil {
		tpcLis, err := g.getTCPListener()
		if err != nil {
			return err
		}

		lis = tpcLis
	}

	intercept := grpc.UnaryInterceptor(remoteAddrInterceptor(g.log))
	g.srv = grpc.NewServer(intercept)

	tradingProxySvc := &tradingProxyService{
		log:                  g.log,
		conf:                 g.Config,
		tradingServiceClient: g.vegaTradingServiceClient,
	}
	g.tradingProxySvc = tradingProxySvc
	protoapi.RegisterTradingProxyServiceServer(g.srv, tradingProxySvc)

	tradingDataSvc := &tradingDataService{
		log:                     g.log,
		Config:                  g.Config,
		Stats:                   g.stats,
		OrderService:            g.orderService,
		TradeService:            g.tradeService,
		CandleService:           g.candleService,
		MarketService:           g.marketService,
		PartyService:            g.partyService,
		TimeService:             g.timeService,
		AccountsService:         g.accountsService,
		TransferResponseService: g.transferResponseService,
		RiskService:             g.riskService,
		NotaryService:           g.notaryService,
		governanceService:       g.governanceService,
		AssetService:            g.assetService,
		FeeService:              g.feeService,
		eventService:            g.eventService,
		WithdrawalService:       g.withdrawalService,
		DepositService:          g.depositService,
		MarketDepthService:      g.marketDepthService,
		NetParamsService:        g.netParamsService,
		LiquidityService:        g.liquidityService,
		oracleService:           g.oracleService,
		nodeService:             g.nodeService,
		epochService:            g.epochService,
		delegationService:       g.delegationService,
		rewardsService:          g.rewardsService,
		stakingService:          g.stakingService,
		checkpointService:       g.checkpointSvc,
	}
	g.tradingDataService = tradingDataSvc
	protoapi.RegisterTradingDataServiceServer(g.srv, tradingDataSvc)

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		<-ctx.Done()
		g.stop()
		return ctx.Err()
	})

	eg.Go(func() error {
		return g.srv.Serve(lis)
	})

	return eg.Wait()
}

func (g *GRPCServer) stop() {
	if g.srv == nil {
		return
	}

	done := make(chan struct{})
	go func() {
		g.log.Info("Gracefully stopping gRPC based API")
		g.srv.GracefulStop()
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		g.log.Info("Force stopping gRPC based API")
		g.srv.Stop()
	}
}
