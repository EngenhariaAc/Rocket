package main

import "fmt"

func main() {
    altitude := 0

    fmt.Println("🚀 Launch simulation started...")

    for i := 0; i < 5; i++ {
        altitude += 100
        fmt.Printf("Altitude: %d km\n", altitude)
    }

    fmt.Println("🌌 Rocket reached orbit!")
}
