package cmd

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tj",
	Short: "go from anything to json",
	Long: `A command line utility for converting almost any file type to json. 
Complete documentation is available at https://github.com/cpendery/tj`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error().Err(err).Msg("unable to execute root command")
		os.Exit(1)
	}
}
