package main

import (
	_ "embed"
	"fmt"
)

//go:embed main.go
var code string

func main() {
	fmt.Println(code)
}
