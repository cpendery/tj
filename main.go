package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

func main() {
	f, err := os.Open("t.toml")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	var i interface{}
	err = toml.Unmarshal(b, &i)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(i)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
