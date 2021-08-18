package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findFirstStringInBracket("Agam"))
	fmt.Println(findFirstStringInBracket("(Agam)"))
	fmt.Println(findFirstStringInBracket("(Firman Agam)"))
	fmt.Println(findFirstStringInBracket(""))
}

func findFirstStringInBracket(str string) string {
	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound >= 0 {
		wordsAfterFirstBracket := str[indexFirstBracketFound:]
		indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
		if indexClosingBracketFound >= 0 {
			return wordsAfterFirstBracket[1:indexClosingBracketFound]
		}
	}
	return ""
}
