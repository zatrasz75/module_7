// Пакет состоит из базовых функций и констант
// для реализации калькулятора
package calc

import "log"

const (
	operatorAddition       = "+"
	operatorSubtraction    = "-"
	operatorMultiplication = "*"
	operatorDivision       = "/"
)

// Не экспортируемая структура calculator.
type calculator struct{}

// NewCalculator Экспортируемая функция-конструктор
// для создания экземпляра структуры calculator.
func NewCalculator() calculator {
	return calculator{}
}

// Calculate Экспортируемый публичный метод Calculate с Указателем на структуру calculator,
// Принимающий в качестве параметров 2 числа типа float64 и строчный оператор (+, -, *, /)
// И возвращающий значение типа float64.
// Метод Calculate должен в зависимости от переданного оператора вызывать один из приватных методов.
// (Если в качестве оператора передан +, то должен быть вызван приватный метод для сложения чисел).
func (c *calculator) Calculate(number1, number2 float64, operator string) float64 {
	switch operator {
	case operatorAddition:
		return c.add(number1, number2)
	case operatorSubtraction:
		return c.subtract(number1, number2)
	case operatorMultiplication:
		return c.multiply(number1, number2)
	case operatorDivision:
		return c.divide(number1, number2)
	default:
		log.Printf("несуществующий оператор: %s\n", operator)
		return 0
	}
}

// Приватный метод операции Сложения принимающий на вход
// 2 числа типа float64 и возвращать значение типа float64
func (c *calculator) add(number1, number2 float64) float64 {
	return number1 + number2
}

// Приватный метод операции вычитание принимающий на вход
// 2 числа типа float64 и возвращать значение типа float64
func (c *calculator) subtract(number1, number2 float64) float64 {
	return number1 - number2
}

// Приватный метод операции Умножение принимающий на вход
// 2 числа типа float64 и возвращать значение типа float64
func (c *calculator) multiply(number1, number2 float64) float64 {
	return number1 * number2
}

// Приватный метод операции Деление принимающий на вход
// 2 числа типа float64 и возвращать значение типа float64
// Так же осуществляет проверку Деления на 0.
func (c *calculator) divide(number1, number2 float64) float64 {
	if number2 == 0 {
		log.Println("ошибка, деление на 0")
		return 0
	}

	return number1 / number2
}
