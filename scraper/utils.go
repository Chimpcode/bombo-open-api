package scraper

import "strings"

func GetIdFromMatchUrl(urlMatch string) string {
	fixedURL := ""
	if strings.HasSuffix(urlMatch, "/") {
		fixedURL = urlMatch[0 : len(urlMatch)-2]
	} else {
		fixedURL = urlMatch
	}

	chunks := strings.Split(fixedURL, "/")
	return chunks[len(chunks)-1]
}

func ParseToValidCamelCase(text string) string {
	finalText := strings.ToLower(text)
	finalText = strings.Replace(finalText, " ", "_", -1)
	finalText = strings.Replace(finalText, "-", "_", -1)
	return finalText
}
