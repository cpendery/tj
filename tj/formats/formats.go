package formats

import (
	"fmt"

	"github.com/cpendery/tj/tj/formats/csv"
	"github.com/cpendery/tj/tj/formats/toml"
	"github.com/cpendery/tj/tj/formats/xml"
	"github.com/cpendery/tj/tj/formats/yaml"
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

func RunFormatter(blob []byte, formatter string) ([]byte, error) {
	return nil, nil
}

func RunAllFormatters(blob []byte) ([]byte, error) {
	for _, f := range formatters {
		res, err := f.Encode(blob)
		if err == nil {
			return res, nil
		}
	}
	return nil, fmt.Errorf("unable to process input using any formatter")
}
