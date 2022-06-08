package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == "ES" {
		return "Hola, " + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Manuel", ""))
}
