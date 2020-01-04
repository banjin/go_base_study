package pointers

import "fmt"


type Wallet struct {
	balance int
 }

func (w *Wallet) Deposit(a int){
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += a 
}

func (w *Wallet) Balance() int{
	return w.balance
}