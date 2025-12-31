package main

import "fmt"

func main() {
	
	// age := 10
	// if age >= 18 {
	// 	fmt.Println("You are an adult.")
	// } else if age >=12{
	// 	fmt.Println("You are a teenager.")
	// } else {
	// 	fmt.Println("You are a child.")
	// }

	var role = "user"
	var hasPermission = true

	if role == "admin" && hasPermission {
		fmt.Println("Access granted to admin panel.")
	} else if role == "user" && hasPermission {
		fmt.Println("Access granted to user dashboard.")
	} else {
		fmt.Println("Access denied.")
	}
}
