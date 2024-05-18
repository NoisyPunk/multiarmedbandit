package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	rotator "github.com/NoisyPunk/multiarmedbandit/internal/app"
	"github.com/NoisyPunk/multiarmedbandit/internal/configs"
	"github.com/NoisyPunk/multiarmedbandit/internal/logger"
	internalgrpc "github.com/NoisyPunk/multiarmedbandit/internal/server/grpc"
	"github.com/NoisyPunk/multiarmedbandit/internal/storage"
	"go.uber.org/zap"
)

var configFile string

func init() {
	flag.StringVar(&configFile,
		"rotator_config",
		"./configs/rotator_config.yaml",
		"path to rotator configuration file",
	)
}

func main() {
	flag.Parse()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	config, err := rotatorconfig.GetConfig(configFile)
	if err != nil {
		fmt.Printf("can't get config from config file: %s", err.Error())
		os.Exit(1) //nolint:gocritic
	}

	log := logger.New(config.Server.LogLevel)
	ctx = logger.ContextLogger(ctx, log)

	err = storage.Migrate(config)
	if err != nil {
		log.Error("migration has failed", zap.String("error_message", err.Error()))
		os.Exit(1)
	}
	app, err := rotator.New(ctx, config)
	if err != nil {
		log.Error("app creation has failed", zap.String("error_message", err.Error()))
		os.Exit(1)
	}

	grpcServer := internalgrpc.NewGRPCServer(ctx, app, config.Server.GrpcPort)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		<-ctx.Done()
		grpcServer.Stop()
		wg.Done()
	}()

	go func() {
		if err = grpcServer.Start(); err != nil {
			log.Error("failed to start grpc server", zap.String("error", err.Error()))
			grpcServer.Stop()
		}
	}()

	log.Info("rotator is running...", zap.String("start_time", time.Now().String()))
	wg.Wait()
}
