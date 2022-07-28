package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func deleteGeneratedFiles(t *testing.T, pathToFile string) {
	err := os.Remove(pathToFile)
	if err != nil {
		t.Fatalf("Failed to delete generated file %s: %v.", pathToFile, err)
	}
}

// Based on https://eli.thegreenplace.net/2022/file-driven-testing-in-go/
func TestGenerateFiles(t *testing.T) {
	// Find the paths of all input files in the data directory.
	paths, err := filepath.Glob(filepath.Join("testdata", "*.input"))
	if err != nil {
		t.Fatal(err)
	}
	for _, configurationPath := range paths {
		_, filename := filepath.Split(configurationPath)
		testname := filename[:len(filename)-len(filepath.Ext(configurationPath))]
		// Each path turns into a test: the test name is the filename without the
		// extension.
		t.Run(testname, func(t *testing.T) {
			outputFile := filepath.Join("testdata", testname+".generated")
			defer deleteGeneratedFiles(t, outputFile)
			err := generateYgotGeneratorFileFromInput(configurationPath, outputFile)
			if err != nil {
				t.Fatal(err)
			}

			// Each input file is expected to have a "golden output" file, with the
			// same path except the .input extension is replaced by .golden
			goldenfile := filepath.Join("testdata", testname+".golden")
			want, err := os.ReadFile(goldenfile)
			if err != nil {
				t.Fatal("error reading golden file:", err)
			}
			generated, err := os.ReadFile(outputFile)
			if err != nil {
				t.Fatal("error reading generated file:", err)
			}

			if !bytes.Equal(generated, want) {
				t.Errorf("\n==== got:\n%s\n==== want:\n%s\n", generated, want)
			}
		})
	}

}
