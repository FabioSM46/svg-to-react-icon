package parser

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/FabioSM46/svg-to-react-icon/utils"
)

func ReadSVGFiles(folder string) (map[string]string, error) {
	files := make(map[string]string)
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(info.Name(), ".svg") {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		baseName := strings.TrimSuffix(info.Name(), ".svg")
		normalizedBaseName := utils.NormalizeName(baseName)
		files[normalizedBaseName] = string(content)

		return nil
	})

	return files, err
}
