package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	s := ""

	bi := &BigInt{}
	if _, ok := bi.SetString(s, 10); !ok {
		log.Fatal("invalid integer string")
	}

	fmt.Println(bi)

}

type BigInt struct {
	big.Int
}
