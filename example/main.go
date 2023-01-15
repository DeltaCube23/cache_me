package main

import (
	"fmt"

	"github.com/DeltaCube23/cache_me"
)

func Trialrun(cm cache_me.Cache) {
	cm.Put("Adithya", "Rajesh")
	cm.Put("England", "London")
	fmt.Println(cm.GetLength())

	fmt.Println(cm.Get("Australia"))
	fmt.Println(cm.Get("Adithya"))
	cm.Put("Australia", "Sydney")

	cm.Put("Prime", "23")
	fmt.Println(cm.Get("Adithya"))
	fmt.Println(cm.Get("Australia"))

	cm.GetStats()
}

func main() {
	// dry runs
	cm1 := cache_me.NewLRU(3)
	fmt.Println("---LRU---")
	Trialrun(cm1)

	cm2 := cache_me.NewLIFO(3)
	fmt.Println("---LIFO---")
	Trialrun(cm2)

	cm3 := cache_me.NewFIFO(3)
	fmt.Println("---FIFO---")
	Trialrun(cm3)
}
