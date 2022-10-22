package yaml

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type YamlFormatter struct{}

func (YamlFormatter) Encode(blob []byte) ([]byte, error) {
	var i interface{}
	if err := yaml.Unmarshal(blob, &i); err != nil {
		return nil, fmt.Errorf("unable to parse yaml")
	}
	output, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal yaml to json")
	}
	return output, nil
}
