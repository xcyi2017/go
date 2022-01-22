package calculator

import (
	"strings"
)

// 声明方法
// 原因是，struct定义的属性是小写开头的，不是public的，这样是不能跨包调用的！

type Triangle struct {
	Size int
}

type Square struct {
	Size int
}

func (t Square) Perimeter() int {
	return t.Size * 4
}

func (t *Triangle) Perimeter() int {

	return t.Size * 3
}

// 方法中的指针
// 如果方法仅可访问接收方的信息，则不需要在接收方变量中使用指针。 但是，依据 Go 的约定，如果结构的任何方法具有指针接收方，则此结构的所有方法都必须具有指针接收方，即使某个方法不需要也是如此。

func (t *Triangle) DoubleSize() {
	t.Size *= 2
}

// 声明其他类型的方法

type UpperString string

func (s UpperString) UpperString() string {
	return strings.ToUpper(string(s))
}

// 嵌入方法
// 也就是说，可以重用来自一个结构的属性，以避免出现重复并保持代码库的一致性。 类似的观点也适用于方法。 即使接收方不同，也可以调用已嵌入结构的方法。

type ColoredTriangle struct {
	Triangle
	Color string
}

// 自动完成
// func (t *ColoredTriangle) Perimeter() int {
// 	return t.Triangle.Perimeter()
// }

// 重载方法

func (t *ColoredTriangle) Perimeter() int {
	return t.Size * 3 * 2
}

// 方法中的封装
// “封装”表示对象的发送方（客户端）无法访问某个方法。 通常，在其他编程语言中，你会将 private 或 public 关键字放在方法名称之前。 在 Go 中，只需使用大写标识符，即可公开方法，使用非大写的标识符将方法设为私有方法。
