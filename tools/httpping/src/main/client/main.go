package main

import (
	"os"

	"github.com/starter-go/starter"
	"github.com/xu-shi-fu/xwxy/tools/httpping"
)

func main() {

	a := os.Args
	m := httpping.ModuleForClient()
	i := starter.Init(a)

	i.MainModule(m)

	i.WithPanic(true).Run()
}
