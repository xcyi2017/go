package calculator

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

// 声明接口
type Shape interface {
	Perimeter() float64
	Area() float64
}

// 实现接口
type Square_I struct {
	Size float64
}

func (s *Square_I) Area() float64 {
	return s.Size * s.Size
}

func (s *Square_I) Perimeter() float64 {
	return s.Size * 4
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintInformation(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter: ", s.Perimeter())
	fmt.Println()
}

// 实现字符串接口

type Stringer interface {
	String() string
}

type Person struct {
	Name, Country string
}

func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}

func PrintPersonInformation() {
	rs := Person{Name: "John Doe", Country: "USA"}
	ab := Person{"Mark COllins", "United Kingdom"}
	fmt.Printf("%s\n%s\n", rs, ab)
}

// 扩展现有实现
func Extend() {
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	writer := CustomWriter{}
	io.Copy(&writer, resp.Body)
}

type CustomWriter struct{}

type GitHubResponse []struct {
	FullName string `json:"full_name"`
}

func (w *CustomWriter) Write(p []byte) (n int, err error) {
	var resp GitHubResponse
	json.Unmarshal(p, &resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), nil
}

// 编写自定义服务器API

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func TestUserHtpp() {
	db := database{"Go T-Shiret": 25, "Go Jacket": 55}
	log.Fatal((http.ListenAndServe("localhost:8000", db)))
}

// 创建用于管理在线商店的程序包

type Account struct {
	FirstName string
	LastName  string
}

func (a *Account) ChangeName(lastName, firstName string) {
	a.FirstName = firstName
	a.LastName = lastName
}

func (a *Account) String() string {
	return fmt.Sprintf("firstName;%s, lastName=%s", a.FirstName, a.LastName)
}

type Employee struct {
	Money float64
	Account
}

func (e *Employee) AddCredits(money float64) {
	e.Money += money
}

func (e *Employee) RemoveCredits(money float64) {
	e.Money -= money
}

func (e *Employee) CheckCredits() {
	fmt.Printf("Money: %v\n", e.Money)
}

func TestMyCredit() {
	a := Account{LastName: "cy", FirstName: "x"}
	e := Employee{1000., a}

	e.CheckCredits()
	fmt.Println(a)
	a.ChangeName("xc", "y")
	fmt.Println(a)

	e.AddCredits(100)
	e.CheckCredits()
	e.RemoveCredits(200)
	e.CheckCredits()

}
