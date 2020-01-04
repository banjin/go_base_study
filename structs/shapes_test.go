package structs


import "testing"

func TestPerimeter(t *testing.T){
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0 
	if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}

func TestArea(t *testing.T){

	t.Run("Rectangle", func(t *testing.T){
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0
		if want != got {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("Cycle", func(t *testing.T){

		circle := Circle{10}
        got := circle.Area()
        want := 314.16

        if got != want {
            t.Errorf("got %.2f want %.2f", got, want)
        }
	})
	
}