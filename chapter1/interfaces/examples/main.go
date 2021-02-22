package main

import (
	"bytes"
	"fmt"

	"github.com/davidnastasi/go-programing-cookbook-2ed/chapter1/interfaces"
)

func main() {
	in := bytes.NewReader([]byte("examples"))
	out := &bytes.Buffer{}
	fmt.Print("stdout on Copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}
	fmt.Println("out bytes buffer =", out.String())
	fmt.Print("stdout on PipeExample = ")
	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}
}
