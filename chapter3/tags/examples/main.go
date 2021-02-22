package main

import (
	"fmt"
	"github.com/davidnastasi/go-programing-cookbook-2ed/chapter3/tags"
)

func main() {
	if err := tags.EmptyStruct(); err != nil {
		panic(err)
	}

	fmt.Println()

	if err := tags.FullStruct(); err != nil {
		panic(err)
	}
}
