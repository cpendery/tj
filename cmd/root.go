package cmd

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/cpendery/tj/tj/formats"
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
	inputIsStrings *bool
)

func init() {
	inputIsStrings = rootCmd.PersistentFlags().BoolP("strings", "s", false, "toggles tj to process the input as strings rather than files")
}

func rootExec(_ *cobra.Command, args []string) error {
	stat, _ := os.Stdin.Stat()
	var blobs [][]byte

	stdinIsFromPipe := (stat.Mode() & os.ModeCharDevice) == 0
	if stdinIsFromPipe {
		blob, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		blobs = append(blobs, blob)
	} else {
		for _, arg := range args {
			if *inputIsStrings {
				blobs = append(blobs, []byte(arg))
			} else {
				f, err := os.Open(arg)
				defer f.Close()
				if err != nil {
					return err
				}
				blob, err := io.ReadAll(f)
				if err != nil {
					return err
				}
				blobs = append(blobs, blob)
			}
		}
	}

	errChan := make(chan error, len(blobs))
	resChan := make(chan []byte, len(blobs))
	wg := sync.WaitGroup{}
	wg.Add(len(blobs))

	for _, b := range blobs {
		blob := b
		go func() {
			defer wg.Done()
			result, err := formats.RunAllFormatters(blob)
			if err != nil {
				errChan <- err
				return
			}
			resChan <- result
		}()
	}

	wg.Wait()

	for i := len(errChan); i > 0; i-- {
		err := <-errChan
		return err
	}
	for i := len(resChan); i > 0; i-- {
		result := <-resChan
		fmt.Fprintf(os.Stdout, "%s\n", string(result))
	}

	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error().Err(err).Msg("unable to execute root command")
		os.Exit(1)
	}
}
