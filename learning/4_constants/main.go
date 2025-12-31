package main 

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)

	const (
		port = 8080
		host = "localhost"
	) 
	fmt.Printf("Listening to http://%s:%d", host, port)
}