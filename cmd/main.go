package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/NoisyPunk/multiarmedbandit/internal/configs"
	"github.com/NoisyPunk/multiarmedbandit/internal/logger"
	"github.com/NoisyPunk/multiarmedbandit/internal/storage"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
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

	// get config
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
	// get app

	//create server

	// shutdown server rules

	// start servers

	// wait for end
}
