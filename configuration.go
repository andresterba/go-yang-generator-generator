package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	packageNameOptionKey = "package_name"
)

type GeneratorOptions struct {
	Option string `yaml:"option,omitempty"`
	Value  string `yaml:"value,omitempty"`
}

type Configuration struct {
	PackageName      string             `yaml:"package_name,omitempty"`
	PathToGenerator  string             `yaml:"path_to_generator,omitempty"`
	PathToModels     string             `yaml:"path_to_models,omitempty"`
	GeneratorOptions []GeneratorOptions `yaml:"generator_options,omitempty"`
	Models           []string           `yaml:"models,omitempty"`
}

func readConfiguration(path string) (*Configuration, error) {
	config := Configuration{}

	readManifest, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(readManifest, &config)
	if err != nil {
		return nil, err
	}

	err = config.Validate()
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Configuration) Validate() error {
	for _, option := range c.GeneratorOptions {
		if option.Option == packageNameOptionKey {
			return fmt.Errorf("global package_name and generator_option for package_name is set! This option will be infered from package_name in your global configuration, so there is no need to set it explicitly")
		}
	}

	packageNameGeneratorOption := GeneratorOptions{
		Option: packageNameOptionKey,
		Value:  c.PackageName,
	}

	c.GeneratorOptions = append(c.GeneratorOptions, packageNameGeneratorOption)

	return nil
}
