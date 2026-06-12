package main

import "fmt"

type Entry struct {
	Key   string
	Value int
}

type HashTable struct {
	Table [100][]Entry
}

func (hashTable HashTable) hash(key string) int {
	var sum int = 0
	for _, c := range key {
		sum += int(c)
	}

	return sum % 100
}

func (hashTable *HashTable) put(entry Entry) {
	index := hashTable.hash(entry.Key)
	hashTable.Table[index] = append(hashTable.Table[index], entry)
}

func (hashTable *HashTable) get(key string) (int, bool) {
	index := hashTable.hash(key)
	for _, e := range hashTable.Table[index] {
		if e.Key == key {
			return e.Value, true
		}
	}
	return -1, false
}

func (hashTable *HashTable) delete(key string) {
	index := hashTable.hash(key)
	bucket := hashTable.Table[index]

	for i, e := range bucket {
		if e.Key == key {
			hashTable.Table[index] = append(bucket[:i], bucket[i+1:]...)
			return
		}
	}
}

func TestChaining() {
	table := HashTable{}
	table.put(Entry{Key: "hoge", Value: 10})
	table.put(Entry{Key: "egoh", Value: 1})
	hoge, _ := table.get("hoge")
	egoh, _ := table.get("egoh")
	fmt.Printf("Value of hoge is %d\n", hoge)
	fmt.Printf("Value of egoh is %d\n", egoh)
	table.delete("hoge")
	fmt.Println(table.get("hoge"))
}
