package iteration

import "testing"

func TestIter(t *testing.T){
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	if repeated != expected{
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}