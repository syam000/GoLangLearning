package main

import (

	"fmt"
	"os"
	"io/ioutil"
)

func main() {

	fmt.Printf("***********Read Json File From Command Line***********\n")

	arg := os.Args[1]

	if arg != "" {

		b, err := ioutil.ReadFile(arg)
		if err != nil {
			panic(err)
			return
		}

		fmt.Print(string(b))
	}
}
