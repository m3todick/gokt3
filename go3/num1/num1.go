package main

import (
	"fmt"
	"sync"
)

type Animal struct {
	Name              string
	HighSpeed         int
	Size              string
	ClimbTree         bool
	RecognizeDiseases bool
}

func animalInfo(animal Animal, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Animal: %s, HighSpeed: %d, Size: %s, ClimbTree: %t, RecognizeDiseases: %t\n",
		animal.Name, animal.HighSpeed, animal.Size, animal.ClimbTree, animal.RecognizeDiseases)
}

func main() {
	var wg sync.WaitGroup
	animals := []Animal{
		{"Cheetah", 120, "Medium", false, false},
		{"Elephant", 25, "Large", false, true},
		{"Monkey", 30, "Medium", true, false},
		{"Dog", 45, "Medium", false, true},
		{"Cat", 30, "Small", true, false},
	}

	for _, animal := range animals {
		wg.Add(1)
		go animalInfo(animal, &wg)
	}

	wg.Wait()
}
