package scraper

import "strings"

func GetIdFromMatchUrl(urlMatch string) string {
	fixedUrl := ""
	if strings.HasSuffix(urlMatch, "/") {
		fixedUrl = urlMatch[0 : len(urlMatch)-2]
	} else {
		fixedUrl = urlMatch
	}

	chunks := strings.Split(fixedUrl, "/")
	return chunks[len(chunks)-1]
}
