package toml

import (
	"encoding/json"
	"fmt"

	"github.com/BurntSushi/toml"
)

func Encode(blob []byte) ([]byte, error) {
	var i interface{}
	if err := toml.Unmarshal(blob, &i); err != nil {
		return nil, fmt.Errorf("unable to parse toml")
	}
	output, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal toml to json")
	}
	return output, nil
}
