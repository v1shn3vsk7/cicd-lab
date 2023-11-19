package app

import (
	"context"
	"github.com/v1shn3vsk7/cicd-lab/internal/app/http"
	"github.com/v1shn3vsk7/cicd-lab/internal/utils"

	"github.com/v1shn3vsk7/cicd-lab/internal/config"
)

func Run(ctx context.Context, cfg *config.Config) error {
	_, cancel := context.WithCancel(ctx)

	httpServer := http.New(cfg)
	httpServer.Start()

	utils.GracefulShutdown(cancel)

	return nil
}
