package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/FabioSM46/svg-to-react-icon/utils"
)

func GenerateTSX(name, filledSVG, strokeSVG string, hasStroke bool) string {
	componentName := utils.NormalizeName(name)
	filledSVG = utils.TransformSVG(filledSVG, "filled")
	strokeSVG = utils.TransformSVG(strokeSVG, "stroke")

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
		return fmt.Sprintf(tsxTemplate, name, fmt.Sprintf(`
    if (filled)
        return (%s);
    return (%s);
`, filledSVG, strokeSVG), name)
	}

	return fmt.Sprintf(tsxTemplate, componentName, fmt.Sprintf("return (%s);", filledSVG), componentName)
}

func GenerateIndexFile(outputFolder string) error {
	files, err := os.ReadDir(outputFolder)
	if err != nil {
		return err
	}
	var imports []string
	var exports []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".tsx") && file.Name() != "index.ts" {
			name := strings.TrimSuffix(file.Name(), ".tsx")
			imports = append(imports, fmt.Sprintf("import %s from './%s';", name, name))
			exports = append(exports, name)
		}
	}
	exportStatement := fmt.Sprintf("export { %s };", strings.Join(exports, ", "))
	indexContent := strings.Join(imports, "\n") + "\n\n" + exportStatement

	return os.WriteFile(filepath.Join(outputFolder, "index.ts"), []byte(indexContent), 0644)
}
