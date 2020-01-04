package pointers

import "fmt"

type Bitcoin int  //类型别名

type Wallet struct {
	balance Bitcoin
 }

func (w *Wallet) Deposit(a Bitcoin){
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += a 
}

func (w *Wallet) Balance() Bitcoin{
	return w.balance
}


type Stringer interface {
	String() string
}

func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC", b)
}