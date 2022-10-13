package cmd

import (
	"encoding/json"
	"io"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
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
	var i interface{}
	if err := yaml.Unmarshal(input, &i); err != nil {
		log.Error().Err(err).Msg("unable to parse yaml")
		return
	}
	output, err := json.Marshal(i)
	if err != nil {
		log.Error().Err(err).Msg("unable to marshal output to json")
		return
	}
	if _, err := os.Stdout.Write(output); err != nil {
		log.Error().Err(err).Msg("unable to write output to stdout")
		return
	}
}
