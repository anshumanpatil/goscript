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

func Reverse(str string) (result string) {
	strArr := []rune(str)
	for i := len(strArr) - 1; i >= 0; i-- {
		result += string(strArr[i])
	}
	return
}

func (str JSString) IndexOf(strToCompare string) (result int) {
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

func (str JSString) LastIndexOf(strToCompare string) (result int) {
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

func (str JSString) ToMap() map[int]string {
	strArr := []rune(string(str))
	var result map[int]string = map[int]string{}
	for i := 0; i <= len(strArr)-1; i++ {
		result[i] = string(strArr[i])
	}

	return result
}

func (str JSString) ToSlice() []string {
	strArr := []rune(string(str))
	result := []string{}
	for i := 0; i <= len(strArr)-1; i++ {
		result = append(result, string(strArr[i]))
	}

	return result
}

func (str JSString) RemoveFirstChar() string {
	strArr := []rune(string(str))
	result := ""
	for i := 1; i <= len(strArr)-1; i++ {
		result += string(strArr[i])
	}

	return result
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

func IntSliceHas(s []int, index int) (result bool) {
	for _, v := range s {
		if v == index {
			result = true
		}
	}
	return
}

// func (str JSString) Remove(occurer JSString) (result JSString) {
// 	occurerLength := len(occurer)
// 	strArr := str.ToSlice()
// 	indexes := []int{}
// 	resultAppend := []string{}
// 	for i, v := range strArr {
// 		if len(strArr[i:i+occurerLength]) >= occurerLength {
// 			comparer := ""
// 			for _, vc := range strArr[i : i+occurerLength] {
// 				comparer += vc
// 			}
// 			if comparer == string(occurer) {
// 				indexes = append(indexes, i)
// 				resultAppend = append(resultAppend, "remove")
// 			} else {
// 				resultAppend = append(resultAppend, v)
// 			}
// 		}
// 	}
// 	for _, v := range resultAppend {
// 		if v != "remove" {
// 			result += JSString(v)
// 		}
// 	}
// 	return
// }

func (str JSString) Contains(occurer string) (result int) {
	strArr := []rune(string(str))
	myresult := []string{}
	var p JSString = ""
	for i := 0; i <= len(strArr)-1; i++ {
		p += JSString(strArr[i])
		if len(p) > len((occurer)) {
			p.RemoveFirstChar()
		}
		if len(p) == len((occurer)) && string(p) == occurer {
			myresult = append(myresult, string(p))
			result++
		}
	}
	return
}

func (str JSString) ParseInt() (result int) {
	rc := []rune(str)
	for i, v := range rc {
		recs := 0
		dec := ((i - len(rc)) * -1)
		ver := 1
		for i := 1; i < dec; i++ {
			ver = ver * 10
		}
		switch int(v) {
		case 48:
			recs = 0
		case 49:
			recs = 1
		case 50:
			recs = 2
		case 51:
			recs = 3
		case 52:
			recs = 4
		case 53:
			recs = 5
		case 54:
			recs = 6
		case 55:
			recs = 7
		case 56:
			recs = 8
		case 57:
			recs = 9
		}
		result += recs * ver
	}
	return
}

func (str JSString) Split(occurer JSString) (myresult []JSString) {
	strCopy := str
	isAvailableAt := -1
	freshValues := []string{}
	resultStart := ""
	resultEnd := ""
	result := ""
	for {
		myresult = str.Chunks(len(occurer))
		for i, v := range myresult {
			if v == occurer {
				isAvailableAt = i

				break
			}
		}
		if isAvailableAt >= 0 || len(str) == 0 {
			break
		}
		freshValues = append(freshValues, (str.ToSlice())[0])
		str = JSString(str.RemoveFirstChar())
	}
	if isAvailableAt == -1 { // occurer not available in string
		fmt.Println("isAvailableAt", isAvailableAt, str)
		myresult = []JSString{strCopy}
		return myresult
	}
	for _, v := range myresult[0:isAvailableAt] {
		resultStart += string(v)
	}
	fmt.Println(myresult, isAvailableAt)
	for _, v := range myresult[isAvailableAt+1:] {
		resultEnd += string(v)
	}
	for _, v := range freshValues {
		result += string(v)
	}
	myresult = []JSString{}
	myresult = append(myresult, JSString(string(result)+string(resultStart)), JSString(resultEnd))
	fmt.Println(" result : ", result)
	fmt.Println(" resultStart : ", resultStart)
	fmt.Println(" resultEnd : ", resultEnd)
	fmt.Println(" freshValues : ", freshValues)
	fmt.Println(" isAvailableAt : ", isAvailableAt)
	fmt.Println(" strCopy : ", strCopy)
	fmt.Println(" occurer : ", occurer)
	return
}

func (str JSString) SplitWithChar(splitter JSString) []string {
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

func (str JSString) Dummy(occurer JSString) (n int) {
	occurerLength := len(occurer)
	strArr := str.ToSlice()
	type occurerIndex struct {
		from int
		to   int
	}
	// indexes := []occurerIndex{}
	result := []string{}
	foo := occurerIndex{}
	for i, v := range strArr {

		if len(strArr[i:i+occurerLength]) >= occurerLength {
			comparer := ""
			for _, vc := range strArr[i : i+occurerLength] {
				comparer += vc
			}
			if comparer == string(occurer) {
				// indexes = append(indexes, occurerIndex{from: i, to: i + occurerLength})
				foo.from = i
				foo.to = i + occurerLength
				result = append(result, "removeStart")
				result = append(result, v)
			} else {
				result = append(result, v)
			}
		}
		if foo.from >= 0 && foo.to >= 0 && (foo.to-1) == i {
			result = append(result, "removeEnd")
		}
	}
	temp := ""
	allTemps := []string{}
	for _, v := range result {
		if v == "remove" {
			if len(temp) > 0 {
				allTemps = append(allTemps, temp)
			}
			temp = ""
		} else {
			temp += v
		}
	}
	if len(temp) > 0 {
		allTemps = append(allTemps, temp)
	}
	for _, v := range allTemps {
		fmt.Println("allTemps   ", v)

	}
	// fmt.Println(indexes)
	fmt.Println(result)
	// fmt.Println(indexes)
	return

}
