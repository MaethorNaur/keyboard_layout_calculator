package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func (config *Config) Load() {
	var err error
	var f *os.File
	var d []byte
	if f, err = os.Open(config.File); err != nil {
		return
	}
	defer f.Close()
	if d, err = io.ReadAll(f); err != nil {
		return
	}
	_ = yaml.Unmarshal(d, &config.data)
}

func (config *Config) save() (err error) {
	var f *os.File
	var d []byte

	if f, err = os.Create(config.File); err != nil {
		return
	}
	defer f.Close()

	if d, err = yaml.Marshal(config.data); err != nil {
		return
	}
	_, err = f.Write(d)
	return
}
