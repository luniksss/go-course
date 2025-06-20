package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func printErrorMessage(err error) {
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
}

func readTwoInts(reader *bufio.Reader) (int, int, error) {
    input, err := reader.ReadString('\n')
    if err != nil {
        return 0, 0, err
    }
    input = strings.TrimSpace(input)
    var x, y int
    _, err = fmt.Sscanf(input, "%d %d", &x, &y)
    if err != nil {
        return 0, 0, err
    }
    return x, y, nil
}

func addition(x int, y int) int {
	return x + y
}

func multiply(x int, y int) int {
	return x * y
}

func subtraction(x int, y int) int {
	return x - y
}

func division(x int, y int) int {
	if y == 0 {
		return 0
	}
	return x / y
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите операцию(addition, multiply, subtraction, division) или exit для выхода: ")
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		if input == "exit" {
			fmt.Println("Вы вышлии из калькулятора")
			break
		}

		switch input {
		case "addition": 
			fmt.Print("Введите слагаемые через пробел: ")
			x, y, err := readTwoInts(reader)
			printErrorMessage(err)
			fmt.Println("Сумма: ", addition(x, y))
		case "multiply":
			fmt.Print("Введите множители через пробел: ")
			x, y, err := readTwoInts(reader)
			printErrorMessage(err)
			fmt.Println("Произведение: ", multiply(x, y))
		case "subtraction":
			fmt.Print("Введите уменьшаемое и вычитаемое через пробел: ")
			x, y, err := readTwoInts(reader)
			printErrorMessage(err)
			fmt.Println("Разность: ", subtraction(x, y))
		case "division":
			fmt.Print("Введите делимое и делитель через пробел: ")
			x, y, err := readTwoInts(reader)
			printErrorMessage(err)
			fmt.Println("Частное(деление нацело): ", division(x, y))
		default:
			fmt.Println("Неизвестная операция")
		}
	}
}