package goscript

func IsPalindrome(str string) (ok bool, result string) {
	strArr := []rune(str)
	for i := len(strArr) - 1; i >= 0; i-- {
		result += string(strArr[i])
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

func StringContainsChar(str string, strToCompare string) (result int) {
	strArr := []rune(str)
	for i := len(strArr) - 1; i >= 0; i-- {
		if string(strArr[i]) == strToCompare {
			result++
		}
	}
	return
}

func StringIndexOf(str string, strToCompare string) (result int) {
	strArr := []rune(str)
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

func StringLastIndexOf(str string, strToCompare string) (result int) {
	strArr := []rune(str)
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

func StringLength(str string) (length int) {
	length = len([]rune(str))
	return
}

func StringToMap(str string) map[int]string {
	strArr := []rune(str)
	var result map[int]string = map[int]string{}
	for i := 0; i <= len(strArr)-1; i++ {
		result[i] = string(strArr[i])
	}

	return result
}

func StringToSlice(str string) []string {
	strArr := []rune(str)
	result := []string{}
	for i := 0; i <= len(strArr)-1; i++ {
		result = append(result, string(strArr[i]))
	}

	return result
}

func StringRemoveFirstChar(str string) string {
	strArr := []rune(str)
	result := ""
	for i := 1; i <= len(strArr)-1; i++ {
		result += string(strArr[i])
	}

	return result
}

func StringRemoveLastChar(str string) string {
	strArr := []rune(str)
	result := ""
	for i := 0; i <= len(strArr)-1; i++ {
		if i != len(strArr)-1 {
			result += string(strArr[i])
		}
	}

	return result
}

func StringCharInstances(str string, occurer string) (result int) {
	if StringLength(occurer) > 1 {
		return 0
	}
	strArr := []rune(str)
	for i := 0; i <= len(strArr)-1; i++ {
		if string(strArr[i]) == occurer {
			result += 1
		}
	}

	return
}

func IntSliceHas(s []int, index int) (result bool) {
	for _, v := range s {
		if v == index {
			result = true
		}
	}
	return
}

func StringRemove(str string, occurer string) (result string) {
	myresult := []int{}
	strarr := StringToSlice(str)
	occurerresult := StringToSlice(occurer)
	for i, v := range strarr {
		for _, vx := range occurerresult {
			if v == vx {
				myresult = append(myresult, i)
			}
		}
	}
	for i, v := range strarr {
		if !IntSliceHas(myresult, i) {
			result += v
		}

	}
	return
}

func StringContains(str string, occurer string) (result int) {
	strArr := []rune(str)
	myresult := []string{}
	p := ""
	for i := 0; i <= len(strArr)-1; i++ {
		p += string(strArr[i])
		if StringLength(p) > StringLength((occurer)) {
			p = StringRemoveFirstChar(p)
		}
		if StringLength(p) == StringLength((occurer)) && p == occurer {
			myresult = append(myresult, p)
			result++
		}
	}
	return
}

func StringSplit(str string, occurer string) (myresult []string) {
	strArr := []rune(str)
	// myresult := []string{}
	p := ""
	t := ""
	for i := 0; i <= len(strArr)-1; i++ {
		p += string(strArr[i])
		t += string(strArr[i])
		if StringLength(p) > StringLength((occurer)) {
			p = StringRemoveFirstChar(p)
		}
		if StringLength(p) == StringLength((occurer)) && p == occurer {
			myresult = append(myresult, StringRemove(t, occurer))
			t = ""
			// result++
		}
	}
	myresult = append(myresult, StringRemove(t, occurer))
	return
}

func StringSplitWithChar(str string, splitter string) []string {
	// str += splitter
	strArr := []rune(str)
	myresult := []string{}
	p := ""
	for i := 0; i <= len(strArr)-1; i++ {
		// fmt.Println("index - ", i, " string - ", string(strArr[i]), " p -- ", p, " comparison ", string(strArr[i]) == splitter)
		if string(strArr[i]) == splitter {
			myresult = append(myresult, p)
			p = ""
		} else {
			p += string(strArr[i])
		}
	}
	myresult = append(myresult, p)

	return myresult
}

func StringChunks(str string, splitter int) []string {
	strArr := []rune(str)

	parts := len(strArr) % splitter

	if parts == 0 {
		parts = len(strArr) / splitter
	} else {
		parts = (len(strArr) - parts) / splitter
	}
	myresult := []string{}

	p := ""
	for i := 0; i <= len(strArr)-1; i++ {
		// fmt.Println("index - ", i, " string - ", string(strArr[i]), " p -- ", p)
		p += string(strArr[i])
		if StringLength(p) == splitter {
			myresult = append(myresult, p)
			p = ""
		}

	}
	myresult = append(myresult, p)

	return myresult
}
