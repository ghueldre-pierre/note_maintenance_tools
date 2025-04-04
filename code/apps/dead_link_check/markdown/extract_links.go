package markdown

import (
	"strings"
)

func ExtractAllLinks(document string) []string {
	tokens := [4]rune{'[', ']', '(', ')'}
	checkPos := 0
	var prevRune rune
	linkBuilder := strings.Builder{}
	foundLinkList := []string{}
	for _, r := range document {
		if r == tokens[checkPos] {
			if checkPos == 0 && prevRune == '!' {
				continue
			}
			checkPos += 1
			if checkPos == 3 {
				// don't write '('
				continue
			}
		}
		prevRune = r
		if checkPos < 3 {
			continue
		}
		if checkPos == 3 {
			if r == '"' {
				checkPos += 1
			} else {
				linkBuilder.WriteRune(r)
				continue
			}
		}
		if checkPos == 4 {
			foundLinkList = append(foundLinkList, strings.TrimSpace(linkBuilder.String()))
			linkBuilder.Reset()
			checkPos = 0
		}
	}
	return foundLinkList
}
