package main

import "fmt"

// START animals OMIT

type Cow struct{}

type Fox struct{}

func (f Fox) Talk() string {
	return "Ring-ding-ding-ding-dingeringeding"
}

// END animals OMIT

// START main OMIT

type Talker interface {
	Talk() string
}

func ask(n Talker) {
	fmt.Printf("What goes \"%s?\"\n", n.Talk())
}

func main() {
	ask(Cow{})
	ask(Fox{})
}

// END main OMIT
