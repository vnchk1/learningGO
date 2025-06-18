package main

import (
	"errors"
	"fmt"
	"github.com/vnchk1/mymath"
	"os"
	"strings"
	"unicode"
)

// C:\Users\user\GolandProjects\awesomeProject
// for task 1.1
func temp(C float64) {
	var F float64

	F = C*1.8 + 32

	fmt.Printf("Output: %.2f C = %.2f F", C, F)
}

// for task 1.2
func fizzBuzz(fizzArr [100]int) {
	for _, value := range fizzArr {
		switch {
		case value%3 == 0 && value%5 == 0:
			fmt.Println("FizzBuzz")
		case value%3 == 0:
			fmt.Println("Fizz")
		case value%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(value)
		}
	}
}

// for task 1.3
func SumDigits(n int) int {
	if n < 0 {
		n = -n
	}
	s := 0
	for n != 0 {
		s = s + n%10
		n /= 10
	}
	return s
}

// for task 2.1
type currencyConverter struct {
	rate float64
}

func (c currencyConverter) convertToUSD(n float64) float64 {
	return n / c.rate
}

func (c currencyConverter) convertToRUB(n float64) float64 {
	return n * c.rate
}

// for task 2.2
func filterEven(numbers ...int) []int {
	newSlice := make([]int, len(numbers))
	k := 0
	for _, number := range numbers {
		if number%2 == 0 {
			newSlice[k] = number
			k += 1
		}
	}
	return newSlice
}

// for task 2.3
func (c *bankAccount) deposit(amount float64) {
	c.initialAmount += amount
}

func (c *bankAccount) withdraw(amount float64) error {
	if c.initialAmount-amount < 0 {
		return errors.New("Amount you want to withdraw is more than curren amount")
	} else {
		c.initialAmount -= amount
		return nil
	}
}

func (c bankAccount) balance() float64 {
	return c.initialAmount
}

type bankAccount struct {
	initialAmount float64
	//deposit(amount float64)
	//withdraw(amount float64) error
	//balance(float64)
}

// for task 3.1
func onceIsUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

func ValidatePassword(pass string) error {
	if len(pass) < 8 {
		return errors.New("Password must be at least 8 characters")
	}
	if !strings.ContainsAny(pass, "0123456789") {
		return errors.New("Password must contain at least one number")
	}
	if !onceIsUpper(pass) {
		return errors.New("Password must contain at least one uppercase character")
	}
	return nil
}

// for task 3.2
const maxFileSize = 1 << 20

func ReadFile(path string) (string, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", errors.New("File does not exist")
		}
		return "", errors.New("Error with opening file")
	}
	if fileInfo.Size() > maxFileSize {
		return "", errors.New("File too large")
	}
	content, err1 := os.ReadFile(path)
	if err1 != nil {
		return "", errors.New("Error with reading file")
	}
	return string(content), nil
}

// for task 4.1
type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * 3.14159265359
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {

	//task 1.1
	//var C float64
	//fmt.Scan(&C)
	//temp(C)

	//task 1.2
	//fizzArr := [100]int{}
	//for i := 0; i < 100; i++ {
	//	fizzArr[i] = i
	//}
	//fizzBuzz(fizzArr)

	// task 1.3
	//var n int
	//fmt.Scan(&n)
	//fmt.Printf("SumDigits(%v) // %v", n, SumDigits(n))

	//task 2.1
	//converter := currencyConverter{75.5}
	//fmt.Println(converter.convertToUSD(755))

	//task 2.2
	//variadicSlice := make([]int, 100)
	//for i := range variadicSlice {
	//	variadicSlice[i] = i * i
	//}
	//fmt.Println(filterEven(variadicSlice...))

	//task 2.3
	//var input23, AmountInput float64
	//bankAcc := bankAccount{initialAmount: 100}
	//for {
	//	fmt.Println("Menu: \n 1.Deposit \n 2.Withdraw \n 3.Balance \n 4.Exit")
	//	fmt.Scan(&input23)
	//	switch input23 {
	//	case 1:
	//		fmt.Println("Enter amount you'd like to deposit")
	//		fmt.Scan(&AmountInput)
	//		bankAcc.deposit(AmountInput)
	//		continue
	//	case 2:
	//		fmt.Println("Enter amount you'd like to withdraw")
	//		fmt.Scan(&AmountInput)
	//		err := bankAcc.withdraw(AmountInput)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		continue
	//	case 3:
	//		fmt.Printf("Balance %v\n", bankAcc.balance())
	//		continue
	//	case 4:
	//		return
	//	default:
	//		fmt.Println("Choose a number")
	//		continue
	//	}
	//}

	//task 3.1
	//err := ValidatePassword("ggааааааа")
	//fmt.Println(err)

	//task 3.2
	//path := "C:/Users/user/Desktop/testDocument.txt"
	//reader, err := ReadFile(path)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(reader)

	//task 4.1
	//circle := Circle{radius: 5}
	//rectangle := Rectangle{width: 10, height: 5}
	//fmt.Println(circle.Area())
	//fmt.Println(rectangle.Area())

	//task 4.2
	numbers := []float64{1, 2, 3, 4, 5}
	avg := mymath.Average(numbers)
	fmt.Printf("Average of numbers: %.2f", avg)
}
