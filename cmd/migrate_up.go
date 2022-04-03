// nolint:dupl
package cmd

import (
	"github.com/VyalovAlexander/Mytheresa/internal/config"
	"github.com/jmoiron/sqlx"
	"os"

	"github.com/rs/zerolog"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

// migrateUpCmd represents the up command
var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "migrates the database to the most recent version available",
	Run: func(cmd *cobra.Command, args []string) {
		log := zerolog.New(os.Stdout).With().Caller().Logger().
			With().Str("command", "migrate up").Timestamp().Logger()

		conn, err := sqlx.Connect(config.AppConfig.DBDriver, config.AppConfig.DBSource)
		if err != nil {
			log.Fatal().Err(err).Msg("create connection to db")
		}

		log.Info().Msgf("connected to db")

		n, err := migrate.Exec(conn.DB, "mysql", Migrations, migrate.Up)
		if err != nil {
			log.Fatal().Err(err).Msg("apply migration to mysql")
		}

		log.Info().Msgf("mysql: applied %d migrations", n)

	},
}

func init() {
	migrateCmd.AddCommand(migrateUpCmd)
}
