package main

import (
	"fmt"
)

// "math/rand"
// "time"

func main() {

	message := "Hello World"

	for i, v := range message {
		fmt.Println(i, v)
	}

}

// myMap := make(map[string]int)

// myMap["student"] = 5
// myMap["alice"] = 2

// v, s := myMap["alice"]
// // delete(myMap)
// fmt.Println(v, s)

// studentsScore := map[string]int{
// 	"alice": 90,
// 	"jame":  87,
// 	"jake":  80,
// }

// fmt.Println(studentsScore)

// maps.Equal()

// a := [5]int{1, 2, 3, 4, 5}

// s := a[1:4]

// fmt.Println(s, cap(s))

// s = append(s, 2, 4, 5)

// fmt.Println(s, cap(s))

// sliceCopy := make([]int, len(s))
// copy(sliceCopy, s)

// fmt.Println(s, sliceCopy)

// for i, v := range s {
// 	fmt.Println(i, v)
// }

// fmt.Println(slices.Equal(s, sliceCopy))

// originalArray := [3]int{1, 2, 3}

// var coppiedArray *[3]int

// coppiedArray = &originalArray

// coppiedArray[0] = 100

// fmt.Println(originalArray, *coppiedArray)

// var numbers [5]int

// numbers[len(numbers)-1] = 20
// fmt.Println(numbers)

// numbers[0] = 10
// fmt.Println(numbers)

// for index, value := range numbers {
// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
// }

// for i := 0; i < len(numbers); i++ {
// 	fmt.Println("Index:", i, ":", numbers[i])
// }

// fruits := [4]string{"apple", "banana", "chery", "orange"}
// fmt.Println("Fruits array:", fruits)

// originalArray := [5]int{1, 2, 3, 4, 5}

// copiedArray := originalArray

// copiedArray[0] = 10
// fmt.Println("Original array:", originalArray)
// fmt.Println("Copied array:", copiedArray)

// fruit := "apple"

// switch fruit {
// case "apple":
// 	fmt.Println("You selected an apple.")
// case "banana":
// 	fmt.Println("You selected a banana.")
// default:
// 	fmt.Println("Unknown fruit selected.")
// }

// day := "Monday"

// switch day {
// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
// 	fmt.Println("It's a weekday.")
// case "Saturday", "Sunday":
// 	fmt.Println("It's a weekend.")
// default:
// 	fmt.Println("Unknown day selected.")
// }

// number := 15

// switch {
// case number < 10:
// 	fmt.Println("The number is less than 10.")
// case number >= 10 && number < 20:
// 	fmt.Println("The number is between 10 and 20.")
// case number >= 20:
// 	fmt.Println("The number is 20 or greater.")
// default:
// 	fmt.Println("Unknown number range.")
// }

// num := 2

// switch {
// case num > 1:
// 	fmt.Println("The number is greater than 1.")
// 	fallthrough
// case num == 2:
// 	fmt.Println("The number is equal to 2.")
// case num < 1:
// 	fmt.Println("The number is less than 1.")
// default:
// 	fmt.Println("The number is unknown.")
// }

// score := 85

// if score >= 90 {
// 	fmt.Println("Grade: A")
// } else if score >= 80 {
// 	fmt.Println("Grade: B")
// } else if score >= 70 {
// 	fmt.Println("Grade: C")
// } else {
// 	fmt.Println("Grade: D")
// }

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
