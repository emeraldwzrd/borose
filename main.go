package main

import (
	"fmt"
	"math/rand"
)

type Caretaker struct {
	name           string
	hours          int64
	favoriteFlower string
}

type Randomizer interface {
	RandomizeSelect([]string) string
}
type GardenString string

func RandomizeSelect(stringStream []string) string {
	return stringStream[rand.Intn(len(stringStream))]
}

type Flower struct {
	name            string
	color           string
	magicalProperty string
}

func (f *Flower) String() string {
	return fmt.Sprintf("Flower: %v, Color: %v Magical Property: %v\n",
		f.name,
		f.color,
		f.magicalProperty)

}

type Garden struct {
	caretaker Caretaker
	flowers   []Flower
}

func (g *Garden) Flowers() []Flower {
	return g.flowers
}

func GrowGarden(caretaker Caretaker, flowers []Flower) *Garden {
	return &Garden{
		caretaker: caretaker,
		flowers:   flowers,
	}
}

func GrowCaretaker() Caretaker {
	return Caretaker{
		name:           "Sharrie",
		hours:          16,
		favoriteFlower: "Tulip",
	}
}

func GrowFlowersRandom(numFlowers int) []Flower {
	names := []string{"Blue Orchid", "Rose", "Tulip"}
	colors := []string{"green", "blue", "yellow", "scarlet"}
	magicalProperties := []string{"dexterity", "strength", "intellect"}
	var flowers []Flower
	for i := 0; i < numFlowers; i++ {
		flowers = append(flowers, Flower{
			name:            RandomizeSelect(names),
			color:           RandomizeSelect(colors),
			magicalProperty: RandomizeSelect(magicalProperties),
		})
	}
	return flowers
}

func main() {
	garden := initializeGarden(24)
	for _, flower := range garden.Flowers() {
		fmt.Println(flower)
	}
}
