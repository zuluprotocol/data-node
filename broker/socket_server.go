package broker

import (
	"context"
	"fmt"

	"code.vegaprotocol.io/data-node/events"
	"code.vegaprotocol.io/data-node/logging"
	eventspb "code.vegaprotocol.io/protos/vega/events/v1"
	"golang.org/x/sync/errgroup"

	"github.com/golang/protobuf/proto"
	mangos "go.nanomsg.org/mangos/v3"
	mangosErr "go.nanomsg.org/mangos/v3/errors"
	"go.nanomsg.org/mangos/v3/protocol"
	"go.nanomsg.org/mangos/v3/protocol/pull"
	_ "go.nanomsg.org/mangos/v3/transport/tcp"
)

const defaultEventChannelBufferSize = 256

// SocketServer receives events from a remote broker.
// This is used by the data node to receive events from a non-validating core node.
type SocketServer struct {
	log    *logging.Logger
	config *SocketConfig

	sock protocol.Socket
}

func pipeEventToString(pe mangos.PipeEvent) string {
	switch pe {
	case 1:
		return "Attached"
	case 2:
		return "Detached"
	default:
		return "Attaching"
	}
}

func NewSocketServer(log *logging.Logger, config *SocketConfig) (*SocketServer, error) {
	sock, err := pull.NewSocket()
	if err != nil {
		return nil, fmt.Errorf("failed to create new socket: %w", err)
	}

	return &SocketServer{
		log:    log,
		config: config,
		sock:   sock,
	}, nil
}

func (s SocketServer) Listen() error {
	addr := fmt.Sprintf("tcp://%s:%d", s.config.IP, s.config.Port)
	if err := s.sock.Listen(addr); err != nil {
		return fmt.Errorf("failed to listen on %v: %w", addr, err)
	}

	s.log.Info("Starting broker socket server", logging.String("addr", s.config.IP), logging.Int("port", s.config.Port))

	s.sock.SetPipeEventHook(func(pe mangos.PipeEvent, p mangos.Pipe) {
		s.log.Info(
			"New broker connection event",
			logging.String("eventType", pipeEventToString(pe)),
			logging.Uint32("id", p.ID()),
			logging.String("address", p.Address()),
		)
	})

	return nil
}

func (s SocketServer) Receive(ctx context.Context) (<-chan events.Event, <-chan error) {
	receiveCh := make(chan events.Event, defaultEventChannelBufferSize)
	errCh := make(chan error, 1)

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		<-ctx.Done()
		if err := s.close(); err != nil {
			return fmt.Errorf("failed to close socket: %w", err)
		}
		return nil
	})

	eg.Go(func() error {
		defer close(receiveCh)

		var recvTimeouts int

		for {
			var be eventspb.BusEvent

			msg, err := s.sock.Recv()
			if err != nil {
				switch err {
				case mangosErr.ErrRecvTimeout:
					s.log.Warn("Receive socket timeout", logging.Error(err))
					recvTimeouts++
					if recvTimeouts > 3 {
						return fmt.Errorf("more then a 3 socket timeouts occurred: %w", err)
					}
				case mangosErr.ErrBadVersion:
				case mangosErr.ErrClosed:
					return fmt.Errorf("fatal error socket: %w", err)
				default:
					s.log.Error("Failed to receive message", logging.Error(err))
					continue
				}
			}

			if err := proto.Unmarshal(msg, &be); err != nil {
				s.log.Error("Failed to unmarshal received event", logging.Error(err))
				continue
			}

			evt := toEvent(ctx, &be)
			receiveCh <- evt
		}
	})

	go func() {
		defer close(errCh)
		if err := eg.Wait(); err != nil {
			errCh <- err
		}
	}()

	return receiveCh, errCh
}

func (s SocketServer) close() error {
	s.log.Info("Closing socket server")
	return s.sock.Close()
}
