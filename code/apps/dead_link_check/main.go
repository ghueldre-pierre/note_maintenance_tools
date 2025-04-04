package main

import (
	"fmt"

	"github.com/ghueldre-pierre/note_maintenance_tools/apps/dead_link_check/markdown"
)

func main() {
	lil := markdown.ExtractAllLinks("[a](b)")

	for _, li := range lil {
		fmt.Println(li)
	}
}
