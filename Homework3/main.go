package main

import (
	"fmt"
)

// 1- misol

func main() {

	var son int

	fmt.Print("Enter number: ")
	fmt.Scan(&son)

	tuborNotTub(son)

}

func tuborNotTub(son int) {

	count := 0

	for i := 1; i <= son; i++ {
		if son%i == 0 {
			count++
		}
	}
	if count == 2 {
		fmt.Println("Son toq")
	} else {
		fmt.Println("toq emas")
	}
}

/* 2- misol

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	reverseDigits(num)
}

func reverseDigits(num int) {
	reversed := 0
	for num != 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	fmt.Println("Reversed number:", reversed)
}
*/

/* 3- misol

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	sum := numOfDigits(n)

	fmt.Println("Sum of digits:", sum)
}

func numOfDigits(num int) int {
	sum := 0
	for num != 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
*/

/* 4- misol

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	numOfDigits(n)
}

func numOfDigits(num int) {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	for i := 1; i < num; i++ {
		fmt.Print(i, " + ")
	}
	fmt.Println(num, " = ", sum)
}
*/

/* 5- misol

func main() {
	var num1, num2 int

	fmt.Print("Enter A number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter B number: ")
	fmt.Scan(&num2)

	reverseDigits(num1, num2)
}

func reverseDigits(num1, num2 int) {
	if num1 < num2 {
		for i := num1; i <= num2; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
	for i := num1; i >= num2; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
*/

/* 6- misol

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	numOfDigits(n)
}

func numOfDigits(num int) {
	for i := num; i > 1; i-- {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Print()
}
*/

/* 7- misol

func main() {
	var num1, num2, num3 int

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	fmt.Print("Enter third number: ")
	fmt.Scan(&num3)

	reverseDigits(num1, num2, num3)
}

func reverseDigits(num1, num2, num3 int) {
	min := num1
	if num2 < min {
		min = num2
	}
	if num3 < min {
		min = num3
	}

	max := num1
	if num2 > max {
		max = num2
	}
	if num3 > max {
		max = num3
	}

	fmt.Println(max - min)
}
*/

/* 8- misol

func main() {
	var num int

	fmt.Print("Enter number: ")
	fmt.Scan(&num)

	fmt.Println(trueOrFalse(num))
}

func trueOrFalse(num int) bool {
	sum := 0
	for num != 0 {
		sum += num % 10
		num /= 10
	}

	if sum%2 == 0 {
		return true
	}
	return false
}
*/

/* 9- misol

func main() {

	var num1, num2 float32

	fmt.Print("Enter A number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter B number: ")
	fmt.Scan(&num2)

	reverseDigits(num1, num2)
}

func reverseDigits(num1, num2 float32) {
	fmt.Println((num1 / num2))
}
*/

/* 10- misol

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	sum := numOfDigits(n)

	fmt.Println("Sum of digits:", sum)
}

func numOfDigits(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}
*/

/* 13- misol

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	sum := numOfDigits(n)

	fmt.Println("Sum of digits:", sum)
}

func numOfDigits(num int) int {
	sum := 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	return sum
}
*/

/* 14- misol

func main() {
	var num1, num2, num3 int

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	fmt.Print("Enter third number: ")
	fmt.Scan(&num3)

	fmt.Println(reverseDigits(num1, num2, num3))
}

func reverseDigits(num1, num2, num3 int) int {
	if num1+num2 == num3 {
		return 1
	}
	return 0
}
*/

/* 15- misol

func main() {

	var haqiqiySon float64
	var butunSon int

	fmt.Print("Haqiqiy sonni kiriting: ")
	fmt.Scan(&haqiqiySon)

	fmt.Print("Butun sonni kiriting: ")
	fmt.Scan(&butunSon)

	natija := darajaTopish(haqiqiySon, butunSon)

	fmt.Printf("Javob: %f\n", natija)
}

func darajaTopish(haqiqiySon float64, butunSon int) float64 {
	daraja := math.Pow(haqiqiySon, float64(butunSon))
	return daraja
}
*/

/* 16- misol

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	numOfDigits(n)
}

func numOfDigits(num int) {
	for i := 0; i < num; i++ {
		if i%2 > 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
*/

/* 17- misol

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	numOfDigits(n)
}

func numOfDigits(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println(i, " * ", j, " = ", i*j)
		}
		fmt.Println("---------------------------------")
	}
}
*/
