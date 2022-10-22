package xml

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type XmlFormatter struct{}

func (XmlFormatter) Encode(blob []byte) ([]byte, error) {
	var i interface{}
	if err := xml.Unmarshal(blob, &i); err != nil {
		return nil, fmt.Errorf("unable to parse xml")
	}
	output, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal xml to json")
	}
	return output, nil
}
