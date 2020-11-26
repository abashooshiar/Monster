package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	dimensionsArray := [21][3]int{{},
		{2, 4, 3}, {1, 5, 6}, {1, 7, 8}, {1, 9, 10}, {2, 10, 11},
		{2, 7, 12}, {3, 6, 13}, {3, 9, 14}, {4, 8, 15}, {4, 5, 16},
		{5, 12, 17}, {6, 11, 18}, {7, 14, 18}, {8, 13, 19}, {9, 16, 19},
		{10, 15, 17}, {11, 16, 20}, {12, 13, 20}, {14, 15, 20}, {17, 18, 19}}

	rand.Seed(time.Now().UnixNano())
	random := randomInt(1, 20)

	Angrenzend(random, dimensionsArray)

}
func Angrenzend(random int, dimensionsArray [21][3]int) {
	fmt.Println("Du befindest dich in Raum : ", random)
	fmt.Println("Angrenzende Räume sind : ", dimensionsArray[random])
	fmt.Print("Wohin möchtest Du gehen? : ")

	var WohinGehen int
	if _, err := fmt.Scan(&WohinGehen); err != nil {
		if !(WohinGehen >= '0' && WohinGehen <= '9') {
			fmt.Println("Bitte geben Sie eine Nummer ein")
			Angrenzend(random, dimensionsArray)
			return
		}
		fmt.Fprintln(os.Stderr, err)
	}

	Finder(random, WohinGehen, dimensionsArray)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
func Finder(random, WohinGehen int, dimensional [21][3]int) {

	IndexOfDimensional := dimensional[random]
	sw := FindIn(IndexOfDimensional[:], WohinGehen)
	if sw == true {
		random = WohinGehen
		Angrenzend(random, dimensional)
		return
	} else {
		fmt.Println("Du kannst von hier nicht in Raum ", WohinGehen)
		Angrenzend(random, dimensional)
		return
	}
}
func FindIn(DimensionalContent []int, WohinGehen int) bool {
	SwitchKey := false
	for _, content := range DimensionalContent {
		if content == WohinGehen {
			SwitchKey = true
		}
	}
	return SwitchKey
}
