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
	pb.UnimplementedRotatorServer
}

func (g *GRPCServer) Start() error {
	l := logger.FromContext(g.ctx)

	g.server = grpc.NewServer(grpc.UnaryInterceptor(g.loggingInterceptor))

	listener, err := net.Listen("tcp", g.port)
	if err != nil {
		return err
	}
	pb.RegisterRotatorServer(g.server, g)

	go func() error {
		err = g.server.Serve(listener)
		if err != nil {
			return err
		}
		return nil
	}()
	l.Debug("grpc server started", zap.String("server port", g.port))
	<-g.ctx.Done()
	return nil
}

func (g *GRPCServer) Stop() {
	g.server.GracefulStop()
}

func NewGRPCServer(ctx context.Context, app rotator.Application, port string) *GRPCServer {
	return &GRPCServer{
		ctx:                        ctx,
		application:                app,
		port:                       ":" + port,
		UnimplementedRotatorServer: pb.UnimplementedRotatorServer{},
	}
}
