package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	/* 2- misol

	var numbers [10]int
	var sum int

	for i := 0; i < 10; i++ {
		fmt.Printf("Sonni kiriting (%d/%d): ", i+1, 10)
		fmt.Scan(&numbers[i])
	}

	fmt.Println("Tub sonlar:")
	for _, num := range numbers {
		count := 0
		for j := 2; j <= num/2; j++ {
			if num%j == 0 {
				count++
			}
		}
		if count == 0 {
			fmt.Println(num)
			sum++
		}
	}
	fmt.Printf("Tub sonlar soni: %d\n", sum)
	*/

	/* 3- misol

	var n int

	fmt.Print("Massiv uzunligini kiriting: ")
	fmt.Scan(&n)

	var numbers []int
	for i := 0; i < n; i++ {
		var num int
		fmt.Printf("Massivning %d-elementini kiriting: ", i+1)
		fmt.Scan(&num)
		numbers = append(numbers, num)
	}

	sumOfValues := 0
	sumOfIndexes := 0
	for i, num := range numbers {
		sumOfValues += num
		sumOfIndexes += i
	}

	if sumOfValues > sumOfIndexes {
		fmt.Println("INDEX")
	} else if sumOfValues < sumOfIndexes {
		fmt.Println("VALUES")
	} else {
		fmt.Println("WOW")
	}
	*/

	/* 4- misol
	rand.Seed(time.Now().UnixNano())

	N := 10
	var series []int

	for i := 0; i < N; i++ {
		num := rand.Intn(39) - 12
		series = append(series, num)
	}

	fmt.Println("Asli: ", series)

	var modifiedSeries []int
	for _, num := range series {
		if num < 0 {
			modifiedSeries = append(modifiedSeries, 0)
		} else {
			modifiedSeries = append(modifiedSeries, 1)
		}
	}

	fmt.Println("Almashtirilgani:", modifiedSeries)

	*/

	/* 5- misol

	rand.Seed(time.Now().UnixNano())

	N := 10
	var series []int

	for i := 0; i < N; i++ {
		num := rand.Intn(22) + 14
		series = append(series, num)
	}

	fmt.Println("Asli:", series)

	fmt.Print("Almashtirilgani: ")
	for _, num := range series {
		if num%2 == 0 {
			fmt.Printf("+ ")
		} else {
			fmt.Printf("%d ", num)
		}
	}
	fmt.Println()
	*/

	/* 6- misol

	rand.Seed(time.Now().UnixNano())

	N := 10
	var series []int

	for i := 0; i < N; i++ {
		num := rand.Intn(22) + 14
		series = append(series, num)
	}

	fmt.Println("Asli:", series)

	fmt.Print("Almashtirilgani: ")
	for _, num := range series {
		if num%2 != 0 {
			fmt.Printf("_ ")
		} else {
			fmt.Printf("%d ", num)
		}
	}
	fmt.Println()
	*/

	/* 7- misol

	rand.Seed(time.Now().UnixNano())

	N := 5
	var series []int

	for i := 0; i < N; i++ {
		num := rand.Intn(31) - 5
		series = append(series, num)
	}

	fmt.Println("Asl seriya:", series)

	for i := 1; i < 5; i++ {
		fmt.Print(series[i], " ")
	}
	fmt.Println(series[0])
	*/

	/* 8- misol

	rand.Seed(time.Now().UnixNano())

	N := 5
	var series []int

	for i := 0; i < N; i++ {
		num := rand.Intn(31) - 5
		series = append(series, num)
	}

	fmt.Println("Asl seriya:", series)

	for i := 1; i < 5; i++ {
		fmt.Print(series[i], " ")
	}
	fmt.Println(series[0])
	*/

	rand.Seed(time.Now().UnixNano())

	N := 5
	var series []int

	for i := 0; i < N; i++ {
		num := rand.Intn(24) - 15
		series = append(series, num)
	}

	fmt.Println("Asl seriya:", series)

	fmt.Print(series[N-1], " ")
	for i := 0; i < 4; i++ {
		fmt.Print(series[i], " ")
	}
	fmt.Println()

}
