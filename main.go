package main

import (
	"os"
	"fmt"
	"github.com/iower/goshapes"
)

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


	// arrays and slices

	xs := [4]int{1, 2, 3, 4}

	for _, value := range xs {
		fmt.Printf("- %d\n", value)
	}

	ys := []int{1, 2 , 3, 4, 5}

	for _, value := range ys {
		fmt.Printf("-- %d\n", value)
	}

	zs := []int{}
	zs = append(zs, 1)
	zs = append(zs, 2)
	zs = append(zs, 3)

	for _, value := range zs {
		fmt.Printf("--- %d\n", value)
	}

	fmt.Printf("\n")
	zs[1] = 22

	for _, value := range zs {
		fmt.Printf("--- %d\n", value)
	}

	fmt.Printf("\n %d \n", zs[1])


	// maps

	testmap := map[string]int{
		"key1": 11,
		"key2": 22,
	}

	fmt.Println("value1 =", testmap["key1"])

	for key, value := range testmap {
		fmt.Printf("%s: %d \n", key, value)
	}

	// looping

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("sum", sum)

	for i := 0; i < 100; i++ {
		fmt.Println(FizzBuzz(i))
	}


	// os

	user := os.Getenv("USER")
	Hello(user)


	// error handling

	path := "no/such/file"
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("No such file", path, err)
	} else {
		fs, err := f.Stat()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("File size", fs.Size())
		}
	}


	// files

	AppendFile("testfile", "testcontent")


	// packages

	c := goshapes.Circle{Radius: 5}
	fmt.Printf("Area of %s: %0.2f\n", c, c.Area())


	// multiple
	fmt.Println(multiple(2, 3))


	// arrays
	fmt.Println("Arrays")

	var arr [4]int
	for i, value := range arr {
		fmt.Println(i, value)
	}

	arr2 := [...]int{1, 2, 3, 4, 5}
	for i, value := range arr2 {
		fmt.Println(i, value)
	}

	// arrays have value semantics
	arr3 := arr2
	arr3[0] = -1
	arr3[4] = -5
	for i, value := range arr2 {
		fmt.Println(i, value)
	}
	for i, value := range arr3 {
		fmt.Println(i, value)
	}

	// slices
	sl1 := []int{1, 2, 3}
	sl2 := make([]int, 4)
	var sl3 [][]int32
	sl4 := []byte("slice")

	fmt.Println("Slice 1:")
	for i, value := range sl1 {
		fmt.Println(i, value)
	}
	fmt.Println("Slice 2:")
	for i, value := range sl2 {
		fmt.Println(i, value)
	}
	fmt.Println("Slice 3:")
	for i, value := range sl3 {
		fmt.Println(i, value)
	}
	fmt.Println("Slice 4:")
	for i, value := range sl4 {
		fmt.Println(i, value)
	}

	sl5 := []int{1, 2, 3}
	sl6 := sl5
	sl6[0] = 0
	fmt.Println("slices have reference semantics:", sl5[0] == sl6[0])

	sl5 = append(sl5, 4, 5, 6)
	fmt.Println("sl5*", sl5)

	sl5 = append(sl5, []int{7, 8, 9}...)
	fmt.Println("sl5**", sl5)

	sl5 = append(sl5, sl5...)
	fmt.Println("sl5***")


	// pointers
	p, q := learnMemory()
	fmt.Println("p", p, *p)
	fmt.Println("q", q, *q)


	// maps

	m := map[string]int{
		"three": 3,
		"four": 4,
	}
	m["one"] = 1
	fmt.Println("Map", m)


	// switching on the type

	var data interface{}
	data = 12
	switch c := data.(type) {
	case string:
		fmt.Println(c, "is a string")
	case int:
		fmt.Println(c, "is an int64")
	default:
		fmt.Println("unknown")
	}
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


// flow control

func Guess(n int) string {
	N := 42
	if n == N {
		return "You got it!"
	} else if n < N {
		return "Too low1"
	} else {
		return "Too high!"
	}
}

func FizzBuzz(n int) string {
	switch true {
	case n % 15 == 0:
		return "FizzBuzz"
	case n % 3 == 0:
		return "Fizz"
	case n % 5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", n)
	}
}


// deferring

func Goodbye(name string) {
	fmt.Printf("Goodbye, %s\n", name)
}

func Hello(name string) {
	defer Goodbye(name)
	fmt.Printf("Hello, %s\n", name)
}


// error handling

func AppendFile(fn, text string) error {
	f, err := os.OpenFile(fn, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Error opening file for writing: %w", err)
	}
	defer f.Close()

	if _, err := f.Write([]byte(text)); err != nil {
		return fmt.Errorf("Error writing text to file: %w", err)
	}

	return nil
}


/* multiline comment
and multiple return
*/

func multiple(x, y int) (sum, prod int) {
	return x + y, x * y
}


// pointers

func learnMemory() (p, q *int) {
	// p = new(int)
	s := make([]int, 20)
	s[3] = 7
	r := -2
	return &s[3], &r
}


// named return

func learnNamedReturn(x, y int) (z int) {
	z = x * y
	return
}
