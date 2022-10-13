package cmd

import (
	"io"
	"os"

	"github.com/cpendery/tj/tj/formats/xml"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(xmlCmd)
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

	output, err := xml.Encode(input)
	if err != nil {
		log.Error().Err(err).Msg("unable to convert xml -> json")
		return
	}

	if _, err := os.Stdout.Write(output); err != nil {
		log.Error().Err(err).Msg("unable to write output to stdout")
		return
	}
}