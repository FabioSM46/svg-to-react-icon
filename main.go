package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/FabioSM46/svg-to-react-icon/generator"
	"github.com/FabioSM46/svg-to-react-icon/parser"
)

const outputFolder = "./icons"

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run main.go <filledFolder> <strokeFolder>")
	}
	firstFolder := os.Args[1]
	secondFolder := os.Args[2]
	if err := os.MkdirAll(outputFolder, os.ModePerm); err != nil {
		log.Fatal("Failed to create output directory:", err)
	}
	svgFiles1, err := parser.ReadSVGFiles(firstFolder)
	if err != nil {
		log.Fatal("Error reading first folder:", err)
	}
	svgFiles2, err := parser.ReadSVGFiles(secondFolder)
	if err != nil {
		log.Fatal("Error reading second folder:", err)
	}
	for normalizedName, filledSVG := range svgFiles1 {
		strokeSVG, exists := svgFiles2[normalizedName]
		tsxContent := generator.GenerateTSX(normalizedName, filledSVG, strokeSVG, exists)

		outputFilePath := filepath.Join(outputFolder, normalizedName+".tsx")
		err := os.WriteFile(outputFilePath, []byte(tsxContent), 0644)
		if err != nil {
			log.Printf("Failed to write %s: %v\n", outputFilePath, err)
		}
	}
	if err := generator.GenerateIndexFile(outputFolder); err != nil {
		log.Fatal("Error generating index.ts:", err)
	}

	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting root directory: %v", err)
	}
	prettierConfigPath := filepath.Join(rootDir, ".prettierrc")
	if _, err := os.Stat(prettierConfigPath); os.IsNotExist(err) {
		prettierConfigPath = filepath.Join(rootDir, ".prettierrc")
	}
	if _, err := os.Stat(prettierConfigPath); !os.IsNotExist(err) {
		err = filepath.Walk(outputFolder, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() || !strings.HasSuffix(info.Name(), ".tsx") && !strings.HasSuffix(info.Name(), ".ts") {
				return nil
			}
			prettierArgs := []string{"--write", path}
			if _, err := os.Stat(prettierConfigPath); !os.IsNotExist(err) {
				prettierArgs = append(prettierArgs, "--config", prettierConfigPath)
			}
			cmd := exec.Command("prettier", prettierArgs...)

			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("Prettier error for %s: %v\nOutput:%s\n", path, err, output)
				log.Printf("Error running prettier for %s: %v", path, err)
			} else {
				fmt.Printf("Prettier output for %s: %s\n", path, output)
			}

			return nil
		})
		if err != nil {
			log.Fatalf("Error running prettier: %v", err)
		}
	} else {
		fmt.Println(".prettierrc not found in root directory. Skipping formatting.")
	}
	fmt.Println("TSX files successfully generated!")
}
