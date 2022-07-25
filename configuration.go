package main

import (
	"os"

	"gopkg.in/yaml.v2"
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
		panic(err)
	}

	err = yaml.Unmarshal(readManifest, &config)
	if err != nil {
		panic(err)
	}

	return &config, nil
}
