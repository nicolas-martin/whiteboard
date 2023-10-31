package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

func main() {
	dst := map[string]int{
		"four":  4,
		"ten":   10,
		"three": 3,
		"two":   2,
	}

	src := map[string]int{
		"four":  44,
		"nine":  9,
		"one":   1,
		"three": 3,
		"two":   2,
	}

	src2 := map[string]int{
		"eleven": 11,
	}

	fmt.Printf("dst: %v \nsrc: %v\n", dst, src)

	maps.Copy(dst, src)
	maps.Copy(dst, src2)

	fmt.Printf("\ncopied variable (dst): %v\n", dst)
}
