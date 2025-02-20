package main

import (
	"os"
	"path/filepath"
	"strings"
)

// readSVGFiles scans a folder and loads all SVG files into a map
func readSVGFiles(folder string) (map[string]string, error) {
	files := make(map[string]string)

	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(info.Name(), ".svg") {
			return nil
		}

		// Read file content
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Normalize filename for matching
		baseName := strings.TrimSuffix(info.Name(), ".svg")
		files[normalizeName(baseName)] = string(content)

		return nil
	})

	return files, err
}
