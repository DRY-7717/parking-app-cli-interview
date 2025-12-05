package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ParkingLot struct {
	capacity int
	slots    []string 
}

func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		capacity: capacity,
		slots:    make([]string, capacity),
	}
}

func (p *ParkingLot) Park(carNumber string) {
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] == "" {
			p.slots[i] = carNumber
			fmt.Printf("Allocated slot number: %d\n", i+1)
			return
		}
	}
	fmt.Println("Sorry, parking lot is full")
}

func (p *ParkingLot) Leave(carNumber string, hours int) {
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] == carNumber {
			p.slots[i] = ""

		
			charge := 10
			if hours > 2 {
				charge += (hours - 2) * 10
			}

			fmt.Printf(
				"Registration number %s with Slot Number %d is free with Charge $%d\n",
				carNumber,
				i+1,
				charge,
			)
			return
		}
	}
	fmt.Printf("Registration number %s not found\n", carNumber)
}

func (p *ParkingLot) Status() {
	fmt.Println("Slot No. Registration No.")
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] != "" {
			fmt.Printf("%d %s\n", i+1, p.slots[i])
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lot *ParkingLot

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		switch parts[0] {
		case "create_parking_lot":
			capacity, _ := strconv.Atoi(parts[1])
			lot = NewParkingLot(capacity)

		case "park":
			lot.Park(parts[1])

		case "leave":
			hours, _ := strconv.Atoi(parts[2])
			lot.Leave(parts[1], hours)

		case "status":
			lot.Status()
		}
	}
}
