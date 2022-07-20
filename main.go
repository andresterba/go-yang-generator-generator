package main

import (
	"fmt"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

var generatorTemplate = `package {{.PackageName}}

// This file is a placeholder in order to ensure that Go does not
// find this directory to contain an empty package.

//go:generate {{.PathToGenerator}} -path={{.PathToModels}} {{range .GeneratorOptions}} -{{.Option}}={{.Value}} {{end}} {{range .Models}} {{.}} {{end}}
`

type GeneratorOptions struct {
	Option string `yaml:"option,omitempty"`
	Value  string `yaml:"value,omitempty"`
}

type Manifest struct {
	PackageName      string             `yaml:"package_name,omitempty"`
	PathToGenerator  string             `yaml:"path_to_generator,omitempty"`
	PathToModels     string             `yaml:"path_to_models,omitempty"`
	GeneratorOptions []GeneratorOptions `yaml:"generator_options,omitempty"`
	Models           []string           `yaml:"models,omitempty"`
}

func main() {
	providedArgs := os.Args

	if len(providedArgs) != 3 {
		showHelp()
	}

	inputPath := providedArgs[1]
	outputPath := providedArgs[2]

	err := generateYgotGeneratorFileFromInput(inputPath, outputPath)
	if err != nil {
		panic(err)
	}

}

func showHelp() {
	fmt.Println(`Usage:
go-yang-generator-generator <path-to-input-file> <path-to-output-file>`)

	os.Exit(1)
}

func generateYgotGeneratorFileFromInput(inputPath string, outputPath string) error {
	manifest, err := readManifest(inputPath)
	if err != nil {
		return err
	}

	err = executeAndWriteManifest(outputPath, manifest)
	if err != nil {
		return err
	}

	return nil
}

func readManifest(path string) (*Manifest, error) {
	config := Manifest{}

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

func executeAndWriteManifest(path string, manifest *Manifest) error {
	t1 := template.New("template")
	t1parsed, err := t1.Parse(generatorTemplate)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	err = t1parsed.Execute(f, &manifest)
	if err != nil {
		panic(err)
	}

	return nil
}
