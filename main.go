package main

// import (
// 	"bufio"
// 	"errors"
// 	"fmt"
// 	"os"
// 	"slices"
// 	"strconv"
// 	"strings"
// 	"tasks/utils"
// )

// const menu = "1 - Создать задачу\n2 - Удалить задачу\n3 - Добавить/убрать отметку\n4 - Выйти\n"

// func main() {
// 	startApp()
// }

// func startApp() {
// 	var tasks []string
// 	var marks []bool

// 	for {
// 		input, _ := scanChoice(tasks, marks, true)

// 		switch input {
// 		case 1:
// 			tasks, marks = newTask(tasks, marks)
// 		case 2:
// 			tasks, marks = deleteTask(tasks, marks)
// 		case 3:
// 			tasks, marks = markTask(tasks, marks)
// 		case 4:
// 			return
// 		}
// 	}
// }

// func scanChoice(tasks []string, marks []bool, start bool) (int, error) {
// 	reader := bufio.NewReader(os.Stdin)

// 	utils.ClearScreen()
// 	printTasks(tasks, marks)

// 	if start {
// 		fmt.Print(menu)
// 		utils.PrintLine(tasks)
// 		fmt.Print("[1-4] ")
// 	} else {
// 		if tasks == nil {
// 			return -1, errors.New("tasks slice is nil")
// 		}
// 		if len(tasks) == 1 {
// 			fmt.Print("[1] ")
// 		} else {
// 			fmt.Printf("[1-%v] ", len(tasks))
// 		}
// 	}

// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		return -1, errors.New("input error")
// 	}

// 	input = strings.TrimSpace(input)

// 	if utils.IsDigitString(input) {
// 		input_num, err := strconv.Atoi(input)
// 		if err != nil {
// 			return -1, errors.New("input error")
// 		}
// 		if start {
// 			if input_num > 0 && input_num < 6 {
// 				return input_num, nil
// 			}
// 		} else {
// 			if input_num > 0 && input_num <= len(tasks) {
// 				return input_num, nil
// 			}
// 		}
// 	}
// 	return -1, errors.New("некорректный ввод")
// }

// func printTasks(tasks []string, marks []bool) {
// 	if len(tasks) == 0 {
// 		return
// 	}

// 	for i, task := range tasks {
// 		if marks[i] {
// 			fmt.Printf("%d: [x] %s\n", i+1, task)
// 		} else {
// 			fmt.Printf("%d: [ ] %s\n", i+1, task)
// 		}
// 	}
// 	utils.PrintLine(tasks)
// }

// func newTask(tasks []string, marks []bool) ([]string, []bool) {
// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		utils.ClearScreen()
// 		printTasks(tasks, marks)

// 		if len(tasks) == 0 {
// 			fmt.Println("Введите текст")
// 		}
// 		fmt.Print("[т] ")
// 		text, err := reader.ReadString('\n')
// 		if err != nil {
// 			fmt.Println("Ошибка ввода:", err)
// 			continue
// 		}

// 		text = strings.TrimSpace(text)

// 		tasks = append(tasks, text)
// 		marks = append(marks, false)

// 		break
// 	}

// 	return tasks, marks
// }

// func deleteTask(tasks []string, marks []bool) ([]string, []bool) {
// 	if len(tasks) == 0 {
// 		return tasks, marks
// 	}

// 	index, err := scanChoice(tasks, marks, false)

// 	if index > 0 && index <= len(tasks) && err == nil {
// 		tasks = slices.Delete(tasks, index-1, index)
// 		marks = slices.Delete(marks, index-1, index)
// 	}
// 	return tasks, marks
// }

// func markTask(tasks []string, marks []bool) ([]string, []bool) {
// 	if len(tasks) == 0 {
// 		return tasks, marks
// 	}

// 	index, err := scanChoice(tasks, marks, false)

// 	if index > 0 && index <= len(tasks) && err == nil {
// 		if marks[index-1] {
// 			marks[index-1] = false
// 		} else {
// 			marks[index-1] = true
// 		}
// 	}
// 	return tasks, marks
// }
