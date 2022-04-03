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

// migrateDownCmd represents the down command
var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "undo a database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		log := zerolog.New(os.Stdout).With().Caller().Logger().
			With().Str("command", "migrate down").Timestamp().Logger()

		conn, err := sqlx.Connect(config.AppConfig.DBDriver, config.AppConfig.DBSource)
		if err != nil {
			log.Fatal().Err(err).Msg("create connection to db")
		}

		log.Info().Msgf("connected to db")

		n, err := migrate.Exec(conn.DB, "mysql", Migrations, migrate.Down)
		if err != nil {
			log.Fatal().Err(err).Msg("undo migration to db")
		}

		log.Info().Msgf("sb: undo %d migrations", n)
	},
}

func init() {
	migrateCmd.AddCommand(migrateDownCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateDownCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateDownCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
