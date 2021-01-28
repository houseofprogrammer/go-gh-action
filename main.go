package main

import (
	"fmt"
	"log"
	"time"

	"github.com/houseofprogrammer/go-gh-action/cmd"
)

// Version of package
var Version string

func sum(from, to int) int {
	res := 0
	for i := from; i <= to; i++ {
		res += i
	}

	return res
}

func main() {
	fmt.Println("version in main", Version)
	fmt.Println("version in root", cmd.GetVersion())

	var s1, s2 int

	go func() {
		s1 = sum(1, 100)
	}()

	s2 = sum(1, 10)

	time.Sleep(time.Second)

	log.Println(s1, s2)
}
