package main

import (
	"fmt"
	"os"
	"text/template"
)

var generatorTemplate = `package {{.PackageName}}

// This file is a placeholder in order to ensure that Go does not
// find this directory to contain an empty package.

//go:generate {{.PathToGenerator}} -path={{.PathToModels}} {{range .GeneratorOptions}} -{{.Option}}={{.Value}} {{end}} {{range .Models}} {{.}} {{end}}
`


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

func generateYgotGeneratorFileFromInput(configurationPath string, outputPath string) error {
	manifest, err := readConfiguration(configurationPath)
	if err != nil {
		return err
	}

	err = executeAndWriteManifest(outputPath, manifest)
	if err != nil {
		return err
	}

	return nil
}
func executeAndWriteManifest(path string, manifest *Configuration) error {
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
