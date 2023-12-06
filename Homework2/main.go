package main

import (
	"fmt"
)

func main() {

	/* 1- misol

	A := 5

	B := 8

	if A < B {
		c := A
		d := B

		fmt.Printf("A son = %d va B son = %d\n", d, c)
	}

	*/

	/* 2- misol

	var (
		A int
		B int
	)

	fmt.Printf("A sonni kiriting: ")
	fmt.Scan(&A)
	fmt.Printf("B sonni kiriting: ")
	fmt.Scan(&B)

	if A == B {
		fmt.Println("0")
	} else {
		C := A + B
		fmt.Printf("Javob: %d\n", C)
	}

	*/

	/* 3- misol

	var (
		A int
		B int
	)

	fmt.Printf("A sonni kiriting: ")
	fmt.Scan(&A)
	fmt.Printf("B sonni kiriting: ")
	fmt.Scan(&B)

	if A == B {
		fmt.Println("A son = 0 va B son = 0")
	} else {
		if A > B {
			fmt.Printf("Ikala son ham %d ga teng \n", A)
		}
		if A < B {
			fmt.Printf("Ikala son ham %d ga teng \n", B)
		}
	}
	*/

	/* 4- misol

	var (
		A int
		B int
		C int
	)

	fmt.Printf("A sonni kiriting: ")
	fmt.Scan(&A)
	fmt.Printf("B sonni kiriting: ")
	fmt.Scan(&B)
	fmt.Printf("C sonni kiriting: ")
	fmt.Scan(&C)

	min := A

	if B < min {
		min = B
	}

	if C < min {
		min = C
	}

	fmt.Printf("Eng kichik son: %d\n", min)

	*/

	/* 5- misol

	var (
		A int
		B int
		C int
	)

	fmt.Printf("A sonni kiriting: ")
	fmt.Scan(&A)
	fmt.Printf("B sonni kiriting: ")
	fmt.Scan(&B)
	fmt.Printf("C sonni kiriting: ")
	fmt.Scan(&C)

	urtacha := (A + B + C) / 3

	fmt.Printf("O'rtacha son: %d\n", urtacha)
	*/

	/* 6- misol

	var (
		A int
		B int
		C int
	)

	fmt.Printf("A sonni kiriting: ")
	fmt.Scan(&A)
	fmt.Printf("B sonni kiriting: ")
	fmt.Scan(&B)
	fmt.Printf("C sonni kiriting: ")
	fmt.Scan(&C)

	min := A

	if B < min {
		min = B
	}

	if C < min {
		min = C
	}

	max := A

	if B > max {
		max = B
	}

	if C > max {
		max = C
	}

	fmt.Printf("Eng kichik son: %d va eng katta son: %d\n", min, max)
	*/

	/* 7- misol

	var (
		A int
		B int
		C int
	)

	fmt.Printf("A sonni kiriting: ")
	fmt.Scan(&A)
	fmt.Printf("B sonni kiriting: ")
	fmt.Scan(&B)
	fmt.Printf("C sonni kiriting: ")
	fmt.Scan(&C)

	sum1 := A + B
	sum2 := B + C
	sum3 := A + C

	maxSum := sum1
	maxPair := 1

	if sum2 > maxSum {
		maxSum = sum2
		maxPair = 2
	}

	if sum3 > maxSum {
		maxSum = sum3
		maxPair = 3
	}

	switch maxPair {
	case 1:
		fmt.Printf("Eng katta yig'indi bo'lgan ikkita son: %d va %d\n", A, B)
	case 2:
		fmt.Printf("Eng katta yig'indi bo'lgan ikkita son: %d va %d\n", B, C)
	case 3:
		fmt.Printf("Eng katta yig'indi bo'lgan ikkita son: %d va %d\n", A, C)
	}
	*/

	/* 8- misol

	var (
		A int
		B int
		C int
	)

	fmt.Printf("A sonni kiriting: ")
	fmt.Scan(&A)
	fmt.Printf("B sonni kiriting: ")
	fmt.Scan(&B)
	fmt.Printf("C sonni kiriting: ")
	fmt.Scan(&C)

	if A < B && B < C {
		fmt.Printf("Sonlar o'sish tartibida bo'ldi, ularni ikkilantiramiz:\n")
		fmt.Printf("A soni: %d\n", A*2)
		fmt.Printf("B soni: %d\n", B*2)
		fmt.Printf("C soni: %d\n", C*2)
	} else {
		fmt.Println("-%d, -%d, -%d", A, B, C)
	}
	*/

	/* 9- misol

	var (
		A int
		B int
		C int
	)

	fmt.Printf("A sonni kiriting: ")
	fmt.Scan(&A)
	fmt.Printf("B sonni kiriting: ")
	fmt.Scan(&B)
	fmt.Printf("C sonni kiriting: ")
	fmt.Scan(&C)

	if A < B && B < C {
		fmt.Printf("Sonlar o'sish tartibida bo'ldi, ularni ikkilantiramiz:\n")
		fmt.Printf("A soni: %d\n", A*2)
		fmt.Printf("B soni: %d\n", B*2)
		fmt.Printf("C soni: %d\n", C*2)
	} else if A > B && B > C {
		fmt.Printf("Sonlar o'kamayish tartibida bo'ldi, ularni ikkilantiramiz:\n")
		fmt.Printf("A soni: %d\n", A*2)
		fmt.Printf("B soni: %d\n", B*2)
		fmt.Printf("C soni: %d\n", C*2)
	} else {
		fmt.Println("-%d, -%d, -%d", A, B, C)
	}
	*/

	/* 10- misol

	var (
		A int
		B int
		C int
	)

	fmt.Print("Birinchi sonni kiriting: ")
	fmt.Scan(&A)
	fmt.Print("Ikkinchi sonni kiriting: ")
	fmt.Scan(&B)
	fmt.Print("Uchinchi sonni kiriting: ")
	fmt.Scan(&C)

	if A == B {
		fmt.Printf("Uchinchi son: %d\n", C)
	} else if B == C {
		fmt.Printf("Birinchi son: %d\n", A)
	} else if A == C {
		fmt.Printf("Ikkinchi son: %d\n", B)
	} else {
		fmt.Println("Hech qanday ikkita son bir-biriga teng emas")
	}
	*/

	/* 11- misol

	var (
		A int
		B int
		C int
		D int
	)

	fmt.Print("Birinchi sonni kiriting: ")
	fmt.Scan(&A)
	fmt.Print("Ikkinchi sonni kiriting: ")
	fmt.Scan(&B)
	fmt.Print("Uchinchi sonni kiriting: ")
	fmt.Scan(&C)
	fmt.Print("To'rtinchi sonni kiriting: ")
	fmt.Scan(&D)

	if A == B && B == C {
		fmt.Printf("To'rtinchi son: %d\n", D)
	} else if B == C && C == D {
		fmt.Printf("Birinchi son: %d\n", A)
	} else if A == C && C == D {
		fmt.Printf("Ikkinchi son: %d\n", B)
	} else if A == B && B == D {
		fmt.Printf("Uchinchi son: %d\n", C)
	} else {
		fmt.Println("Hech qanday uchta son bir-biriga teng emas")
	}
	*/

	////////////////////////////////////////////////////////  for  \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\

	/* 1- misol

	var son int

	fmt.Print("Sonni kiriting: ")
	fmt.Scan(&son)
	if son%2 == 0 {
		for i := son; i >= 2; i -= 2 {
			fmt.Println(i)
		}
	} else {
		for i := son - 1; i >= 2; i -= 2 {
			fmt.Println(i)
		}
	}
	*/

	/* 2- misol

	var son int

	fmt.Print("Sonni kiriting: ")
	fmt.Scan(&son)
	if son%2 == 0 {
		for i := son - 1; i >= 1; i -= 2 {
			fmt.Println(i)
		}
	} else {
		for i := son; i >= 1; i -= 2 {
			fmt.Println(i)
		}
	}
	*/

	/* 3- misol

	for i := 10; i < 100; i++ {
		count := 0
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 0 {
			fmt.Println(i)
		}
	}
	*/

	/* 4- misol

	for i := 100; i < 1000; i++ {
		yuzlik := i / 100
		unlik := (i % 100) / 10
		birlik := i % 10
		if yuzlik == unlik || unlik == birlik || yuzlik == birlik {
			fmt.Println(i)
		}
	}
	*/

	/* 5- misol

	var (
		n       = 0
		tempNUm = 0
	)
	fmt.Print("Number: ")
	fmt.Scan(&n)

	tempNUm = n
	count := 0

	for tempNUm != 0 {
		tempNUm = tempNUm / 10
		count++
	}

	var temp = 1
	for i := 0; i < count; i++ {
		if i == 0 {
			for j := 1; j < count; j++ {
				temp *= 10
			}
		}
		fmt.Print(n/temp, " ")
		n %= temp
		temp /= 10
	}
	fmt.Println()
	*/

	/* 6- misol

	var n = 0
	fmt.Print("n = ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		for j := i; j < n-1; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			if k == 0 || k == i {
				fmt.Print("* ")
			} else if i == n-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	*/

	/* 7- misol

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == 0 || i == 4-1 || j == 0 || j == 4-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	*/

	////////////////////////////////////////////////////////  switch  \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\

	/* 1- misol

	var direction string
	var command int

	fmt.Print("Yo'nalishni tanlang (s/j/q/g): ")
	fmt.Scan(&direction)

	fmt.Print("Kamandani tanlang (0 - harakatni davom ettir, 1 - chapga buril, 2 - o'ngga buril): ")
	fmt.Scan(&command)

	// Berilgan kamandaga qarab robot holatini yangilash
	switch direction {
	case "s":
		// Shimolda yo'nalish
		switch command {
		case 0:
			fmt.Println("Robot harakatni davom ettiradi (shimolga)")
		case 1:
			fmt.Println("Robot chapga buriladi")
		case 2:
			fmt.Println("Robot o'ngga buriladi")
		}
	case "j":
		// Janubda yo'nalish
		switch command {
		case 0:
			fmt.Println("Robot harakatni davom ettiradi (janubga)")
		case 1:
			fmt.Println("Robot chapga buriladi")
		case 2:
			fmt.Println("Robot o'ngga buriladi")
		}
	case "q":
		// Sharqda yo'nalish
		switch command {
		case 0:
			fmt.Println("Robot harakatni davom ettiradi (sharqqa)")
		case 1:
			fmt.Println("Robot chapga buriladi")
		case 2:
			fmt.Println("Robot o'ngga buriladi")
		}
	case "g":
		// G'arbda yo'nalish
		switch command {
		case 0:
			fmt.Println("Robot harakatni davom ettiradi (g'arbga)")
		case 1:
			fmt.Println("Robot chapga buriladi")
		case 2:
			fmt.Println("Robot o'ngga buriladi")
		}
	default:
		fmt.Println("Noto'g'ri yo'nalish")
	}

	*/

	/* 2- misol

	var direction string
	var command1, command2 int

	fmt.Print("Yo'nalishni tanlang (s/j/q/g): ")
	fmt.Scan(&direction)

	fmt.Print("Birinchi kamandani tanlang (0 - o'ngga buril, 1 - chapga buril, 2 - burilish 180): ")
	fmt.Scan(&command1)

	fmt.Print("Ikkinchi kamandani tanlang (0 - o'ngga buril, 1 - chapga buril, 2 - burilish 180): ")
	fmt.Scan(&command2)

	switch direction {
	case "s":
		// Shimolda yo'nalish
		switch command1 {
		case 0:
			fmt.Println("Lakatr o'ngga buriladi (shimolga)")
		case 1:
			fmt.Println("Lakatr chapga buriladi (shimolga)")
		case 2:
			fmt.Println("Lakatr buriladi 180 gradus (shimolga)")
		}
		switch command2 {
		case 0:
			fmt.Println("Lakatr o'ngga buriladi (shimolga)")
		case 1:
			fmt.Println("Lakatr chapga buriladi (shimolga)")
		case 2:
			fmt.Println("Lakatr buriladi 180 gradus (shimolga)")
		}
	case "j":
		// Janubda yo'nalish
		switch command1 {
		case 0:
			fmt.Println("Lakatr o'ngga buriladi (janubga)")
		case 1:
			fmt.Println("Lakatr chapga buriladi (janubga)")
		case 2:
			fmt.Println("Lakatr buriladi 180 gradus (janubga)")
		}
		switch command2 {
		case 0:
			fmt.Println("Lakatr o'ngga buriladi (janubga)")
		case 1:
			fmt.Println("Lakatr chapga buriladi (janubga)")
		case 2:
			fmt.Println("Lakatr buriladi 180 gradus (janubga)")
		}
	case "q":
		// Sharqda yo'nalish
		switch command1 {
		case 0:
			fmt.Println("Lakatr o'ngga buriladi (sharqqa)")
		case 1:
			fmt.Println("Lakatr chapga buriladi (sharqqa)")
		case 2:
			fmt.Println("Lakatr buriladi 180 gradus (sharqqa)")
		}
		switch command2 {
		case 0:
			fmt.Println("Lakatr o'ngga buriladi (sharqqa)")
		case 1:
			fmt.Println("Lakatr chapga buriladi (sharqqa)")
		case 2:
			fmt.Println("Lakatr buriladi 180 gradus (sharqqa)")
		}
	case "g":
		// G'arbda yo'nalish
		switch command1 {
		case 0:
			fmt.Println("Lakatr o'ngga buriladi (g'arbga)")
		case 1:
			fmt.Println("Lakatr chapga buriladi (g'arbga)")
		case 2:
			fmt.Println("Lakatr buriladi 180 gradus (g'arbga)")
		}
		switch command2 {
		case 0:
			fmt.Println("Lakatr o'ngga buriladi (g'arbga)")
		case 1:
			fmt.Println("Lakatr chapga buriladi (g'arbga)")
		case 2:
			fmt.Println("Lakatr buriladi 180 gradus (g'arbga)")
		}
	default:
		fmt.Println("Noto'g'ri yo'nalish")
	}

	*/

}
