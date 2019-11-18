package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("This code has no any dependencies!")
	r := httprouter.New()
	r.RedirectTrailingSlash = true
}
