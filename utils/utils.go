package utils

import (
	"regexp"
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
		"fill-opacity":      "fillOpacity",
		"stroke-width":      "strokeWidth",
		"stroke-linecap":    "strokeLinecap",
		"stroke-linejoin":   "strokeLinejoin",
		"clip-path":         "clipPath",
		"stroke-miterlimit": "strokeMiterlimit",
		"stroke-opacity":    "strokeOpacity",
		"text-anchor":       "textAnchor",
		"dominant-baseline": "dominantBaseline",
	}
	for dash, camel := range dashToCamel {
		svgContent = strings.ReplaceAll(svgContent, dash+"=", camel+"=")
	}
	reSVG := regexp.MustCompile(`<svg([^>]*)>`)
	svgContent = reSVG.ReplaceAllStringFunc(svgContent, func(match string) string {
		svgTag := match
		attrs := svgTag[4 : len(svgTag)-1]
		reWidth := regexp.MustCompile(`width="[^"]*"`)
		attrs = reWidth.ReplaceAllString(attrs, "")
		reHeight := regexp.MustCompile(`height="[^"]*"`)
		attrs = reHeight.ReplaceAllString(attrs, "")
		newAttrs := ""
		newAttrs += ` width={size}`
		newAttrs += ` height={size}`
		newAttrs += ` id="` + id + `" className={className} onClick={onClick}`
		if len(attrs) > 0 {
			newAttrs = " " + attrs + newAttrs
		}

		return `<svg` + newAttrs + `>`
	})
	reFill := regexp.MustCompile(`fill="([^"]*)"`)
	svgContent = reFill.ReplaceAllStringFunc(svgContent, func(match string) string {
		fillValue := reFill.FindStringSubmatch(match)[1]
		if fillValue != "none" {
			return `fill={color}`
		}
		return match
	})
	reStroke := regexp.MustCompile(`stroke="([^"]*)"`)
	svgContent = reStroke.ReplaceAllStringFunc(svgContent, func(match string) string {
		strokeValue := reStroke.FindStringSubmatch(match)[1]
		if strokeValue != "none" {
			return `stroke={color}`
		}
		return match
	})
	reStrokeWidth := regexp.MustCompile(`strokeWidth="([^"]*)"`)
	svgContent = reStrokeWidth.ReplaceAllStringFunc(svgContent, func(match string) string {
		return `strokeWidth={strokeWidth}`
	})

	return svgContent
}
