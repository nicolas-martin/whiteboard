package iparam

import "fmt"

type IDoSomething interface {
	DoSomething() error
}

func Hello(a IDoSomething) {
	if err := a.DoSomething(); err != nil {
		fmt.Println("error occurred")
		return
	}
	fmt.Println("no error occurred")
}
