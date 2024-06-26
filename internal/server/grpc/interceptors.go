package internalgrpc

import (
	"context"
	"time"

	"github.com/NoisyPunk/multiarmedbandit/internal/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func (g *GRPCServer) loggingInterceptor(ctx context.Context, request interface{},
	serverInfo *grpc.UnaryServerInfo, grpcHandler grpc.UnaryHandler,
) (interface{}, error) {
	l := logger.FromContext(g.ctx)
	start := time.Now()
	dateTime := time.DateTime

	response, err := grpcHandler(ctx, request)

	l.Info(
		"grpc request stats",
		zap.Duration("latency", time.Since(start)),
		zap.String("date_time", dateTime),
		zap.String("method", serverInfo.FullMethod),
	)
	return response, err
}
