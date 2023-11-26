package tests

import (
	"strings"
	"testing"

	tink "github.com/KirillVelichk0/TinkConfigParser/pkg/ConfServer"

	"gopkg.in/yaml.v3"
)

func compareYamlAndNotConfigs(configExpected tink.YamlServerConfiguration, configResult tink.ServerConfiguration) bool {
	return configExpected.Host == configResult.Host &&
		configExpected.Port == configResult.Port &&
		configExpected.Timeout == configResult.Timeout
}

func TestYamlOk(t *testing.T) {
	p := tink.GetYamlParser()
	var configExpected tink.YamlServerConfiguration
	configExpected.Host = "some_host"
	configExpected.Port = 322
	configExpected.Timeout = 100
	out, err := yaml.Marshal(&configExpected)
	if err != nil {
		t.Fatal(err.Error())
	}
	s := string(out)
	r := strings.NewReader(s)
	configResult, err := p.ParseConfig(r)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !compareYamlAndNotConfigs(configExpected, configResult) {
		t.Fatal("Configs not equal")
	}
}

func TestYamlNotOk(t *testing.T) {
	p := tink.GetYamlParser()

	s := "data"
	r := strings.NewReader(s)
	_, err := p.ParseConfig(r)
	if err == nil {
		t.Fatalf("Successful parsed %s", err.Error())
	}
}
