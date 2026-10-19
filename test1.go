package main

import "fmt"

func main() {
	fmt.Print("Hello World")
	fmt.Println("hello world")

	fmt.Print("a", "b", "c")
	fmt.Println("a", "b", "c")

	var str = "Hello World"
	fmt.Printf("%s %T", str, str)
	var a = 10
	fmt.Printf("a=%d %T", a, a)

}
