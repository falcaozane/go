package main

import "fmt"

func main() {

	fmt.Println("Loops in Go")

	fmt.Println("-------------")
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------------")
	// for-each loop
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	fmt.Println("-------------")

	// while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	fmt.Println("-------------")

	// do-while loop
	k := 0
	for {
		fmt.Println(k)
		k++
		if k == 5 {
			break
		}
	}
}