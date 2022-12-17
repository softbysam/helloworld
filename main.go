package main

import (
	"fmt"

	"github.com/Samvel333/gox"
)

func helloworld() string {
	return "Hello World!!"
}

func main() {
	fmt.Println(helloworld())
	println(gox.SaySomething("Hayastan"))
}
