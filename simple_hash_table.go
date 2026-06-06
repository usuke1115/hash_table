package main

import (
	"fmt"
)

type SimpleHashTable struct {
	Table [100]int
}

func (hashTable SimpleHashTable) s_hash(key string) int {
	var sum int = 0
	for _, c := range key {
		sum += int(c)
	}

	return sum % 100
}

func (hashTable *SimpleHashTable) s_put(key string, value int) {
	index := hashTable.s_hash(key)
	hashTable.Table[index] = value
}

func (hashTable *SimpleHashTable) s_get(key string) int {
	index := hashTable.s_hash(key)
	return hashTable.Table[index]
}

func TestSimpleHashTable() {
	table := SimpleHashTable{}
	table.s_put("hoge", 1)
	table.s_put("egoh", 2)
	fmt.Printf("Value of hoge is %d\n", table.s_get("hoge"))
	fmt.Printf("Value of egoh is %d\n", table.s_get("egoh"))
}
