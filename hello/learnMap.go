package main

import "fmt"

type City struct {
	name   string
	street string
	city   string
}

func learningMap() {

	var cities = make(map[int]City)

	city1 := City{"Vani", "Home handler", "Random city"}
	city2 := City{"Whatever", "Random street", "Random city1"}
	city3 := City{"Testing", "What's this", "Another random"}

	cities[0] = city1
	cities[1] = city2
	cities[2] = city3

	fmt.Println(cities)

	fmt.Println("Whatever")

}
