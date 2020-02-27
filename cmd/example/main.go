package main

import (
	"fmt"
	"os"
	"github.com/lesia-s/architecture-lab1"
)

func main() {
	var expression string
	for _, element := range os.Args[1:] {
		expression += element + " "
	}

	res, err := lab1.CalculatePostfix(expression)
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println("Error!", err)
	}

	fmt.Println(buildVersion)
}
