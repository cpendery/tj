package cmd

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(yamlCmd)
}

var xmlCmd = &cobra.Command{
	Use:   "xml",
	Short: "convert xml to json",
	Run:   xmlExec,
}

func xmlExec(_ *cobra.Command, args []string) {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Error().Err(err).Msg("unable to read from stdin")
		return
	}
	var i interface{}
	if err := xml.Unmarshal(input, &i); err != nil {
		log.Error().Err(err).Msg("unable to parse xml")
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
