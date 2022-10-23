package processor

import (
	"fmt"
	"io"
	"os"
	"sort"
	"sync"

	"github.com/cpendery/tj/tj/formats"
	"github.com/rs/zerolog/log"
)

type rankedBlob struct {
	blob []byte
	rank int
}

func LoadBlobs(stdinIsFromPipe, inputIsStrings bool, args []string) ([][]byte, error) {
	var blobs [][]byte

	if stdinIsFromPipe {
		log.Debug().Msg("detected user input from pipe, loading...")
		blob, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("unable to read from stdin: %w", err)
		}
		blobs = append(blobs, blob)
	} else {
		for _, arg := range args {
			if inputIsStrings {
				log.Debug().Msg("detected user input is set of strings, loading...")
				blobs = append(blobs, []byte(arg))
			} else {
				log.Debug().Msg("detected user input is set of files, loading...")
				f, err := os.Open(arg)
				if err != nil {
					return nil, fmt.Errorf("unable to open %s: %w", arg, err)
				}
				blob, err := io.ReadAll(f)
				if err != nil {
					return nil, fmt.Errorf("unable to read all bytes from %s: %w", arg, err)
				}
				err = f.Close()
				if err != nil {
					return nil, fmt.Errorf("unable to close %s: %w", arg, err)
				}
				blobs = append(blobs, blob)
			}
		}
	}
	log.Debug().Msg("completed loading user input")
	return blobs, nil
}

func Run(blobs [][]byte) error {
	errChan := make(chan error, len(blobs))
	resChan := make(chan rankedBlob, len(blobs))
	wg := sync.WaitGroup{}
	wg.Add(len(blobs))

	for i, b := range blobs {
		blob := b
		idx := i
		go func() {
			defer wg.Done()
			result, err := formats.RunAllFormatters(blob)
			if err != nil {
				errChan <- fmt.Errorf("failed to process user input #%d: %w", idx, err)
				return
			}
			resChan <- rankedBlob{
				blob: result,
				rank: idx,
			}
		}()
	}

	wg.Wait()

	numErrors := len(errChan)
	var err error = nil
	for i := len(errChan); i > 0; i-- {
		err = <-errChan
		log.Error().Err(err)
	}
	if err != nil {
		return fmt.Errorf("formatters failed on %d/%d user inputs", numErrors, len(blobs))
	}

	numResults := len(resChan)
	results := make([]rankedBlob, numResults)
	for i := 0; i < numResults; i++ {
		results[i] = <-resChan
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].rank < results[j].rank
	})

	for _, result := range results {
		fmt.Println(string(result.blob))
	}
	return nil
}
