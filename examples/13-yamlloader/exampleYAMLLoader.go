// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

// Example for: *yaml*

// This example demonstrates loading an SPDX YAML document from disk into memory,
// and then logging some attributes to the console.
// Run project: go run exampleYAMLLoader.go ../sample-docs/yaml/SPDXYAMLExample-2.2.spdx.yaml
package main

import (
	"fmt"
	"os"
	"strings"

	spdx_yaml "github.com/spdx/tools-golang/yaml"
)

func main() {

	// check that we've received the right number of arguments
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %v <json-file-in>\n", args[0])
		fmt.Printf("  Load SPDX 2.2 JSON file <spdx-file-in>, and\n")
		fmt.Printf("  print portions of its creation info data.\n")
		return
	}

	// open the SPDX file
	fileIn := args[1]
	r, err := os.Open(fileIn)
	if err != nil {
		fmt.Printf("Error while opening %v for reading: %v", fileIn, err)
		return
	}
	defer r.Close()

	// try to load the SPDX file's contents as a json file, version 2.2
	doc, err := spdx_yaml.Load2_2(r)
	if err != nil {
		fmt.Printf("Error while parsing %v: %v", args[1], err)
		return
	}

	// if we got here, the file is now loaded into memory.
	fmt.Printf("Successfully loaded %s\n", args[1])

	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("Some Attributes of the Document:")
	fmt.Printf("Document Name:         %s\n", doc.DocumentName)
	fmt.Printf("DataLicense:           %s\n", doc.DataLicense)
	fmt.Printf("Document Namespace:    %s\n", doc.DocumentNamespace)
	fmt.Printf("SPDX Version:          %s\n", doc.SPDXVersion)
	fmt.Println(strings.Repeat("=", 80))
}
