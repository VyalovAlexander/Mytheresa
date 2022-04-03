package cmd

import (
	"context"
	"io"
	http2 "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/VyalovAlexander/Mytheresa/internal/config"
	"github.com/VyalovAlexander/Mytheresa/internal/server/protocol/http"
	v1 "github.com/VyalovAlexander/Mytheresa/internal/server/protocol/http/api/v1"
	"github.com/VyalovAlexander/Mytheresa/internal/service"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/ziflex/lecho/v3"
)

// serviceStartCmd represents the start command
var serviceStartCmd = &cobra.Command{
	Use:   "start",
	Short: "starts the service",
	Run: func(cmd *cobra.Command, args []string) {
		var lw io.Writer
		os.Stderr = os.Stdout
		lw = os.Stdout

		log := zerolog.New(lw).With().Caller().Str("service", "app").Timestamp().Logger()

		/*conn, err := sqlx.Connect(config.AppConfig.DBDriver, config.AppConfig.DBSource)
		if err != nil {
			log.Fatal().Err(err).Msg("connect to db")
		}*/

		h := v1.NewHandler(
			v1.WithLogger(log),
		)

		srv := http.NewServer(
			http.WithAddr(config.AppConfig.ServerAddress),
			http.WithLogger(lecho.From(log)),
			http.WithHandler(h),
		)

		// create service
		svc := service.New(
			service.WithServer(srv),
			service.WithLogger(log),
		)

		go func() {
			if err := svc.Start(); err != nil && err != http2.ErrServerClosed {
				log.Error().Err(err).Msg("service start failed")
			}
		}()

		defer func() {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if err := svc.Shutdown(ctx); err != nil {
				log.Fatal().Err(err).Msg("shutdown failed")
			} else {
				log.Info().Msg("shutdown succeed")
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 10 seconds.
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
	},
}

func init() {
	serviceCmd.AddCommand(serviceStartCmd)
}
