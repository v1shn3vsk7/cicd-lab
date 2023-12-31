package main

import (
	"context"
	
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"

	"github.com/v1shn3vsk7/cicd-lab/internal/app"
	"github.com/v1shn3vsk7/cicd-lab/internal/config"
)

func main() {
	_ = godotenv.Load(".env")

	cfg := &config.Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal().Msgf("failed to marshal .env variables to config, err: %+v", err)
	}

	if err := app.Run(context.Background(), cfg); err != nil {
		log.Fatal().Msgf("error running app, err: %+v", err)
	}
}
