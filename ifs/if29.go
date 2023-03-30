package ifs

import (
	"fmt"
	"log"
	"github.com/vladalexeev-git/homework_if_29.03/tools"
)


func PrintIntDesc() {
	fmt.Println(tools.Desc29)
	fmt.Print("Введите целое число: ")
	var i int
	_, err := fmt.Scan(&i)
	if err != nil {
		log.Println(err)
		return
	}

	if i > 0 {
		if tools.IsEven(i) {
			fmt.Println("Положительное, чётное число")
			return
		}
		fmt.Println("Положительное, нечётное число")
		return
	} else if i < 0 {
		if tools.IsEven(i) {
			fmt.Println("Отрицательное, чётное число")
			return
		}
		fmt.Println("Отрицательное, нечётное число")
		return
	} else {
		fmt.Println("Нулевое число")
	}
}

