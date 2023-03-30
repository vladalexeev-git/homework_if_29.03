package ifs

import (
	"fmt"
	"log"

	"github.com/vladalexeev-git/homework_if_29.03/tools"
)

// Условие:
// "Дан номер года (положительное целое число). Определить количество
// дней в этом году, учитывая, что обычный год насчитывает 365 дней, а
// високосный — 366 дней. Високосным считается год, делящийся на 4, за
// исключением тех годов, которые делятся на 100 и не делятся на 400
// (например, годы 300, 1300 и 1900 не являются високосными, а 1200 и 2000
// — являются)."

func HowManyDays() {
	var year int
	fmt.Println(tools.Desc28)
	fmt.Print("Введите номер года: ")
	_, err := fmt.Scan(&year)
	if err != nil {
		log.Println(err)
		return
	}
	if year <= 0 {
		fmt.Println("Это число не подходит")
		return
	}

	if isLeapYear(year) {
		fmt.Println("Этот год високосный, количество дней:",366)
		return
	} else {
		fmt.Println("Этот год невисокосный, количество дней:",365)
		return
	}

}

func isLeapYear(y int) bool {

	if rest := y % 400; rest == 0{
		return true
	}

	if rest := y % 100; rest == 0 {
		return false
	}

	if rest := y % 4; rest == 0 {
		return true
	}

	return false
}