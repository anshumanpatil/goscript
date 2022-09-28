package main

import (
	"fmt"

	"github.com/anshumanpatil/goscript"
)

func main() {
	var str goscript.JSString
	str = "anshuman"
	is, res := str.IsPalindrome()
	fmt.Println("IsPalindrome -- ", is, res)
	goscript.ColorPrint([]any{"IsPalindrome -- ", is, res})
	goscript.ColorPrint([]any{
		"StringRemove -- ",
		str.StringRemove("hu"),
	})
	goscript.ColorPrint([]any{
		"StringContains -- ",
		str.StringContains("ans"),
	})
	goscript.ColorPrint([]any{
		"StringIndexOf -- ",
		str.StringIndexOf("n"),
	})
	goscript.ColorPrint([]any{
		"StringLastIndexOf -- ",
		str.StringLastIndexOf("a"),
	})
	goscript.ColorPrint([]any{
		"StringToMap -- ",
		str.StringToMap(),
	})
	goscript.ColorPrint([]any{
		"StringToSlice -- ",
		str.StringToSlice(),
	})
	goscript.ColorPrint([]any{
		"StringSplitWithChar -- ",
		str.StringSplitWithChar("a"),
	})
	goscript.ColorPrint([]any{
		"StringChunks -- ",
		str.StringChunks(3),
	})
	goscript.ColorPrint([]any{
		"StringRemoveFirstChar -- ",
		str.StringRemoveFirstChar(),
	})
	goscript.ColorPrint([]any{
		"StringRemoveLastChar -- ",
		str.StringRemoveLastChar(),
	})
	str = "anshuman-is-very-good-boy-to-answer"
	goscript.ColorPrint([]any{
		"StringSplit -- ",
		str.StringSplit("ans"),
	})
	// fmt.Println("StringSplit -- ", goscript.StringSplit(string(str), "shum"))
}
