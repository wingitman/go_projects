package main

import (
	"fmt"
)

type syphers struct {
}

func main() {
	input := "Hello world!"
  en := encode(input)
  de := decode(en)
  fmt.Println("Input:",input)
  fmt.Println("Encode:",en)
  fmt.Println("Decode:",de)
}

func encode(str string) []byte {
  result := []byte(str)
  return result
}

func decode(arr []byte)string {
  result := string(arr)
  return result
}

