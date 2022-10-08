package main

import (
	"fmt"

	"github.com/anshumanpatil/goscript"
)

func main() {

	var str goscript.JSString

	// IsPalindrome Example
	str = "anshuman"
	goscript.ColorPrint("\nIsPalindrome Example : ")
	goscript.ColorPrint("IsPalindrome: ", str, str.IsPalindrome())
	str = "ana"
	goscript.ColorPrint("IsPalindrome: ", str, str.IsPalindrome())
	goscript.ColorPrint("-----------------------------------------------------------\n")

	// Reverse Example
	str = "anshuman"
	goscript.ColorPrint("\nReverse Example : ")
	goscript.ColorPrint("Reverse: ", str, str.Reverse())
	goscript.ColorPrint("-----------------------------------------------------------\n")

	// Contains Example
	str = "anshuman is in pune and he is there and"
	goscript.ColorPrint("\nContains Example : ")
	goscript.ColorPrint("Contains: ", str, str.Contains("n"))
	goscript.ColorPrint("-----------------------------------------------------------\n")

	// IndexOf Example
	str = "anshuman is in pune"
	goscript.ColorPrint("\nIndexOf Example : ")
	goscript.ColorPrint("IndexOf: ", str, str.IndexOf("ans"))
	// for i, v := range str {
	// 	fmt.Println(i, string(v))
	// }
	goscript.ColorPrint("-----------------------------------------------------------\n")

	// LastIndexOf Example
	str = "anshuman is in pune pune pune"
	goscript.ColorPrint("\nLastIndexOf Example : ")
	goscript.ColorPrint("LastIndexOf: ", str, str.LastIndexOf("pun"))
	// for i, v := range str {
	// 	fmt.Println(i, string(v))
	// }
	goscript.ColorPrint("-----------------------------------------------------------\n")

	// ToSlice Example
	str = "anshuman"
	goscript.ColorPrint("\nToSlice Example : ")
	goscript.ColorPrint("ToSlice: ", str, str.ToSlice())
	goscript.ColorPrint("-----------------------------------------------------------\n")

	// RemoveFirstChar Example
	str = "anshuman"
	goscript.ColorPrint("\nRemoveFirstChar Example : ")
	goscript.ColorPrint("RemoveFirstChar: ", str, str.RemoveFirstChar())
	goscript.ColorPrint("-----------------------------------------------------------\n")

	// RemoveLastChar Example
	str = "anshuman"
	goscript.ColorPrint("\nRemoveLastChar Example : ")
	goscript.ColorPrint("RemoveLastChar: ", str, str.RemoveLastChar())
	goscript.ColorPrint("-----------------------------------------------------------\n")

	// Chunks Example
	str = "anshuman"
	goscript.ColorPrint("\nChunks Example : ")
	goscript.ColorPrint("Chunks: ", str, str.Chunks(4))
	goscript.ColorPrint("-----------------------------------------------------------\n")

	// time.Sleep(2 * time.Second)
	// time.Sleep(2 * time.Second)
	// time.Sleep(2 * time.Second)
	// time.Sleep(2 * time.Second)
	// time.Sleep(2 * time.Second)
	// time.Sleep(2 * time.Second)
	// time.Sleep(2 * time.Second)
	// time.Sleep(2 * time.Second)
	// time.Sleep(2 * time.Second)

	// Replace Example
	str = "anshuman is in pune and mumbai pune pune punpun"
	goscript.ColorPrint("\nReplace Example : ")
	goscript.ColorPrint("Replace: ", str, str.Replace("pun", ""))
	goscript.ColorPrint("-----------------------------------------------------------\n")

	// Split Example
	str = "anshuman is in pune and mumbai pune pune punpun"
	goscript.ColorPrint("\nSplit Example : ")
	goscript.ColorPrint("Split: ", str, str.Split("pun"))
	goscript.ColorPrint("-----------------------------------------------------------\n")

	goscript.ColorPrinter(0).Clear()

	// ForEach Example
	str = "anshuman"
	goscript.ColorPrint("\nForEach Example : ")
	// goscript.ColorPrint("ForEach: ", str, )

	bing := str.ForEach(func(e uint8, i interface{}) goscript.JSString {
		return goscript.JSString("-" + string(e))
	})
	fmt.Println(bing)
	goscript.ColorPrint("-----------------------------------------------------------\n")

	goscript.ColorPrinter(0).
		Red("Red Print").
		BGRed().
		White("BG Dark Gray").
		Blinking().
		Red("anshu").
		Green(bing, "Green\n")
	// Puma([]interface{}{1, 2, "anshuman"}).MAP(func(elem, idx interface{}) {
	// 	fmt.Println(elem, idx)
	// })
	// // fmt.Println("StringSplit -- ", goscript.StringSplit(string(str), "shum"))
	// fmt.Println(str.ToSlice())
	// arr := []interface{}{"anshu", 2, 3, 4, 5}
	// // str.ToSlice()
	// goscript.ARJS(arr).MapLikeJS(func(entry interface{}, idx int) {
	// 	fmt.Println(entry, idx)
	// })

}
