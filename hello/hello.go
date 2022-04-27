package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(s string, l string) string {
	if s == "" {
		return Hello("World", l)
	} else {
		return greetingsPrefix(l) + s
	}
}

func greetingsPrefix(l string) (prefix string) {
	switch l {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", "English"))
}
