package main

import (
	"fmt"
	"log"
	"os"

	spdxjson "github.com/spdx/tools-golang/json"
)

func main() {
	file, err := os.Open("photon.spdx.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Decode (unmarshal) the SPDX JSON data
	doc, err := spdxjson.Read(file)
	if err != nil {
		log.Fatalf("Error unmarshaling SPDX JSON: %v", err)
	}

	// Retrieve and print the SPDX version
	fmt.Println("SPDX Version:", doc.SPDXVersion) // Result: SPDX Version: SPDX-2.3
}
