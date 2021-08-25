package main

import (
	"fmt"
	"strings"
)

// this function to find all string inside the bracket
func findAllStringInBracket(str string) string {
	indexOpen := strings.Index(str, "(")
	if indexOpen >= 0 {
		indexClose := strings.Index(str, ")")
		if indexClose >= 0 {
			return str[indexOpen+1 : indexClose]
		}
	}
	return ""
}

// this function to find first string inside the bracket
func findFirstStringInBracket(str string) string {
	indexOpen := strings.Index(str, "(")
	if indexOpen >= 0 {
		indexClose := strings.Index(str, ")")
		firstSpace := strings.Index(str, " ")
		if indexClose >= 0 {
			return str[indexOpen+1 : firstSpace]
		}
	}
	return ""
}

func main() {
	fmt.Println(findAllStringInBracket("(alvin lesmana)"))
	fmt.Println(findFirstStringInBracket("(alvin lesmana)"))
}
