package crawler
import (
	"strings"
)
func ExtractLinks(html string) ([]string) {
	var links []string
	for { 
		idx := strings.Index(html, "href=\"")
		if idx == -1 {break}
		html = html[idx+6:]
		end := strings.Index(html, "\"")
		if end == -1 {break}
		url := html[:end]
		links = append(links, url)
		html = html[end:]
	 }
	 return links
}