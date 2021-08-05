package node

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"code.vegaprotocol.io/data-node/accounts"
	"code.vegaprotocol.io/data-node/api"
	"code.vegaprotocol.io/data-node/assets"
	"code.vegaprotocol.io/data-node/broker"
	"code.vegaprotocol.io/data-node/candles"
	"code.vegaprotocol.io/data-node/config"
	"code.vegaprotocol.io/data-node/fee"
	"code.vegaprotocol.io/data-node/gateway/server"
	"code.vegaprotocol.io/data-node/governance"
	"code.vegaprotocol.io/data-node/liquidity"
	"code.vegaprotocol.io/data-node/logging"
	"code.vegaprotocol.io/data-node/markets"
	"code.vegaprotocol.io/data-node/metrics"
	"code.vegaprotocol.io/data-node/netparams"
	"code.vegaprotocol.io/data-node/notary"
	"code.vegaprotocol.io/data-node/oracles"
	"code.vegaprotocol.io/data-node/orders"
	"code.vegaprotocol.io/data-node/parties"
	"code.vegaprotocol.io/data-node/plugins"
	"code.vegaprotocol.io/data-node/pprof"
	"code.vegaprotocol.io/data-node/risk"
	"code.vegaprotocol.io/data-node/stats"
	"code.vegaprotocol.io/data-node/storage"
	"code.vegaprotocol.io/data-node/subscribers"
	"code.vegaprotocol.io/data-node/trades"
	"code.vegaprotocol.io/data-node/transfers"
	"code.vegaprotocol.io/data-node/vegatime"
	types "code.vegaprotocol.io/protos/vega"
	vegaprotoapi "code.vegaprotocol.io/protos/vega/api"
	"golang.org/x/sync/errgroup"
)

type AccountStore interface {
	accounts.AccountStore
	SaveBatch([]*types.Account) error
	Close() error
	ReloadConf(storage.Config)
}

type CandleStore interface {
	FetchLastCandle(marketID string, interval types.Interval) (*types.Candle, error)
	GenerateCandlesFromBuffer(marketID string, previousCandlesBuf map[string]types.Candle) error
	candles.CandleStore
	Close() error
	ReloadConf(storage.Config)
}

type OrderStore interface {
	orders.OrderStore
	SaveBatch([]types.Order) error
	Close() error
	ReloadConf(storage.Config)
}

type TradeStore interface {
	trades.TradeStore
	SaveBatch([]types.Trade) error
	Close() error
	ReloadConf(storage.Config)
}

// NodeCommand use to implement 'node' command.
type NodeCommand struct {
	ctx    context.Context
	cancel context.CancelFunc

	accounts              AccountStore
	candleStore           CandleStore
	orderStore            OrderStore
	marketStore           *storage.Market
	marketDataStore       *storage.MarketData
	tradeStore            TradeStore
	partyStore            *storage.Party
	riskStore             *storage.Risk
	transferResponseStore *storage.TransferResponse

	vegaTradingServiceClient vegaprotoapi.TradingServiceClient

	broker *broker.Broker

	transferSub      *subscribers.TransferResponse
	marketEventSub   *subscribers.MarketEvent
	orderSub         *subscribers.OrderEvent
	accountSub       *subscribers.AccountSub
	partySub         *subscribers.PartySub
	tradeSub         *subscribers.TradeSub
	marginLevelSub   *subscribers.MarginLevelSub
	governanceSub    *subscribers.GovernanceDataSub
	voteSub          *subscribers.VoteSub
	marketDataSub    *subscribers.MarketDataSub
	newMarketSub     *subscribers.Market
	marketUpdatedSub *subscribers.MarketUpdated
	candleSub        *subscribers.CandleSub
	riskFactorSub    *subscribers.RiskFactorSub
	marketDepthSub   *subscribers.MarketDepthBuilder

	candleService     *candles.Svc
	tradeService      *trades.Svc
	marketService     *markets.Svc
	orderService      *orders.Svc
	liquidityService  *liquidity.Svc
	partyService      *parties.Svc
	timeService       *vegatime.Svc
	accountsService   *accounts.Svc
	transfersService  *transfers.Svc
	riskService       *risk.Svc
	governanceService *governance.Svc
	notaryService     *notary.Svc
	assetService      *assets.Svc
	feeService        *fee.Svc
	eventService      *subscribers.Service
	netParamsService  *netparams.Service
	oracleService     *oracles.Service

	pproffhandlr *pprof.Pprofhandler
	configPath   string
	conf         config.Config
	stats        *stats.Stats
	Log          *logging.Logger
	cfgwatchr    *config.Watcher

	// plugins
	settlePlugin     *plugins.Positions
	notaryPlugin     *plugins.Notary
	assetPlugin      *plugins.Asset
	withdrawalPlugin *plugins.Withdrawal
	depositPlugin    *plugins.Deposit

	Version     string
	VersionHash string
}

func (l *NodeCommand) Run(cfgwatchr *config.Watcher, rootPath string, args []string) error {
	l.cfgwatchr = cfgwatchr

	l.conf, l.configPath = cfgwatchr.Get(), rootPath

	stages := []func([]string) error{
		l.persistentPre,
		l.preRun,
		l.runNode,
		l.postRun,
		l.persistentPost,
	}
	for _, fn := range stages {
		if err := fn(args); err != nil {
			return err
		}
	}

	return nil
}

// runNode is the entry of node command.
func (l *NodeCommand) runNode(args []string) error {
	defer l.cancel()

	// gRPC server
	grpcServer := api.NewGRPCServer(
		l.Log,
		l.conf.API,
		l.stats,
		l.vegaTradingServiceClient,
		l.timeService,
		l.marketService,
		l.partyService,
		l.orderService,
		l.liquidityService,
		l.tradeService,
		l.candleService,
		l.accountsService,
		l.transfersService,
		l.riskService,
		l.governanceService,
		l.notaryService,
		l.assetService,
		l.feeService,
		l.eventService,
		l.oracleService,
		l.withdrawalPlugin,
		l.depositPlugin,
		l.marketDepthSub,
		l.netParamsService,
	)

	// watch configs
	l.cfgwatchr.OnConfigUpdate(
		func(cfg config.Config) { grpcServer.ReloadConf(cfg.API) },
	)

	eg, ctx := errgroup.WithContext(l.ctx)
	ctx, cancel := context.WithCancel(ctx)

	// start the grpc server
	eg.Go(func() error { return grpcServer.Start(ctx) })

	// start gateway
	if l.conf.GatewayEnabled {
		gty := server.New(l.conf.Gateway, l.Log)

		eg.Go(func() error { return gty.Start(ctx) })
	}

	eg.Go(func() error {
		return l.broker.Receive(ctx)
	})

	// waitSig will wait for a sigterm or sigint interrupt.
	eg.Go(func() error {
		var gracefulStop = make(chan os.Signal, 1)
		signal.Notify(gracefulStop, syscall.SIGTERM)
		signal.Notify(gracefulStop, syscall.SIGINT)

		select {
		case sig := <-gracefulStop:
			l.Log.Info("Caught signal", logging.String("name", fmt.Sprintf("%+v", sig)))
			cancel()
		case <-ctx.Done():
			return ctx.Err()
		}

		return nil
	})

	metrics.Start(l.conf.Metrics)

	l.Log.Info("Vega startup complete")

	return eg.Wait()
}

func flagProvided(flag string) bool {
	for _, v := range os.Args[1:] {
		if v == flag {
			return true
		}
	}

	return false
}
