package config

import (
	"flag"
	"fmt"
	"io"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Data struct {
	Port int `yaml:"port"`
	File string
}

func (d *Data) InitFlag() {
	filePathPtr := flag.String("c", "config.yaml", "Relative path of yaml config file")
	flag.Parse()
	d.File = *filePathPtr
}
func (d *Data) ReadYAML() ([]byte, error) {
	f, err := os.Open(d.File)
	if err != nil {
		return nil, fmt.Errorf("yaml read error: %w", err)
	}
	return io.ReadAll(f)
}

func (d *Data) DecodeYAML(y []byte) (err error) {
	type config struct {
		Server Data `yaml:"server"`
	}
	var s config
	err = yaml.Unmarshal(y, &s)
	if err != nil {
		return fmt.Errorf("yaml decode error: %w", err)
	}
	d.Port = s.Server.Port
	return
}
