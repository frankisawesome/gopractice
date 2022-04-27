package main

import "fmt"

func Hello(s string) string {
	return "Hello, " + s
}

func main() {
	fmt.Println(Hello("World"))
}
