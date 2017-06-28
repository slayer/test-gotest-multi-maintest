package main

import (
	"fmt"

	"github.com/slayer/test-gotest-multi-maintest/pkg1"
	"github.com/slayer/test-gotest-multi-maintest/pkg2"
	"github.com/slayer/test-gotest-multi-maintest/pkg3"
)

func main() {
	fmt.Println("This is main()")
	fmt.Printf("pkg1 FooFunc(): %s\n", pkg1.FooFunc())
	fmt.Printf("pkg2 FooFunc(): %s\n", pkg2.FooFunc())
	fmt.Printf("pkg3 FooFunc(): %s\n", pkg3.FooFunc())
}
