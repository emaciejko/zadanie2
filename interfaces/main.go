package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sp)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//very custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	//very custom logic for generating an english greeting
	return "Hola!"
}
