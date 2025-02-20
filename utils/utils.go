package utils

import (
	"strings"
	"unicode"
)

func NormalizeName(name string) string {
	var result []rune
	var capitalizeNext bool
	for i, r := range name {
	if r == '-' { 
		continue
	}
			if i == 0 && unicode.IsDigit(r) {
					result = append(result, '_', r)
					capitalizeNext = false 
			} else if i == 0 { 
		result = append(result, unicode.ToUpper(r)) 
	} else if r == '_' {
					capitalizeNext = true
			} else if capitalizeNext {
					result = append(result, unicode.ToUpper(r))
					capitalizeNext = false
			} else if i > 0 && unicode.IsDigit(r) && name[i-1] == '_' {
				result = append(result, unicode.ToUpper(r))
			} else {
					result = append(result, unicode.ToLower(r))
			}
	}

	return string(result)
}

func TransformSVG(svgContent, id string) string {
	dashToCamel := map[string]string{
		"fill-opacity":       "fillOpacity",
		"stroke-width":       "strokeWidth",
		"stroke-linecap":     "strokeLinecap",
		"stroke-linejoin":    "strokeLinejoin",
		"clip-path":          "clipPath",
		"stroke-miterlimit":  "strokeMiterlimit",
		"stroke-opacity":     "strokeOpacity",
		"text-anchor":        "textAnchor",
		"dominant-baseline":  "dominantBaseline",
	}
	for dash, camel := range dashToCamel {
		svgContent = strings.ReplaceAll(svgContent, dash+"=", camel+"=")
	}
	svgContent = strings.Replace(svgContent, "<svg", `<svg width={size} height={size} id="`+id+`" className={className} onClick={onClick}`, 1)
	svgContent = strings.Replace(svgContent, `fill="`, `fill={color}"`, -1)
	svgContent = strings.Replace(svgContent, `stroke="`, `stroke={color}"`, -1)
	svgContent = strings.Replace(svgContent, `stroke-width="`, `strokeWidth={strokeWidth}"`, -1)

	return svgContent
}
