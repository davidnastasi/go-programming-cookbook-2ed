package main

import "github.com/davidnastasi/go-programming-cookbook-2ed/chapter1/bytestrings"

func main() {

	err := bytestrings.WorkWithBuffer()
	if err != nil {
		panic(err)
	}
	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}
