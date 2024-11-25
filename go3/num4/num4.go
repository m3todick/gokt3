package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gen2brain/beeep"
)

func notifyAnimal(animal string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	err := beeep.Notify("Animal Notification", fmt.Sprintf("%s has been processed!", animal), "")
	if err != nil {
		fmt.Println("Error sending notification:", err)
	}
}

func main() {
	var wg sync.WaitGroup
	animals := []string{"Lion", "Tiger", "Bear"}

	for _, animal := range animals {
		wg.Add(1)
		go notifyAnimal(animal, &wg)
	}

	wg.Wait()
}
