package pointers

import "testing"
import "fmt"

func TestWallet(t *testing.T){
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10)) //当调用一个函数或方法时，参数会被复制。
	got := wallet.Balance()
	fmt.Println("address of balance in test is", &wallet.balance)
	want := Bitcoin(10)
	if got != want {
        t.Errorf("got %s want %s", got, want)
    }
}