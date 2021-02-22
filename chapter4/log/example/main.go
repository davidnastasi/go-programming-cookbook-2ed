package main

import (
	"fmt"
	"github.com/davidnastasi/go-programing-cookbook-2ed/chapter4/log"
)

func main() {
	fmt.Println("basic logging and modification of logger:")
	log.Log()
	fmt.Println("logging 'handled' errors:")
	log.FinalDestination()
}

