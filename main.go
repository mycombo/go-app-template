package main

import (
	"fmt"
	"inter-knot-server/config"
)

func main() {
	config := config.GetConfig()

	fmt.Printf("Hello %v!\n", config.Name)
}
