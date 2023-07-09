package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Ship struct {
	shipType string
	capacity int
}

var (
	tunnel      = make(chan Ship, 5)
	breadPier   = make(chan Ship)
	bananaPier  = make(chan Ship)
	clothesPier = make(chan Ship)
)

func main() {
	wg := sync.WaitGroup{}

	types := [3]string{"bread", "banana", "clothes"}

	wg.Add(3)
	for _, shipType := range types {
		go func(shipType string) {
			createShips(shipType)
		}(shipType)
	}

	go func() {
		wg.Wait()
		close(tunnel)
	}()

	go func() {
		for ship := range breadPier {
			fmt.Printf("Началась разгрузка корабля \"%s\" с вместимостью \"%d\"\n", ship.shipType, ship.capacity)
			for ship.capacity > 0 {
				ship.capacity = ship.capacity - 10
				time.Sleep(time.Second)
				fmt.Printf("Осталось \"%d\" у корабля \"%s\"\n", ship.capacity, ship.shipType)
			}
		}
	}()

	go func() {
		for ship := range bananaPier {
			fmt.Printf("Началась разгрузка корабля \"%s\" с вместимостью \"%d\"\n", ship.shipType, ship.capacity)
			for ship.capacity > 0 {
				ship.capacity = ship.capacity - 10
				time.Sleep(time.Second)
				fmt.Printf("Осталось \"%d\" у корабля \"%s\"\n", ship.capacity, ship.shipType)
			}
		}
	}()

	go func() {
		for ship := range clothesPier {
			fmt.Printf("Началась разгрузка корабля \"%s\" с вместимостью \"%d\"\n", ship.shipType, ship.capacity)
			for ship.capacity > 0 {
				ship.capacity = ship.capacity - 10
				time.Sleep(time.Second)
				fmt.Printf("Осталось \"%d\" у корабля \"%s\"\n", ship.capacity, ship.shipType)
			}
		}
	}()

	for i := range tunnel {
		fmt.Printf("Новый корабль \"%s\" с вместимостью \"%d\"\n", i.shipType, i.capacity)
		switch i.shipType {
		case "bread":
			breadPier <- i
		case "banana":
			bananaPier <- i
		case "clothes":
			clothesPier <- i
		}
	}

	close(breadPier)
	close(bananaPier)
	close(clothesPier)
}

func createShips(shipType string) {
	arrCap := [...]int{10, 20, 30}
	for i := 0; i < 2; i++ {
		idx := rand.Intn(3)
		newShip := Ship{
			shipType: shipType,
			capacity: arrCap[idx],
		}
		tunnel <- newShip
	}
}
