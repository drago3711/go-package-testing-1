package notbar

import (
	"fmt"
	"github.com/drago3711/go-package-testing-2/pkg/bar"
)

func Print(){
	fmt.Println("not bar")
	bar.Print()
	bar.PrintFoo()
}
