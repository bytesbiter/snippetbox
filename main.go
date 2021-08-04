package main

import (
	"fmt"
	"os"
)

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[1])

	for i := 0; i < 2; i++ {
		fmt.Println(os.Args[i])
	}

	log("test")
	fmt.Println(add(2, 3))

	val, bol := power("test")
	fmt.Println(val)
	fmt.Println(bol)

	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}

	fmt.Println(goku.Name)
	fmt.Println(goku.Power)
}

func log(message string) {
	fmt.Println(message)
}

func add(a int, b int) int {
	return a + b
}

func power(name string) (string, bool) {
	return name, true
}
