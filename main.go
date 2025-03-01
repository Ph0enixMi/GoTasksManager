package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const menu = "1 - Создать задачу\n2 - Удалить задачу\n3 - Добавить/убрать отметку\n4 - Выйти\n"

func main() {
	startApp()
}

func startApp() {
	var tasks []string
	var marks []bool

	for {
		input, _ := scanChoice(tasks, marks, true)

		switch input {
		case 1:
			tasks, marks = newTask(tasks, marks)
		case 2:
			tasks, marks = deleteTask(tasks, marks)
		case 3:
			tasks, marks = markTask(tasks, marks)
		case 4:
			return
		}
	}
}

func scanChoice(tasks []string, marks []bool, start bool) (int, error) {
	reader := bufio.NewReader(os.Stdin)

	clearScreen()
	printTasks(tasks, marks)

	if start {
		fmt.Print(menu)
		printLine(tasks)
		fmt.Print("[1-4] ")
	} else {
		if tasks == nil {
			return -1, errors.New("tasks slice is nil")
		}
		if len(tasks) == 1 {
			fmt.Print("[1] ")
		} else {
			fmt.Printf("[1-%v] ", len(tasks))
		}
	}

	input, err := reader.ReadString('\n')
	if err != nil {
		return -1, errors.New("input error")
	}

	input = strings.TrimSpace(input)

	if isDigitString(input) {
		input_num, err := strconv.Atoi(input)
		if err != nil {
			return -1, errors.New("input error")
		}
		if start {
			if input_num > 0 && input_num < 6 {
				return input_num, nil
			}
		} else {
			if input_num > 0 && input_num <= len(tasks) {
				return input_num, nil
			}
		}
	}
	return -1, errors.New("некорректный ввод")
}

func isDigitString(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func printTasks(tasks []string, marks []bool) {
	if len(tasks) == 0 {
		return
	}

	for i, task := range tasks {
		if marks[i] {
			fmt.Printf("%d: [x] %s\n", i+1, task)
		} else {
			fmt.Printf("%d: [ ] %s\n", i+1, task)
		}
	}
	printLine(tasks)
}

func printLine(tasks []string) {
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

func clearScreen() {
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

func newTask(tasks []string, marks []bool) ([]string, []bool) {
	reader := bufio.NewReader(os.Stdin)

	for {
		clearScreen()
		printTasks(tasks, marks)

		if len(tasks) == 0 {
			fmt.Println("Введите текст")
		}
		fmt.Print("[т] ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		text = strings.TrimSpace(text)

		tasks = append(tasks, text)
		marks = append(marks, false)

		break
	}

	return tasks, marks
}

func deleteTask(tasks []string, marks []bool) ([]string, []bool) {
	if len(tasks) == 0 {
		return tasks, marks
	}

	index, err := scanChoice(tasks, marks, false)

	if index > 0 && index <= len(tasks) && err == nil {
		tasks = append(tasks[:index-1], tasks[index:]...)
		marks = append(marks[:index-1], marks[index:]...)
	}
	return tasks, marks
}

func markTask(tasks []string, marks []bool) ([]string, []bool) {
	if len(tasks) == 0 {
		return tasks, marks
	}

	index, err := scanChoice(tasks, marks, false)

	if index > 0 && index <= len(tasks) && err == nil {
		if marks[index-1] {
			marks[index-1] = false
		} else {
			marks[index-1] = true
		}
	}
	return tasks, marks
}
