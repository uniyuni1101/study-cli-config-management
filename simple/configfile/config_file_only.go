package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/uniyuni1101/studycfg/simple/config"
)

var CurrentConfig = &config.Config{}

func NewJSONConfigLoader(path string) (Loader *JSONConfigLoader, CloseFunc func(), Err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	closeFunc := func() {
		file.Close()
	}

	loader := &JSONConfigLoader{
		dec:           json.NewDecoder(file),
		defaultConfig: config.NewDefault(),
	}

	return loader, closeFunc, nil
}

type JSONConfigLoader struct {
	dec           *json.Decoder
	defaultConfig config.Config
	isSetDefault  bool
}

func (j *JSONConfigLoader) SetDefault(defaultConfig config.Config) {
	j.defaultConfig = defaultConfig
	j.isSetDefault = true
}

func (j *JSONConfigLoader) Load(config *config.Config) error {
	config.Theme = j.defaultConfig.Theme
	config.TickRate = j.defaultConfig.TickRate

	if err := j.dec.Decode(config); err != nil {
		return err
	}

	return nil
}

func main() {

	loader, close, err := NewJSONConfigLoader("config.json")
	if err != nil {
		log.Fatal("config.json", err)
	}
	defer close()

	config := &config.Config{}
	if err := loader.Load(config); err != nil {
		log.Fatal("load error", err)
	}

	fmt.Printf("%#+v\n", config)

}
