package main

import "fmt"

func main() {
	s := `hey
%s
%s`

	r := fmt.Sprintf(s, "a", "b")

	fmt.Println(r)

}
