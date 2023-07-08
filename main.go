package main

import (
	"clinic/initializers"
	"fmt"
)

func init() {
	initializers.SetupDB()
}

func main() {
	fmt.Println("Rede")
}
