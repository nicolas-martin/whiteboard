package main

import (
	"math/big"
)

func main() {
}

func createBig() {
	bi := big.NewInt(10)
	_ = bi
}

func createCustomBig() {
	cbi := fromInt(10)
	_ = cbi
}

type BigInt struct {
	big.Int
}

func fromInt(x int) *BigInt {
	bi := &BigInt{}
	bi.SetInt64(int64(x))
	return bi
}
