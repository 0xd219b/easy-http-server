package parser

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Parser struct {
	URL    string `yaml:"url"`
	Method string `yaml:"method"`
	Resp   string `yaml:"resp"`
	Status int    `yaml:"status"`
}

func FromYaml(yamlFile string) ([]Parser, error) {
	var parsers []Parser
	yamlFileBytes, err := os.ReadFile(yamlFile)
	if err != nil {
		return nil, err
	}

	parsers, err = getParsersFromRawBytes(yamlFileBytes)
	if err != nil {
		return nil, err
	}
	return parsers, nil
}

func getParsersFromRawBytes(rawBytes []byte) ([]Parser, error) {
	var parsers []Parser
	err := yaml.Unmarshal(rawBytes, &parsers)
	if err != nil {
		return nil, err
	}
	return parsers, nil
}

func (p Parser) GetURL() string {
	return p.URL
}

func (p Parser) GetMethodType() string {
	return p.Method
}

func (p Parser) GetResp() []byte {
	return []byte(p.Resp)
}

func (p Parser) GetStatus() int {
	return p.Status
}
