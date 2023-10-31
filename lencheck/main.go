package main

import "fmt"

func main() {
	aa := somestruct{}

	if len(aa.a.b) > 0 {
		fmt.Println("hey")
		return
	}

}

type somestruct struct {
	a *otherstruct
}

type otherstruct struct {
	b []string
}
