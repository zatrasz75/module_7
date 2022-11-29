// Главный пакет калькулятора
package main

import (
	"fmt"
	"m_07/pkg/calc"
)

func main() {
	var number1, number2 float64
	var operator string
	var err error

	//Считывание из консоли первого числа и проверка на ошибку считывания
	fmt.Println("Введите первое число:")
	_, err = fmt.Scanln(&number1)
	if err != nil {
		// обработка ошибок.
		fmt.Printf("ошибка при сканировании первого числа: %v", err)
	}

	//Считывание из консоли оператора и проверка на ошибку считывания
	fmt.Println("Введите оператор:")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		// обработка ошибок.
		fmt.Printf("ошибка при сканировании оператора: %v", err)
	}

	//Считывание из консоли второго числа и проверка на ошибку считывания
	fmt.Println("Введите второе число:")
	_, err = fmt.Scanln(&number2)
	if err != nil {
		// обработка ошибок.
		fmt.Printf("ошибка при сканировании второго числа: %v", err)
	}

	//Экземпляр структуры calculator из пакета calc
	// И вызов метода Calculate
	calculator := calc.NewCalculator()
	// Передача полученных значений в Экземпляр структуры из консоли
	result := calculator.Calculate(number1, number2, operator)

	//--------------
	fmt.Printf("Равно : %g", result)
}
