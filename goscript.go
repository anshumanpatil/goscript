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

func StringContains(str string, strToCompare string) (result bool) {
	strArr := []rune(str)
	for i := len(strArr) - 1; i >= 0; i-- {
		if string(strArr[i]) == strToCompare {
			return true
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

func StringSplit(str string, splitter string) []string {
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
