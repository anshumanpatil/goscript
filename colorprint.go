package goscript

import (
	"fmt"
)

type ColorPrinter int

// **** **** **** **** **** **** **** BACKGROUND COLORS **** **** **** **** **** **** //
func (c ColorPrinter) BGGray() (back ColorPrinter) {
	fmt.Print("\033[1;40m")
	return c
}
func (c ColorPrinter) BGRed() (back ColorPrinter) {
	fmt.Print("\033[1;41m")
	return c
}
func (c ColorPrinter) BGGreen() (back ColorPrinter) {
	fmt.Print("\033[1;42m")
	return c
}
func (c ColorPrinter) BGYellow() (back ColorPrinter) {
	fmt.Print("\033[1;43m")
	return c
}
func (c ColorPrinter) BGBlue() (back ColorPrinter) {
	fmt.Print("\033[1;44m")
	return c
}
func (c ColorPrinter) BGPurple() (back ColorPrinter) {
	fmt.Print("\033[1;45m")
	return c
}
func (c ColorPrinter) BGCyan() (back ColorPrinter) {
	fmt.Print("\033[1;46m")
	return c
}
func (c ColorPrinter) BGWhite() (back ColorPrinter) {
	fmt.Print("\033[1;47m")
	return c
}

// **** **** **** **** **** **** **** EFFECTS **** **** **** **** **** **** //

func (c ColorPrinter) Bold() (back ColorPrinter) {
	fmt.Print("\033[1m")
	return c
}
func (c ColorPrinter) Underlined() (back ColorPrinter) {
	fmt.Print("\033[4m")
	return c
}
func (c ColorPrinter) Blinking() (back ColorPrinter) {
	fmt.Print("\033[5m")
	return c
}
func (c ColorPrinter) Reverse() (back ColorPrinter) {
	fmt.Print("\033[7m")
	return c
}
func (c ColorPrinter) Clear() (back ColorPrinter) {
	fmt.Print("\033[H\033[2J")
	return c
}

// **** **** **** **** **** **** **** FOREGROUND COLORS **** **** **** **** **** **** //
func (c ColorPrinter) Gray(str ...interface{}) (back ColorPrinter) {
	printMatter := ""

	for _, v := range str {
		printMatter += fmt.Sprintf(fmt.Sprintf(("\033[1;30m%s\033[0m"), fmt.Sprintf("%v", v)))
	}
	fmt.Print(printMatter)
	return c
}
func (c ColorPrinter) Red(str ...interface{}) (back ColorPrinter) {
	printMatter := ""

	for _, v := range str {
		printMatter += fmt.Sprintf(fmt.Sprintf(("\033[1;31m%s\033[0m"), fmt.Sprintf("%v", v)))
	}
	fmt.Print(printMatter)
	return c
}
func (c ColorPrinter) Green(str ...interface{}) (back ColorPrinter) {
	printMatter := ""

	for _, v := range str {
		printMatter += fmt.Sprintf(fmt.Sprintf(("\033[1;32m%s\033[0m"), fmt.Sprintf("%v", v)))
	}
	fmt.Print(printMatter)
	return c
}
func (c ColorPrinter) Yellow(str ...interface{}) (back ColorPrinter) {
	printMatter := ""

	for _, v := range str {
		printMatter += fmt.Sprintf(fmt.Sprintf(("\033[1;33m%s\033[0m"), fmt.Sprintf("%v", v)))
	}
	fmt.Print(printMatter)
	return c
}
func (c ColorPrinter) Blue(str ...interface{}) (back ColorPrinter) {
	printMatter := ""

	for _, v := range str {
		printMatter += fmt.Sprintf(fmt.Sprintf(("\033[1;34m%s\033[0m"), fmt.Sprintf("%v", v)))
	}
	fmt.Print(printMatter)
	return c
}
func (c ColorPrinter) Purple(str ...interface{}) (back ColorPrinter) {
	printMatter := ""

	for _, v := range str {
		printMatter += fmt.Sprintf(fmt.Sprintf(("\033[1;35m%s\033[0m"), fmt.Sprintf("%v", v)))
	}
	fmt.Print(printMatter)
	return c
}
func (c ColorPrinter) Cyan(str ...interface{}) (back ColorPrinter) {
	printMatter := ""

	for _, v := range str {
		printMatter += fmt.Sprintf(fmt.Sprintf(("\033[1;36m%s\033[0m"), fmt.Sprintf("%v", v)))
	}
	fmt.Print(printMatter)
	return c
}
func (c ColorPrinter) White(str ...interface{}) (back ColorPrinter) {
	printMatter := ""

	for _, v := range str {
		printMatter += fmt.Sprintf(fmt.Sprintf(("\033[1;37m%s\033[0m"), fmt.Sprintf("%v", v)))
	}
	fmt.Print(printMatter)
	return c
}
