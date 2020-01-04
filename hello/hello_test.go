package main


import "testing"

func TestHello(t *testing.T) {
    assertCorrectMessage := func(t *testing.T, got, want string){
        t.Helper() //告诉测试套件这个方法是辅助函数，当测试失败时候所报告的行号将在函数调用中而不是辅助函数内部
        if got != want {
            t.Errorf("got '%q' want '%q'", got, want)
        }
    }
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("shm")
        want := "hello shm"
        assertCorrectMessage(t, got, want)
    })
    t.Run("empty string defaults to 'world'", func(t *testing.T) {
        got := Hello("")
       
        want := "hello world"
        assertCorrectMessage(t, got ,want)
    	})	
}

