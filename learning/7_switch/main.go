package main

import "fmt"
import "time"

func main() {
	
	i:=2

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Other number")
	}

	fmt.Println("-----")

	// switch with multiple cases
	day:=time.Now().Weekday()
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a workday")
	}

	fmt.Println("-----")

	// type switch
	whoAmI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("I'm an integer and my value is %d\n", v)
		case string:
			fmt.Printf("I'm a string and my value is %s\n", v)
		case bool:
			fmt.Printf("I'm a boolean and my value is %t\n", v)
		default:
			fmt.Printf("Unknown type %T\n", v)
		}
	}
	whoAmI(1)
	whoAmI("Hello")
	whoAmI(true)
}