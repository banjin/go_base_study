package pointers

import "testing"
import "fmt"


func TestWallet(t *testing.T){


	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin){
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	}

	assertError := func(t *testing.T, got error, want string) {
		if got == nil {
			t.Fatal("wanted an error but didnt get one")
		}
		if got.Error() != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T){
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10)) //当调用一个函数或方法时，参数会被复制。
		fmt.Println("address of balance in test is", &wallet.balance)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T){
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err,  "cannot withdraw, insufficient funds")
	})
}