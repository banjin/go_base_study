package iteration

import "fmt"
import "testing"

func TestIter(t *testing.T){
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	if repeated != expected{
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}



// 示例函数, 如果去掉// Output那行，则示例函数则不会执行

// === RUN   TestIter
// --- PASS: TestIter (0.00s)
// PASS
// ok      github.com/banjin/iteration     0.640s

// go test -v 
func ExampleRepeat(){
	expected := Repeat("a", 5)
	fmt.Println(expected)
	// Output: "aaaaa"
	
}


//基准测试  go test -bench=.
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 1)
    }
}