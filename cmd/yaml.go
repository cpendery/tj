package cmd

import (
	"io"
	"os"

	"github.com/cpendery/tj/tj/formats/yaml"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(yamlCmd)
}

var yamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: "convert yaml to json",
	Run:   yamlExec,
}

func yamlExec(_ *cobra.Command, args []string) {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Error().Err(err).Msg("unable to read from stdin")
		return
	}

	output, err := yaml.Encode(input)
	if err != nil {
		log.Error().Err(err).Msg("unable to convert yaml -> json")
		return
	}

	if _, err := os.Stdout.Write(output); err != nil {
		log.Error().Err(err).Msg("unable to write output to stdout")
		return
	}
}
