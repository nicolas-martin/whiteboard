package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(int64(1249123))
	a.Mul(a, big.NewInt(int64(1_000_000_000)))

	fmt.Println(a)

}
