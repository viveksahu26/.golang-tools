package main

import (
	"fmt"
	"log"
	"os"

	spdxtv "github.com/spdx/tools-golang/tagvalue"
)

func main() {
	file, err := os.Open("example-spdx-2-2.spdx")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Decode (unmarshal) the SPDX JSON data
	doc, err := spdxtv.Read(file)
	if err != nil {
		log.Fatalf("Error unmarshaling SPDX JSON: %v", err)
	}

	// Retrieve and print the SPDX version
	fmt.Println("SPDX Version:", doc.SPDXVersion) // Result: SPDX Version: SPDX-2.3
}
