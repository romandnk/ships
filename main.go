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

const (
	bread   = "bread"
	banana  = "bananas"
	clothes = "clothes"
)

var (
	tunnel      = make(chan Ship, 5)
	breadPier   = make(chan Ship)
	bananaPier  = make(chan Ship)
	clothesPier = make(chan Ship)
	done        = make(chan struct{})
)

func main() {
	wg := sync.WaitGroup{}
	types := [3]string{bread, banana, clothes}

	wg.Add(3)
	for _, shipType := range types {
		go func(shipType string) {
			defer wg.Done()
			createShips(shipType)
		}(shipType)
	}

	go func() {
		wg.Wait()
		close(tunnel)
	}()

	go pier(breadPier)
	go pier(bananaPier)
	go pier(clothesPier)

	for ship := range tunnel {
		fmt.Printf("Новый корабль \"%s\" с вместимостью \"%d\"\n", ship.shipType, ship.capacity)
		switch ship.shipType {
		case bread:
			breadPier <- ship
		case banana:
			bananaPier <- ship
		case clothes:
			clothesPier <- ship
		}
	}

}

func createShips(shipType string) {
	arrCap := [...]int{10}
	for i := 0; i < 2; i++ {
		idx := rand.Intn(1)
		newShip := Ship{
			shipType: shipType,
			capacity: arrCap[idx],
		}
		tunnel <- newShip
	}
}

func pier(pier chan Ship) {
	for ship := range pier {
		fmt.Printf("Началась разгрузка корабля \"%s\" с вместимостью \"%d\"\n", ship.shipType, ship.capacity)
		for ship.capacity > 0 {
			ship.capacity = ship.capacity - 10
			time.Sleep(time.Second)
			fmt.Printf("Осталось \"%d\" у корабля \"%s\"\n", ship.capacity, ship.shipType)
		}
	}
}
