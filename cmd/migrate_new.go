package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var cmdTemplateContent = `
-- +migrate Up
-- +migrate Down
`
var tpl *template.Template
var migrationName string

// migrateUpCmd represents the up command
var migrateNewCmd = &cobra.Command{
	Use:   "new",
	Short: "create migration file based on template",
	Run: func(cmd *cobra.Command, args []string) {
		log := zerolog.New(os.Stdout).With().Caller().Logger().
			With().Str("command", "migrate new").Timestamp().Logger()

		if migrationName == "" {
			err := errors.New("a name for the migration is needed")
			log.Fatal().Err(err).Msg("Cannot create migration")
		}

		fileName := fmt.Sprintf("%s-%s.sql", time.Now().Format("20060102150405"), strings.TrimSpace(migrationName))
		pathName := path.Join(Migrations.Dir, fileName)
		f, err := os.Create(pathName)

		if err != nil {
			log.Fatal().Err(err).Msg("Cannot create migration")
		}

		defer func() { _ = f.Close() }()

		if err := tpl.Execute(f, nil); err != nil {
			log.Fatal().Err(err).Msg("Cannot create migration")
		}

		log.Info().Msgf("migration %s was created", migrationName)
	},
}

func init() {
	tpl = template.Must(template.New("new_migration").Parse(cmdTemplateContent))

	migrateCmd.AddCommand(migrateNewCmd)

	migrateNewCmd.Flags().StringVarP(&migrationName, "name", "n", "", "--name filename")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateUpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateUpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
