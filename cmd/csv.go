package cmd

import (
	"io"
	"os"

	"github.com/cpendery/tj/tj/formats/csv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(csvCmd)
}

var csvCmd = &cobra.Command{
	Use:   "csv",
	Short: "convert csv to json",
	Run:   csvExec,
}

func csvExec(_ *cobra.Command, args []string) {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Error().Err(err).Msg("unable to read from stdin")
		return
	}

	output, err := csv.Encode(input)
	if err != nil {
		log.Error().Err(err).Msg("unable to convert csv -> json")
		return
	}

	if _, err := os.Stdout.Write(output); err != nil {
		log.Error().Err(err).Msg("unable to write output to stdout")
		return
	}
}
