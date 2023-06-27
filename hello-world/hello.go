package main

import "fmt"

const french = "French"
const portuguese = "Portuguese"

const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const portugueseHelloPrefix = "Olá, "

func Hello(name string, language string) string {
	if name == "" {
		name = "Pandora"
	}

	prefix:= englishHelloPrefix

	switch language {
	case french:
			prefix = frenchHelloPrefix
	case portuguese:
			prefix = portugueseHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}