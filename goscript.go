package goscript

import "fmt"

type JSString string

func ColorPrint(str []interface{}) {
	printMatter := ""
	for i, v := range str {
		if i == 0 {
			printMatter += fmt.Sprintf(fmt.Sprintf((" \033[1;32m%s\033[0m "), fmt.Sprintf("%v", v)))
		} else {
			printMatter += fmt.Sprintf(fmt.Sprintf((" \033[1;35m%s\033[0m "), fmt.Sprintf("%v", v)))
		}
	}
	fmt.Println(printMatter)
}

func (str JSString) IsPalindrome() (ok bool, result JSString) {
	strArr := []rune(string(str))
	for i := len(strArr) - 1; i >= 0; i-- {
		result += JSString(strArr[i])
	}
	return str == result, result
}

func StringReverse(str string) (result string) {
	strArr := []rune(str)
	for i := len(strArr) - 1; i >= 0; i-- {
		result += string(strArr[i])
	}
	return
}

func (str JSString) StringIndexOf(strToCompare string) (result int) {
	if len(strToCompare) > 1 {
		ColorPrint([]any{"[StringIndexOf] Only works for single character!!!"})
		return -1
	}
	strArr := []rune(string(str))
	for i := 0; i <= len(strArr)-1; i++ {
		if string(strArr[i]) == strToCompare {
			return i
		}
	}
	if result == 0 {
		result = -1
	}
	return
}

func (str JSString) StringLastIndexOf(strToCompare string) (result int) {
	if len(strToCompare) > 1 {
		ColorPrint([]any{"[StringIndexOf] Only works for single character!!!"})
		return -1
	}
	strArr := []rune(string(str))
	for i := len(strArr) - 1; i >= 0; i-- {
		if string(strArr[i]) == strToCompare {
			return i
		}
	}
	if result == 0 {
		result = -1
	}
	return
}

func (str JSString) StringToMap() map[int]string {
	strArr := []rune(string(str))
	var result map[int]string = map[int]string{}
	for i := 0; i <= len(strArr)-1; i++ {
		result[i] = string(strArr[i])
	}

	return result
}

func (str JSString) StringToSlice() []string {
	strArr := []rune(string(str))
	result := []string{}
	for i := 0; i <= len(strArr)-1; i++ {
		result = append(result, string(strArr[i]))
	}

	return result
}

func (str JSString) StringRemoveFirstChar() string {
	strArr := []rune(string(str))
	result := ""
	for i := 1; i <= len(strArr)-1; i++ {
		result += string(strArr[i])
	}

	return result
}

func (str JSString) StringRemoveLastChar() string {
	strArr := []rune(string(str))
	result := ""
	for i := 0; i <= len(strArr)-1; i++ {
		if i != len(strArr)-1 {
			result += string(strArr[i])
		}
	}

	return result
}

func IntSliceHas(s []int, index int) (result bool) {
	for _, v := range s {
		if v == index {
			result = true
		}
	}
	return
}

func (str JSString) StringRemove(occurer JSString) (result JSString) {
	myresult := []int{}
	strarr := str.StringToSlice()
	occurerresult := occurer.StringToSlice()
	for i, v := range strarr {
		for _, vx := range occurerresult {
			if v == vx {
				myresult = append(myresult, i)
			}
		}
	}
	for i, v := range strarr {
		if !IntSliceHas(myresult, i) {
			result += JSString(v)
		}

	}
	return
}

func (str JSString) StringContains(occurer string) (result int) {
	strArr := []rune(string(str))
	myresult := []string{}
	var p JSString = ""
	for i := 0; i <= len(strArr)-1; i++ {
		p += JSString(strArr[i])
		if len(p) > len((occurer)) {
			p.StringRemoveFirstChar()
		}
		if len(p) == len((occurer)) && string(p) == occurer {
			myresult = append(myresult, string(p))
			result++
		}
	}
	return
}

// func (str JSString) StringSplit(occurer JSString) (myresult []JSString) {
// 	strCopy := str
// 	isAvailableAt := -1
// 	freshValues := []string{}
// 	resultStart := ""
// 	resultEnd := ""
// 	result := ""
// 	for {
// 		myresult = str.StringChunks(len(occurer))
// 		for i, v := range myresult {
// 			if v == occurer {
// 				isAvailableAt = i

// 				break
// 			}
// 		}
// 		if isAvailableAt >= 0 || len(str) == 0 {
// 			break
// 		}
// 		freshValues = append(freshValues, (str.StringToSlice())[0])
// 		str = JSString(str.StringRemoveFirstChar())
// 	}
// 	if isAvailableAt == -1 { // occurer not available in string
// 		fmt.Println("isAvailableAt", isAvailableAt, str)
// 		myresult = []JSString{strCopy}
// 		return myresult
// 	}
// 	for _, v := range myresult[0:isAvailableAt] {
// 		resultStart += string(v)
// 	}
// 	for _, v := range myresult[isAvailableAt+len(occurer):] {
// 		resultEnd += string(v)
// 	}
// 	for _, v := range freshValues {
// 		result += string(v)
// 	}
// 	myresult = []JSString{}
// 	myresult = append(myresult, JSString(string(result)+string(resultStart)), JSString(resultEnd))
// 	fmt.Println(" result : ", result)
// 	fmt.Println(" resultStart : ", resultStart)
// 	fmt.Println(" resultEnd : ", resultEnd)
// 	fmt.Println(" freshValues : ", freshValues)
// 	fmt.Println(" isAvailableAt : ", isAvailableAt)
// 	fmt.Println(" strCopy : ", strCopy)
// 	fmt.Println(" occurer : ", occurer)
// 	return
// }

func (str JSString) StringSplitWithChar(splitter JSString) []string {
	// str += splitter
	strArr := []rune(string(str))
	myresult := []string{}
	p := ""
	for i := 0; i <= len(strArr)-1; i++ {
		// fmt.Println("index - ", i, " string - ", string(strArr[i]), " p -- ", p, " comparison ", string(strArr[i]) == splitter)
		if JSString(strArr[i]) == splitter {
			myresult = append(myresult, p)
			p = ""
		} else {
			p += string(strArr[i])
		}
	}
	myresult = append(myresult, p)

	return myresult
}

func (str JSString) StringChunks(splitter int) []JSString {
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
