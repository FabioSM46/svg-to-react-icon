package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/FabioSM46/svg-to-react-icon/generator"
	"github.com/FabioSM46/svg-to-react-icon/parser"
)

const outputFolder = "./icons"

func main() {
	// Ensure two arguments (input folders) are provided
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run main.go <filledFolder> <strokeFolder>")
	}

	firstFolder := os.Args[1]
	secondFolder := os.Args[2]

	// Ensure output folder exists
	if err := os.MkdirAll(outputFolder, os.ModePerm); err != nil {
		log.Fatal("Failed to create output directory:", err)
	}

	// Read SVG files from both input folders
	svgFiles1, err := parser.ReadSVGFiles(firstFolder)
	if err != nil {
		log.Fatal("Error reading first folder:", err)
	}

	svgFiles2, err := parser.ReadSVGFiles(secondFolder)
	if err != nil {
		log.Fatal("Error reading second folder:", err)
	}

	// Match SVGs by filename and generate TSX components
	for name, filledSVG := range svgFiles1 {
		strokeSVG, exists := svgFiles2[name]
		tsxContent := generator.GenerateTSX(name, filledSVG, strokeSVG, exists)

		// Write TSX file
		outputFilePath := filepath.Join(outputFolder, name+".tsx")
		err := os.WriteFile(outputFilePath, []byte(tsxContent), 0644)
		if err != nil {
			log.Printf("Failed to write %s: %v\n", outputFilePath, err)
		}
	}

	// Generate index.ts file
	if err := generator.GenerateIndexFile(outputFolder); err != nil {
		log.Fatal("Error generating index.ts:", err)
	}

	fmt.Println("TSX files successfully generated!")
}
