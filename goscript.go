package goscript

// https://medium.com/@BastianRob/implementing-reduce-in-go-4a3e6e3affc
import (
	"fmt"
	"reflect"
)

type JSString string

// type ARJS []interface{}

func (source JSString) ForEach(calle func(uint8, interface{}) JSString) (result []JSString) {
	calleFunc := reflect.ValueOf(calle)

	for i := 0; i < len(source); i++ {
		print := calleFunc.Call([]reflect.Value{
			reflect.ValueOf(source[i]),
			reflect.ValueOf(i),
		})
		result = append(result, JSString(fmt.Sprint(print[0])))
	}

	return
}

func anyTrue(src []int, occurer int) (result bool) {
	for _, idx := range src {
		if idx == occurer {
			return true
		}
	}
	return
}
func ColorPrint(str ...interface{}) {
	printMatter := ""
	if len(str) == 1 {
		printMatter += fmt.Sprintf(fmt.Sprintf((" \033[1;32m%s\033[0m "), fmt.Sprintf("%v", str[0])))
		fmt.Println(printMatter)
		return
	}

	for _, v := range str {
		// if i == 0 {
		// 	printMatter += fmt.Sprintf(fmt.Sprintf((" \033[1;32m%s\033[0m "), fmt.Sprintf("%v", v)))
		// } else {
		printMatter += fmt.Sprintf(fmt.Sprintf((" \033[1;35m%s\033[0m "), fmt.Sprintf("%v", v)))
		// }
	}
	fmt.Println(printMatter)
}

func (str JSString) Reverse() (result JSString) {
	strArr := []rune(str)
	for i := len(strArr) - 1; i >= 0; i-- {
		result += JSString(strArr[i])
	}
	return
}

func (str JSString) IsPalindrome() (ok bool) {
	return str == str.Reverse()
}

func (str JSString) Contains(occurer string) (result int) {
	strArr := []rune(string(str))
	myresult := []string{}
	var p JSString = ""
	for i := 0; i <= len(strArr)-1; i++ {
		p += JSString(strArr[i])
		if len(p) > len((occurer)) {
			p = p.RemoveFirstChar()
		}
		if len(p) == len((occurer)) && string(p) == occurer {
			myresult = append(myresult, string(p))
			result++
		}
	}
	return
}

func (str JSString) Replace(occurer string, replacer string) (result JSString) {
	strArr := []rune(string(str))
	myresult := []string{}
	resultIndex := []int{}
	var p JSString = ""
	for i := 0; i <= len(strArr)-1; i++ {
		p += JSString(strArr[i])
		if len(p) > len((occurer)) {
			p = p.RemoveFirstChar()
		}
		if len(p) == len((occurer)) && string(p) == occurer {
			myresult = append(myresult, string(p))
			resultIndex = append(resultIndex, i-len(occurer)+1)
			// return
		}
	}
	nums := []int{}
	for _, idx := range resultIndex {
		for i := idx; i < (idx + len(occurer)); i++ {
			nums = append(nums, i)
		}
	}
	for i, v := range str {
		if anyTrue(nums, i) == false {
			result += JSString(v)
		}
		if anyTrue(resultIndex, i) {
			result += JSString(replacer)
		}
	}
	return
}

func (str JSString) Split(occurer string) (betaResult []JSString) {
	strArr := []rune(string(str))
	myresult := []string{}
	resultIndex := []int{}
	var p JSString = ""
	for i := 0; i <= len(strArr)-1; i++ {
		p += JSString(strArr[i])
		if len(p) > len((occurer)) {
			p = p.RemoveFirstChar()
		}
		if len(p) == len((occurer)) && string(p) == occurer {
			myresult = append(myresult, string(p))
			resultIndex = append(resultIndex, i-len(occurer)+1)
		}
	}
	nums := []int{}
	result := ""
	for _, idx := range resultIndex {
		for i := idx; i < (idx + len(occurer)); i++ {
			nums = append(nums, i)
		}
	}
	for i, v := range str {
		if anyTrue(nums, i) == false {
			result += string(v)
		}
		if anyTrue(resultIndex, i) {
			betaResult = append(betaResult, JSString(result))
			result = ""
		}
	}

	return
}

func (str JSString) IndexOf(occurer string) (result int) {
	strArr := []rune(string(str))
	myresult := []string{}
	var p JSString = ""
	for i := 0; i <= len(strArr)-1; i++ {
		p += JSString(strArr[i])
		if len(p) > len((occurer)) {
			p = p.RemoveFirstChar()
		}
		if len(p) == len((occurer)) && string(p) == occurer {
			myresult = append(myresult, string(p))
			result = i - len(occurer) + 1
			return
		}
	}
	return
}

func (str JSString) LastIndexOf(occurer string) (result int) {
	strArr := []rune(string(str))
	myresult := []string{}
	var p JSString = ""
	for i := 0; i <= len(strArr)-1; i++ {
		p += JSString(strArr[i])
		if len(p) > len((occurer)) {
			p = p.RemoveFirstChar()
		}
		if len(p) == len((occurer)) && string(p) == occurer {
			myresult = append(myresult, string(p))
			result = i - len(occurer) + 1
		}
	}
	return
}

func (str JSString) ToSlice() []JSString {
	strArr := []rune(string(str))
	result := []JSString{}
	for i := 0; i <= len(strArr)-1; i++ {
		result = append(result, JSString(strArr[i]))
	}

	return result
}

func (str JSString) RemoveFirstChar() JSString {
	strArr := []rune(string(str))
	result := ""
	for i := 1; i <= len(strArr)-1; i++ {
		result += string(strArr[i])
	}

	return JSString(result)
}

func (str JSString) RemoveLastChar() string {
	strArr := []rune(string(str))
	result := ""
	for i := 0; i <= len(strArr)-1; i++ {
		if i != len(strArr)-1 {
			result += string(strArr[i])
		}
	}

	return result
}

func (str JSString) Chunks(splitter int) []JSString {
	strArr := []rune(string(str))

	parts := len(strArr) % splitter

	if parts == 0 {
		parts = len(strArr) / splitter
	} else {
		parts = (len(strArr) - parts) / splitter
	}
	myresult := []JSString{}

	p := ""
	for i := 0; i <= len(strArr)-1; i++ {
		// fmt.Println("index - ", i, " string - ", string(strArr[i]), " p -- ", p)
		p += string(strArr[i])
		if len(p) == splitter {
			myresult = append(myresult, JSString(p))
			p = ""
		}

	}
	myresult = append(myresult, JSString(p))

	return myresult
}

// func (str JSString) ParseInt() (result int) {
// 	rc := []rune(str)
// 	for i, v := range rc {
// 		recs := 0
// 		dec := ((i - len(rc)) * -1)
// 		ver := 1
// 		for i := 1; i < dec; i++ {
// 			ver = ver * 10
// 		}
// 		switch int(v) {
// 		case 48:
// 			recs = 0
// 		case 49:
// 			recs = 1
// 		case 50:
// 			recs = 2
// 		case 51:
// 			recs = 3
// 		case 52:
// 			recs = 4
// 		case 53:
// 			recs = 5
// 		case 54:
// 			recs = 6
// 		case 55:
// 			recs = 7
// 		case 56:
// 			recs = 8
// 		case 57:
// 			recs = 9
// 		}
// 		result += recs * ver
// 	}
// 	return
// }
