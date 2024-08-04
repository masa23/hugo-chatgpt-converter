package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	OpenAI struct {
		APIToken  string `yaml:"APIToken"`
		Model     string `yaml:"Model"`
		MaxTokens int    `yaml:"MaxTokens"`
	} `yaml:"OpenAI"`
	Prompt string `yaml:"Prompt"`
}

func Load(path string) (*Config, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
