package main

import (
	"fmt"

	"github.com/halnique/go-algorithms/pkg/sort"
)

func main() {
	s := []int{1, 3, 5, 2, 4}
	fmt.Println("Bubble Sort")
	fmt.Printf("in: %v\n", s)
	fmt.Printf("out: %v\n", sort.Bubble(s))
}
