package cmd

import (
	"fmt"
	"math"
	"os"

	"github.com/cpendery/tj/tj/processor"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: `tj [flags] [files...]
  tj [flags] -s [strings...]`,
		Short: "go from anything to json",
		Long: `tj - command-line JSON converter 

tj is a tool for converting other configuration file formats to JSON 
and producing the results on standard output. When reading from stdin,
input will be treated as a single string.

complete documentation is available at https://github.com/cpendery/tj`,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE:          rootExec,
	}
	flagStrings *bool
	verbosity   *int
)

func init() {
	flagStrings = rootCmd.PersistentFlags().BoolP("strings", "s", false, "toggles tj to process the input as strings rather than files")
	verbosity = rootCmd.PersistentFlags().CountP("verbose", "v", "increase verbosity (-v = error, -vv = info)")
}

func rootExec(_ *cobra.Command, args []string) error {
	logLevel := zerolog.Level(math.Max(float64(int(zerolog.FatalLevel)-*verbosity), 0))
	zerolog.SetGlobalLevel(logLevel)

	stat, _ := os.Stdin.Stat()
	stdinIsFromPipe := (stat.Mode() & os.ModeCharDevice) == 0
	blobs, err := processor.LoadBlobs(stdinIsFromPipe, *flagStrings, args)
	if err != nil {
		msg := "failed to load user input"
		log.Error().Err(err).Msg(msg)
		return fmt.Errorf(msg)
	}

	err = processor.Run(blobs)
	if err != nil {
		msg := "failed to run formatters on user input"
		log.Error().Err(err).Msg(msg)
		return fmt.Errorf(msg)
	}
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error().Err(err).Msg("unable to execute root command")
		os.Exit(1)
	}
}
