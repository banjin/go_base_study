package structs

import "math"

type Rectangle struct {
	Width float64
	Height float64
}

//把类型的第一个字母作为接收者变量是 Go 语言的一个惯例。
func (r Rectangle) Area() float64{
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64{
	return c.Radius * c.Radius * math.Pi
}

type Triangle struct{
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {

	return t.Base * t.Height / 2
}

func Perimeter(a Rectangle) float64 {

	return (a.Width + a.Height) * 2
}

//声明接口，这样我们可以定义适合不同参数类型的函数（参数多态）
type Shape interface {
    Area() float64
}

// func Area(a Rectangle) float64{
// 	return a.Height * a.Width
// }


