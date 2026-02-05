package configuration

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ExevoConfig struct {
	Version   string              `yaml:"version"`
	Global    GlobalConfig        `ymal:"global"`
	Resources map[string]Resource `yaml:"resources"`
}

type GlobalConfig struct {
	TerraformVersion string                    `yaml:"terraform_version"`
	Backend          BackendConfig             `yaml:"backend"`
	Providers        map[string]ProviderConfig `yaml:"providers"`
}

type BackendConfig struct {
	Type   string                 `yaml:"type"`
	Config map[string]interface{} `yaml:"config"`
}

type ProviderConfig struct {
	Source  string `yaml:"source"`
	Version string `yaml:"version"`
}

type Resource struct {
	Source          string                 `yaml:"source"`
	PrimaryKey      string                 `yaml:"primary_key"`
	ResourceAddress string                 `yaml:"resource_address"`
	Mappings        map[string]string      `yaml:"mappings"`
	Static          map[string]interface{} `yaml:"static"`
}

func Load(filename string) (*ExevoConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler config: %w", err)
	}

	var cfg ExevoConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("erro ao parsear yaml: %w", err)
	}

	return &cfg, nil
}
