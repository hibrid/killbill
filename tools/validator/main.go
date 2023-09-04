package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	libxml2 "github.com/lestrrat-go/libxml2"
	xsd "github.com/lestrrat-go/libxml2/xsd"
)

func validateXMLAgainstXSD(xmlContent, xsdPath string) error {
	// Load XSD schema
	schema, err := xsd.ParseFromFile(xsdPath)
	if err != nil {
		return fmt.Errorf("failed to parse XSD: %v", err)
	}
	defer schema.Free()

	// Parse XML content
	doc, err := libxml2.ParseString(xmlContent)
	if err != nil {
		return fmt.Errorf("failed to parse XML: %v", err)
	}
	defer doc.Free()

	// Validate XML against schema
	if err := schema.Validate(doc); err != nil {
		for _, e := range err.(xsd.SchemaValidationError).Errors() {
			log.Printf("error: %s", e.Error())
		}
		return fmt.Errorf("validation error: %v", err)
	}

	return nil
}

func main() {
	var xmlType string
	var xmlPath string

	// Define the CLI flags
	flag.StringVar(&xmlType, "type", "", "Type of XML (catalog or overdue)")
	flag.StringVar(&xmlPath, "path", "", "Path to the XML file to validate")
	flag.Parse()

	// Ensure XML type is valid
	if xmlType != "catalog" && xmlType != "overdue" {
		log.Fatalf("Invalid XML type specified. Please choose either 'catalog' or 'overdue'.")
	}

	// Ensure XML path is provided
	if xmlPath == "" {
		log.Fatalf("Please provide a path to the XML file to validate.")
	}

	// Construct the path to the appropriate XSD schema based on XML type
	projectRoot := filepath.Join("..", "..")
	xsdPath := filepath.Join(projectRoot, "schemas", fmt.Sprintf("%s.xsd", xmlType))

	// Read the XML content from the provided file path
	xmlContent, err := os.ReadFile(xmlPath)
	if err != nil {
		log.Fatalf("Failed to read XML file: %v", err)
	}

	// Validate the XML content
	err = validateXMLAgainstXSD(string(xmlContent), xsdPath)
	if err != nil {
		log.Fatalf("XML is not valid: %v", err)
	} else {
		fmt.Println("XML is valid!")
	}
}
