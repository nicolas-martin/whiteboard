package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b, _ := json.Marshal(nil)

	fmt.Println(string(b))

}
