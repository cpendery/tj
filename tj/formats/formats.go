package formats

import (
	"fmt"
	"reflect"

	"github.com/cpendery/tj/tj/formats/csv"
	"github.com/cpendery/tj/tj/formats/toml"
	"github.com/cpendery/tj/tj/formats/xml"
	"github.com/cpendery/tj/tj/formats/yaml"
	"github.com/rs/zerolog/log"
)

type Formatter interface {
	Encode(blob []byte) ([]byte, error)
}

var formatters []Formatter

func init() {
	formatters = []Formatter{
		xml.XmlFormatter{},
		toml.TomlFormatter{},
		yaml.YamlFormatter{},
		csv.CsvFormatter{},
	}
}

func RunAllFormatters(blob []byte) ([]byte, error) {
	for _, f := range formatters {
		res, err := f.Encode(blob)
		formatterType := reflect.TypeOf(f).String()
		if err == nil {
			log.Debug().Msgf("successfully ran using formatter %s", formatterType)
			return res, nil
		} else {
			log.Info().Err(err).Msgf("unable to run %s on blob", formatterType)
		}
	}
	return nil, fmt.Errorf("unable to process input using any formatter")
}
