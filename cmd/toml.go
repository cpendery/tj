package cmd

import (
	"io"
	"os"

	"github.com/cpendery/tj/tj/formats/toml"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tomlCmd)
}

var tomlCmd = &cobra.Command{
	Use:   "toml",
	Short: "convert toml to json",
	Run:   tomlExec,
}

func tomlExec(_ *cobra.Command, _ []string) {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Error().Err(err).Msg("unable to read from stdin")
		return
	}

	output, err := toml.Encode(input)
	if err != nil {
		log.Error().Err(err).Msg("unable to convert toml -> json")
		return
	}

	if _, err := os.Stdout.Write(output); err != nil {
		log.Error().Err(err).Msg("unable to write output to stdout")
		return
	}
}
