//go:generate protoc --proto_path=. --go_out=./pb --go-grpc_out=./pb ./EventService.proto

package internalgrpc

import (
	"context"
	"net"

	"github.com/NoisyPunk/multiarmedbandit/internal/app"
	"github.com/NoisyPunk/multiarmedbandit/internal/logger"
	"github.com/NoisyPunk/multiarmedbandit/internal/server/grpc/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	ctx         context.Context
	application rotator.Application
	server      *grpc.Server
	port        string
	pb.UnimplementedEventsServer
}

func (e *GRPCServer) Start() error {
	l := logger.FromContext(e.ctx)

	e.server = grpc.NewServer(grpc.UnaryInterceptor(e.loggingInterceptor))

	listener, err := net.Listen("tcp", e.port)
	if err != nil {
		return err
	}
	pb.RegisterEventsServer(e.server, e)

	go func() error {
		err = e.server.Serve(listener)
		if err != nil {
			return err
		}
		return nil
	}()
	l.Debug("grpc server started", zap.String("server port", e.port))
	<-e.ctx.Done()
	return nil
}

func (e *GRPCServer) Stop() {
	e.server.GracefulStop()
}

func NewGRPCServer(ctx context.Context, app rotator.Application, port string) *GRPCServer {
	return &GRPCServer{
		ctx:                       ctx,
		application:               app,
		port:                      ":" + port,
		UnimplementedEventsServer: pb.UnimplementedEventsServer{},
	}
}
