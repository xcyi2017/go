package main

import (
	"github.com/myuser/calculator"
)

func main() {
	calculator.FizzBuzz(50)
	// calculator.GetPrimeRange(1, 20)
	// calculator.TestPanic()
	// fmt.Println(calculator.Fib(30))
	// fmt.Println(calculator.RomNumConv("MCMXCIX")) // 1999

	// calculator.TestLogPrint()

	// calculator.TestLogFatal()

	// calculator.TestLogPrefix()

	// calculator.TestLogTimeFieldFormat()

	// calculator.TestLogTimeFieldFormat_()

	// s := calculator.UpperString("Learning Go!")
	// fmt.Println(s)               // Learning Go!
	// fmt.Println(s.UpperString()) // LEARNING GO!

	// t := calculator.ColoredTriangle{calculator.Triangle{3}, "blue"}
	// fmt.Println("Size:", t.Size)
	// fmt.Println("Perimeter:", t.Perimeter())

	// t := calculator.ColoredTriangle{calculator.Triangle{3}, "blue"}
	// fmt.Println("Size: ", t.Size)                             // 3
	// fmt.Println("Perimeter (colored", t.Perimeter())          // 18 重载的方法
	// fmt.Println("Perimeter (normal)", t.Triangle.Perimeter()) // 9 继承的方法

	// var s calculator.Shape = &calculator.Square_I{3}
	// fmt.Printf("%T\n", s)
	// fmt.Println("Area: ", s.Area())
	// fmt.Println("Perimeter: ", s.Perimeter())

	// var s calculator.Shape = &calculator.Square_I{Size: 3}
	// calculator.PrintInformation(s)

	// c := &calculator.Circle{Radius: 6}
	// calculator.PrintInformation(c)

	// calculator.PrintPersonInformation()

	// calculator.Extend()

	// calculator.TestUserHtpp()

	// calculator.TestMyCredit()

	// calculator.Test()

	// calculator.Test_()
	calculator.TestFib()
}
