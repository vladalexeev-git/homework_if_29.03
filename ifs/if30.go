package ifs

import (
	"fmt"

	"github.com/vladalexeev-git/homework_if_29.03/tools"
)

func PrintThreeDigNumberDesc() {
	fmt.Println(tools.Desc30)
	fmt.Print("Введите  целое число, лежащее в диапазоне 1–999: ")
	var n int
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Это число не подходит")
		return
	}
	l := len(fmt.Sprint(n))
	if tools.IsEven(n) {
		switch l{
		case 3:
			fmt.Println("четное трехзначное число")
			return
		case 2:
			fmt.Println("четное двузначное число")
			return
		case 1:
			fmt.Println("четное однозначное число")
			return
		}
		
	} else {
		switch l{
		case 3:
			fmt.Println("нечетное трехзначное число")
			return
		case 2:
			fmt.Println("нечетное двузначное число")
			return
		case 1:
			fmt.Println("нечетное однозначное число")
			return
		}
	}
	
}