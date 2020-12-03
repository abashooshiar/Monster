package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	dimensionsArray := [20][3]int{
		{2, 4, 3}, {1, 5, 6}, {1, 7, 8}, {1, 9, 10}, {2, 10, 11},
		{2, 7, 12}, {3, 6, 13}, {3, 9, 14}, {4, 8, 15}, {4, 5, 16},
		{5, 12, 17}, {6, 11, 18}, {7, 14, 18}, {8, 13, 19}, {9, 16, 19},
		{10, 15, 17}, {11, 16, 20}, {12, 13, 20}, {14, 15, 20}, {17, 18, 19}}

	rand.Seed(time.Now().UnixNano())
	zufall := randomint(1, 20)
	Zufallszahl := randomint(1, 20)
	slice := make([]int, 2)
	slice[0] = randomint(1, 20)
	slice[1] = randomint(1, 20)
	fmt.Println("Zufallszahl(", Zufallszahl, ")")

	heimstart := startraum(slice, zufall, dimensionsArray)
	if heimstart == false {
		angrenzend(zufall, dimensionsArray, Zufallszahl, slice)
	} else if heimstart == true {
		//fmt.Println("Ich glaube, es gibt beim Startraum() ein Problem!!")
		rand.Seed(time.Now().UnixNano())
		slice := make([]int, 2)
		slice[0] = randomint(1, 20)
		slice[1] = randomint(1, 20)
		startraum(slice, zufall, dimensionsArray)
		angrenzend(zufall, dimensionsArray, Zufallszahl, slice)
	}

}
func angrenzend(zufall int, dimensionsArray [20][3]int, Zufallszahl int, slice []int) {
	fmt.Println("Du befindest dich in Raum : ", zufall)
	fmt.Println("Angrenzende Räume sind : ", dimensionsArray[zufall-1])

	for _, inhalts := range dimensionsArray[zufall-1] {
		if inhalts == Zufallszahl {
			fmt.Println("Es stinkt bestialisch.")
		}
		if inhalts == slice[0] || inhalts == slice[1] {
			fmt.Println("Du spürst einen Luftzug.")
		}
	}

	fmt.Print("Wohin möchtest Du gehen? : ")
	var wohingehen int
	if _, err := fmt.Scan(&wohingehen); err != nil {
		fmt.Print("Bitte geben Sie eine Nummer__ ")
	}

	if wohingehen == Zufallszahl {
		fmt.Println("Du hast das Monster gefunden!")
		fmt.Println("drücken Sie 'j', um das Spiel zu beenden.Andernfalls geben Sie eine Zahl ein.")

		var exit int
		if _, err := fmt.Scan(&exit); err != nil {
			os.Exit(0)
		}
	}

	if wohingehen == slice[0] || wohingehen == slice[1] {
		fmt.Println("Du fällst in ein bodenloses Loch und stirbst.")
		os.Exit(0)

	}
	finder(zufall, wohingehen, dimensionsArray, Zufallszahl, slice)
}

func randomint(min, max int) int {
	return min + rand.Intn(max-min)
}
func finder(zufall, wohingehen int, dimensional [20][3]int, Zufallszahl int, slice []int) {
	indexOfDimensional := dimensional[zufall-1]

	sw, _ := findin(indexOfDimensional[:], wohingehen, Zufallszahl, indexOfDimensional[:])
	if sw == true {
		fmt.Println("----------------------------")
		fmt.Println("Zufallszahl(", Zufallszahl, ")")
		fmt.Println("Loch*(", slice[0], " und ", slice[1], ")")
		zufall = wohingehen
		angrenzend(zufall, dimensional, Zufallszahl, slice)
	} else {
		fmt.Println("Du kannst von hier nicht in Raum ", wohingehen)
		angrenzend(zufall, dimensional, Zufallszahl, slice)
	}
}

func findin(arrayinhalt []int, wohingehen int, Zufallszahl int, s []int) (bool, int) {
	switchKey := false
	for _, inhalt := range arrayinhalt {
		if inhalt == wohingehen {
			switchKey = true
			//return switchKey, Zufallszahl
		}
	}
	return switchKey, 0
}
func startraum(slice []int, zufall int, dimensionsArray [20][3]int) bool {

	fmt.Println("Loch(", slice[0], " und ", slice[1], ")")
	switchKey := false

	for _, inhalt := range dimensionsArray[zufall-1] {
		if inhalt == slice[0] || inhalt == slice[1] {
			rand.Seed(time.Now().UnixNano())
			slice := make([]int, 2)
			slice[0] = randomint(1, 20)
			slice[1] = randomint(1, 20)
			switchKey = true
			continue
		} else {
			break
		}
	}
	return switchKey
}
