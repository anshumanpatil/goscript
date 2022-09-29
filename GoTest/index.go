package main

import (
	"github.com/anshumanpatil/goscript"
)

func main() {

	// var str goscript.JSString
	str := goscript.JSString("anshuman is good man")
	// is, res := str.IsPalindrome()
	// goscript.ColorPrint([]any{"IsPalindrome -- ", is, res})

	goscript.ColorPrint([]any{
		"String Contains -- ",
		str,
		"an :=== ",

		str.Split("ans"),
	})
	// goscript.ColorPrint([]any{
	// 	"String IndexOf -- ",
	// 	str.IndexOf("a"),
	// })
	// goscript.ColorPrint([]any{
	// 	"String LastIndexOf -- ",
	// 	str.LastIndexOf("a"),
	// })
	// goscript.ColorPrint([]any{
	// 	"String ToMap -- ",
	// 	str.ToMap(),
	// })
	// goscript.ColorPrint([]any{
	// 	"String ToSlice -- ",
	// 	str.ToSlice(),
	// })
	// goscript.ColorPrint([]any{
	// 	"String SplitWithChar -- ",
	// 	str.SplitWithChar("a"),
	// })
	// goscript.ColorPrint([]any{
	// 	"String Chunks -- ",
	// 	str.Chunks(3),
	// })
	// goscript.ColorPrint([]any{
	// 	"String RemoveFirstChar -- ",
	// 	str.RemoveFirstChar(),
	// })
	// goscript.ColorPrint([]any{
	// 	"String RemoveLastChar -- ",
	// 	str.RemoveLastChar(),
	// })
	// str = "anshuman-is-very-good-booy-to-answer"
	// goscript.ColorPrint([]any{
	// 	"String Split -- ",
	// 	str.Split("ans"),
	// })
	// goscript.ColorPrint([]any{
	// 	"Dummy -- ",
	// 	str.Dummy("oo"),
	// })

	// fmt.Println("StringSplit -- ", goscript.StringSplit(string(str), "shum"))
}
