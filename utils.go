package main

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Normalizes names to PascalCase and handles leading numbers
func normalizeName(name string) string {
	re := regexp.MustCompile(`^(\d+)([A-Za-z])`)
	name = re.ReplaceAllString(name, "_${1}${2}")
	return toPascalCase(name)
}

// Converts a name to PascalCase
func toPascalCase(name string) string {
	titleCaser := cases.Title(language.Und)
	words := strings.FieldsFunc(name, func(r rune) bool {
		return r == '-' || r == '_' || r == ' '
	})

	for i, word := range words {
		words[i] = titleCaser.String(word)
	}

	return strings.Join(words, "")
}

// Modifies SVG attributes
func transformSVG(svgContent, id string) string {
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
