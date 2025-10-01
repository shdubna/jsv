package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/xeipuuv/gojsonschema"
)

var gitTag string

var (
	schemaPath   = flag.String("schema", "schema.json", "Path to json schema file.")
	documentPath = flag.String("document", "package.json", "Path to json document.")
	errorMessage = flag.String("message", "The document is not valid. see errors :", "Message if validation error.")
	version      = flag.Bool("version", false, "Show version number and quit.")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Println(gitTag)
		os.Exit(0)
	}

	schemaAbsPath, err := filepath.Abs(*schemaPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	documentAbsPath, err := filepath.Abs(*documentPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	schemaLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s", schemaAbsPath))
	documentLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s", documentAbsPath))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("%s \n", *errorMessage)
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
		os.Exit(1)
	}
}
