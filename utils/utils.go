package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"unicode"
	"unicode/utf8"
)

func IsDigitString(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func PrintLine(tasks []string) {
	max := 27

	mxNum := len(tasks)
	for _, task := range tasks {
		if mxNum+utf8.RuneCountInString(task)+6 > max {
			max = mxNum + utf8.RuneCountInString(task) + 6
		}
	}

	line := ""
	for range max {
		line += "-"
	}
	fmt.Println(line)
}
