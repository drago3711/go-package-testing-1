package main

import (
	"fmt"
	"github.com/drago3711/go-package-testing-1/pkg/foo"
	"github.com/drago3711/go-package-testing-1/pkg/notbar"
)

func main() {
	fmt.Println("Starting 1")
	foo.Print()
	notbar.Print()
}
