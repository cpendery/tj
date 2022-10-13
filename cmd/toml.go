package cmd

import (
	"encoding/json"
	"io"
	"os"

	"github.com/BurntSushi/toml"
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

func tomlExec(_ *cobra.Command, args []string) {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Error().Err(err).Msg("unable to read from stdin")
		return
	}
	var i interface{}
	if err := toml.Unmarshal(input, &i); err != nil {
		log.Error().Err(err).Msg("unable to parse toml")
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
