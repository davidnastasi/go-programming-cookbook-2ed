package main

import (
	"fmt"
	"github.com/davidnastasi/go-programming-cookbook-2ed/chapter3/nulls"
)

func main() {
	if err := nulls.BaseEncoding(); err != nil {
		panic(err)
	}

	fmt.Println()

	if err := nulls.PointerEncoding(); err != nil {
		panic(err)
	}

	fmt.Println()

	if err := nulls.NullEncoding(); err != nil {
		panic(err)
	}

	fmt.Println()

}
