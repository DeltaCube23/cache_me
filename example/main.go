package main

import (
	"fmt"

	"github.com/DeltaCube23/cache_me"
)

func main() {
	// use NewLIFO or newFIFO if you want to initialize those caches
	lru := cache_me.NewLRU(3)

	lru.Put("Adithya", "Rajesh")
	lru.Put("England", "London")
	fmt.Println(lru.GetLength())

	fmt.Println(lru.Get("Australia"))
	fmt.Println(lru.Get("Adithya"))
	lru.Put("Australia", "Sydney")

	lru.Put("Prime", "23")
	fmt.Println(lru.Get("Adithya"))
	fmt.Println(lru.Get("Australia"))

	lru.GetStats()
}
