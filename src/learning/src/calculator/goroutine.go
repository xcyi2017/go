package calculator

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// 测试并发程序

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		fmt.Printf("Error: %s is down!\n", api)
	}
	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

func Test() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}
	ch := make(chan string)
	for _, api := range apis {
		go checkAPI(api, ch)
	}
	for i := 0; i < len(apis)-2; i++ {
		fmt.Print(<-ch)
	}
	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func send(ch chan string, message string) {
	ch <- message
}

func Test_() {
	size := 2
	ch := make(chan string, size)
	send(ch, "one")
	send(ch, "two")
	send(ch, "three")
	send(ch, "four")
	fmt.Println("All data sent to the channel ...")

	for i := 0; i < size; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Done!")
}

// 并发实现求斐波那契
func Fib_(number float64, ch chan float64) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}
	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)
	ch <- x
}

func TestFib() {
	start := time.Now()
	ch := make(chan float64, 15)
	for i := 1; i < 15; i++ {
		go Fib_(float64(i), ch)
	}

	for i := 1; i < 15; i++ {
		fmt.Printf("Fib(%v): %v\n", i, <-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())

}
