package calculator

import (
	"fmt"
	"math"
	"strconv"
)

// 编写 FizzBuzz 程序
// 首先，编写一个用于输出数字（1 到 100）的程序，其中有以下变化：

// 如果数字可被 3 整除，则输出 Fizz。
// 如果数字可被 5 整除，则输出 Buzz。
// 如果数字可同时被 3 和 5 整除，则输出 FizzBuzz。
// 如果前面的情况都不符合，则输出该数字。
func FizzBuzz(number int) string {
	if t1 := number % 3; t1 == 0 {
		fmt.Println("Fizz")
		return "Fizz"
	} else if t2 := number % 5; t2 == 0 {
		fmt.Println("Buzz")
		return "Buzz"
	} else if t1 == 0 && t2 == 0 {
		fmt.Println("FizzBuzz")
		return "FizzBuzz"
	} else {
		a := strconv.Itoa(number)
		fmt.Println(a)
		return a
	}
}

// 查找质数
// 编写一个程序来查找小于 20 的所有质数。 质数是大于 1 的任意数字，只能被它自己和 1 整除。 “整除”表示经过除法运算后没有余数。
func findPrime(number int) bool {

	for i := 2; i < int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	if number > 0 && number != 1 {
		return true
	} else {
		return false
	}
}

func GetPrimeRange(from, to int) {

	for i := from; i < to; i++ {
		if findPrime(i) {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Println()
}

// test panic
// 要求用户输入一个数字，如果该数字为负数，则进入紧急状态

func TestPanic() {
	val := 0

	for {
		fmt.Print("\nEnter number: ")
		fmt.Scanf("%d", &val)
		fmt.Println("You entered: ", val)
		if val < 0 {
			panic("You entered a negative")
		} else if val == 0 {
			fmt.Println("0 is neither negative nor positive")
		} else {
			fmt.Printf("You entered: %d\n", val)
		}
	}

}

// 编写一个程序来计算斐波纳契数列

func Fib(n int) []int {
	if n < 2 {
		return nil
	}
	i, a, b, res := 2, 1, 1, []int{1, 1}
	for i < n {
		a, b = b, a+b
		res = append(res, b)
		i++
	}
	return res
}

// 创建罗马数字转换器

func RomNumConv(str string) int {
	dict := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	arabicVals := make([]int, len(str)+1)
	for idx, w := range str {
		if val, ok := dict[w]; ok {
			arabicVals[idx] = val
		} else {
			fmt.Printf("Error: The roman numberal %s has a "+
				"bad digit: %c\n", str, w)
			return 0
		}
	}

	number := 0
	for idx := 0; idx < len(str); idx++ {
		if arabicVals[idx] < arabicVals[idx+1] {
			arabicVals[idx] = -arabicVals[idx]
		}
		number += arabicVals[idx]
	}
	return number
}
