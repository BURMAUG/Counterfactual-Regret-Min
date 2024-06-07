package main

import (
	"cfr/cfr"
	"fmt"
)

func main() {
	p1 := cfr.NewPlay()
	fmt.Print(p1.Strategy())
}
