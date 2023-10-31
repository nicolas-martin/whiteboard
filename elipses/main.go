package main

import "fmt"

type filter struct {
	name  string
	value string
}

func main() {
	s := sayHello([]filter{{name: "n1", value: "v1"}, {name: "n2", value: "v2"}})
	fmt.Println(s)
	fmt.Println("---")

}

func sayHello(filter []filter) string {
	s := ""
	for _, v := range filter {
		fmt.Printf("%s\nwhere %s=%s", s, v.name, v.value)
	}
	return s

}
