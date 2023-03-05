package main

import "fmt"

func main() {
	// simple output
	fmt.Println("Hello World!")
	
	// vars & types
	var v1 int
	var v11 uint64
	var v2 int32
	var v22 uint32
	var v3 int64
	var v4 float32
	var v5 float64
	var v6 bool
	var v7 string
	fmt.Println(v1, v11, v2, v22, v3, v4, v5, v6, v7)

	x := 1
	x = 2
	fmt.Println(x)

	test := 2
	fmt.Println(test)

	if isEven2(1) {
		fmt.Println(`even`)
	} else {
		fmt.Println(`not even`)
	}
	
	// structs

	account := Account{}
	fmt.Printf("Id: %v\n", account.id)
	fmt.Printf("Balance: $%0.2f\n", account.balance)
}

func isEven(n int) (bool, error) {
	if n <= 0 {
		return false, fmt.Errorf("Err: n must me > 0")
	}
	return n % 2 == 0, nil
}

func isEven2(n int) (bool) {
	return n % 2 == 0
}

func addN(n int) func(x int) int {
	return func(x int) int {
		return x + n
	}
}

type Account struct {
	id int
	balance float32
}

func (a *Account) String() string {
	return fmt.Sprintf("Account [%d]: $%0.2f", a.id, a.balance)
}

func (a *Account) Deposit(amount float32) float32 {
	a.balance += amount
	return a.balance
}

func (a *Account) Withdraw(amount float32) float32 {
	a.balance -= amount
	return a.balance
}

func (a *Account) Balance() float32 {
	return a.balance
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}