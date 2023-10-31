package switchlist

import "fmt"

func main() {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}
	v := "add_validator"

	switch v {
	case StringInArray(value, first):
		fmt.Println("matches first")
	case StringInArray(value, second):
		fmt.Println("matches second")
	}
}

func StringInArray(str string, arr []string) bool {
	if arr == nil {
		return false
	}

	for _, s := range arr {
		if str == s {
			return true
		}
	}
	return false
}
