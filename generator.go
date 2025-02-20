package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// generateTSX creates the TSX component based on the SVG files
func generateTSX(name, filledSVG, strokeSVG string, hasStroke bool) string {
	componentName := toPascalCase(name)

	// Modify SVG attributes
	filledSVG = transformSVG(filledSVG, "filled")
	strokeSVG = transformSVG(strokeSVG, "stroke")

	tsxTemplate := `import { SvgProps } from '@/helpers/interfaces';
import { FC } from 'react';

const %s: FC<SvgProps> = ({
    filled = false,
    color = '#191C1E',
    onClick = () => {},
    className = '',
    size = 24,
    strokeWidth = 1.5
}) => {
    %s
};

export default %s;
`
	if hasStroke {
		return fmt.Sprintf(tsxTemplate, componentName, fmt.Sprintf(`
    if (filled)
        return (%s);
    return (%s);
`, filledSVG, strokeSVG), componentName)
	}

	return fmt.Sprintf(tsxTemplate, componentName, fmt.Sprintf("return (%s);", filledSVG), componentName)
}

// generateIndexFile creates the `index.ts` file with all exports
func generateIndexFile(outputFolder string) error {
	files, err := os.ReadDir(outputFolder)
	if err != nil {
		return err
	}

	var imports []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".tsx") && file.Name() != "index.ts" {
			name := strings.TrimSuffix(file.Name(), ".tsx")
			imports = append(imports, fmt.Sprintf("import %s from './%s';", name, name))
		}
	}

	indexContent := strings.Join(imports, "\n") + "\n\nexport { " + strings.Join(imports, ", ") + " };"
	return os.WriteFile(filepath.Join(outputFolder, "index.ts"), []byte(indexContent), 0644)
}
