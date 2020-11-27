package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	dimensionsArray := [20][3]int{
		{2, 4, 3}, {1, 5, 6}, {1, 7, 8}, {1, 9, 10}, {2, 10, 11},
		{2, 7, 12}, {3, 6, 13}, {3, 9, 14}, {4, 8, 15}, {4, 5, 16},
		{5, 12, 17}, {6, 11, 18}, {7, 14, 18}, {8, 13, 19}, {9, 16, 19},
		{10, 15, 17}, {11, 16, 20}, {12, 13, 20}, {14, 15, 20}, {17, 18, 19}}

	rand.Seed(time.Now().UnixNano())
	random := randomInt(1, 20)
	angrenzend(random, dimensionsArray)

}
func angrenzend(random int, dimensionsArray [20][3]int) {
	fmt.Println("Du befindest dich in Raum : ", random)
	fmt.Println("Angrenzende Räume sind : ", dimensionsArray[random-1])
	fmt.Print("Wohin möchtest Du gehen? : ")

	var wohingehen int
	if _, err := fmt.Scan(&wohingehen); err != nil {
		fmt.Print("Bitte geben Sie eine Nummer ein : ")
		fmt.Scan(&wohingehen)
	}
	finder(random, wohingehen, dimensionsArray)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
func finder(random, wohingehen int, dimensional [20][3]int) {

	indexOfDimensional := dimensional[random-1]
	sw := findIn(indexOfDimensional[:], wohingehen)
	if sw == true {
		random = wohingehen
		angrenzend(random, dimensional)
	} else {
		fmt.Println("Du kannst von hier nicht in Raum ", wohingehen)
		angrenzend(random, dimensional)
	}
}
func findIn(dimensionalContent []int, wohingehen int) bool {
	switchKey := false
	for _, content := range dimensionalContent {
		if content == wohingehen {
			switchKey = true
		}
	}
	return switchKey
}
