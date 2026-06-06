package main

import (
	"fmt"
)

func main() {
	fmt.Println("---- Simple Hash Table Version ----")
	TestSimpleHashTable()
	fmt.Println("---- Separate Chaining Version ----")
	TestChaining()
}
