package TinkConfigParser

import (
	"errors"
	"io"

	"gopkg.in/yaml.v3"
)

type YamlServerConfiguration struct {
	Port    int32  `yaml:"Port"`
	Host    string `yaml:"Host"`
	Timeout int64  `yaml:"Timeout"`
}

func (c *YamlServerConfiguration) ToConfig() ServerConfiguration {
	var result ServerConfiguration
	result.Host = c.Host
	result.Port = c.Port
	result.Timeout = c.Timeout
	return result
}

type yamlConfigParser struct {
}

func GetYamlParser() IConfigParser {
	return &yamlConfigParser{}
}

func (p *yamlConfigParser) ParseConfig(r io.Reader) (ServerConfiguration, error) {
	var result ServerConfiguration
	var c YamlServerConfiguration
	if r == nil {
		return result, errors.New("r *io.Reader can`t be nil")
	}
	d := yaml.NewDecoder(r)
	err := d.Decode(&c)
	if err != nil {
		return result, err
	}
	result = c.ToConfig()
	return result, nil
}
