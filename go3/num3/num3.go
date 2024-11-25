package main

import (
	"fmt"
	"sync"
	"time"
)

func makeSound(animal string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Printf("%s makes a sound!\n", animal)
}

func main() {
	var wg sync.WaitGroup
	animals := []string{"Dog", "Cat", "Cow", "Sheep", "Duck"}

	for _, animal := range animals {
		wg.Add(1)
		go makeSound(animal, &wg)
	}

	wg.Wait()
}
