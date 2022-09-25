package main

import (
	"fmt"

	"github.com/anshumanpatil/goscript"
)

func main() {
	str := "anshuman"
	is, res := goscript.IsPalindrome(str)
	fmt.Println("IsPalindrome -- ", is, res)
	fmt.Println("StringContainsChar -- ", goscript.StringContainsChar(str, "a"))
	fmt.Println("StringRemove -- ", goscript.StringRemove(str, "hu"))
	fmt.Println("StringContains -- ", goscript.StringContains(str, "hu"))
	fmt.Println("StringIndexOf -- ", goscript.StringIndexOf(str, "a"))
	fmt.Println("StringLastIndexOf -- ", goscript.StringLastIndexOf(str, "a"))
	fmt.Println("StringLength -- ", goscript.StringLength(str))
	fmt.Println("StringToMap -- ", goscript.StringToMap(str))
	fmt.Println("StringToSlice -- ", goscript.StringToSlice(str))
	fmt.Println("StringSplitWithChar -- ", goscript.StringSplitWithChar(str, "a"))
	fmt.Println("StringChunks -- ", goscript.StringChunks(str, 3))
	fmt.Println("StringInstances -- ", goscript.StringCharInstances(str, "a"))
	fmt.Println("StringRemoveFirstChar -- ", goscript.StringRemoveFirstChar(str))
	fmt.Println("StringRemoveLastChar -- ", goscript.StringRemoveLastChar(str))
	fmt.Println("StringSplit -- ", goscript.StringSplit(str, "shum"))
}
