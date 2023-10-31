package main

import (
	"fmt"
)

func main() {
	_ = t()
}

func t() (err error) {
	fmt.Printf("named: %p \r\n", err)

	b, err := e()

	fmt.Printf("returned: %p \r\n", err)
	fmt.Println(b)

	b2, err := e()
	fmt.Printf("returned2: %p \r\n", err)

	fmt.Println(b2)

	return

}

func e() (string, error) {

	return "", fmt.Errorf("TESTERROR")
}
