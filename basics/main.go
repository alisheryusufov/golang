package main

import (
	"fmt"
)

func main() {

	score := 85

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D")
	}

}

// source := rand.NewSource(time.Now().UnixNano())
// random := rand.New(source)

// target := random.Intn(100) + 1

// var guess int
// for {
// 	fmt.Println("Guess a number between 1 and 100:")
// 	fmt.Scanln(&guess)

// 	if guess < target {
// 		fmt.Println("Too low!")
// 	} else if guess > target {
// 		fmt.Println("Too high!")
// 	} else {
// 		fmt.Println("Congratulations! You guessed the number.")
// 		break
// 	}
// }

// num := 1

// 	for num <= 10 {

// 		if num%2 == 0 {
// 			num++
// 			continue
// 		}

// 		fmt.Println("Odd number:", num)
// 		num++
// 	}

// rows := 5

// for i := 1; i <= rows; i++ {
// 	for j := 1; j <= rows-i; j++ {
// 		fmt.Print("'")
// 	}

// 	for k := 1; k <= 2*i-1; k++ {
// 		fmt.Print("*")
// 	}

// 	fmt.Println()
// }

// 	for i := 0; i < 5; i++ {
// 	fmt.Println("iteration number:", i)
//  }

//  numbers := []int{1, 2, 3, 4, 5}
//  for i, num := range numbers {
// 	fmt.Printf("Index: %d, Value: %d\n", i, num)
//  }

// }
// for i := 1; i <= 10; i++ {

// 	if i%2 == 0 {
// 		continue
// 	}
// 	fmt.Println("Odd number:", i)

// 	if i == 7 {
// 		break
// 	}

// 	var a, b int = 10, 3
// 	var result int

// 	var x, y float64 = 10.5, 3.2
// 	var floatResult float64

// 	result = a + b
// 	println("The sum is:", result)

// 	result = a - b
// 	println("The difference is:", result)

// 	result = a * b
// 	println("The product is:", result)

// 	floatResult = x / y
// 	println("The float division result is:", floatResult)

// 	result = a % b
// 	println("The remainder is:", result)
// }

// printName()

// func printName() {
// 	firstName := "John"
// 	fmt.Println(firstName)
