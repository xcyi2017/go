package calculator

var logMessage = "[LOG]"

var Version = "1.0"

func interlSum(number int) int {
	return number - 1
}

func Sum(number1, number2 int) int {

	return number1 + number2 + interlSum(number1)
}
