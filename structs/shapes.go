package structs


type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}


func Perimeter(a Rectangle) float64 {

	return (a.Width + a.Height) * 2
}

func Area(a Rectangle) float64{
	return a.Height * a.Width
}


