package main

import (
	"fmt"
	"wincode/wincode"
)

func main() {
	fmt.Println("Starting main")
	r := wincode.Encode("ooooooaaaa?")
	fmt.Println(r)
}
