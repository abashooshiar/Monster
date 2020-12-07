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

	slice := make([]int, 4)
	raum, zufallszahl, slice := random(slice)
	startraum(slice, raum, zufallszahl, dimensionsArray)

}
func angrenzend(raum int, dimensionsArray [20][3]int, zufallszahl int, slice []int) {
	fmt.Println("Du befindest dich in Raum : ", raum)
	fmt.Println("Angrenzende Räume sind : ", dimensionsArray[raum-1])
	number := map[string]int{"loch1": slice[0], "loch2": slice[1], "fledermaus1": slice[2], "fledermaus2": slice[3]}
	for _, inhalts := range dimensionsArray[raum-1] {

		if inhalts == zufallszahl {
			fmt.Println("Es stinkt bestialisch.")
		}
		if inhalts == number["loch1"] || inhalts == number["loch2"] {
			fmt.Println("Du spürst einen Luftzug.")

		}
		if inhalts == number["fledermaus1"] || inhalts == number["fledermaus2"] {
			fmt.Println("Du hörst ein Flattern.")
		}
	}

	fmt.Print("Wohin möchtest Du gehen? : ")
	var wohingehen int
	if _, err := fmt.Scan(&wohingehen); err != nil {
		fmt.Print("Bitte geben Sie eine Nummer ")
	}

	if wohingehen == zufallszahl {
		fmt.Println("Du hast das Monster gefunden!")
		fmt.Println("drücken Sie 'j', um das Spiel zu beenden.Andernfalls geben Sie eine Zahl ein.")

		var exit int
		if _, err := fmt.Scan(&exit); err != nil {
			os.Exit(0)
		}
	}

	if wohingehen == number["loch1"] || wohingehen == number["loch2"] {
		swi := false
		for _, inhalt := range dimensionsArray[raum-1] {
			if inhalt != wohingehen {
				swi = true
			} else if inhalt == wohingehen {
				swi = false
				break
			}
		}
		if swi == true {
			fmt.Println("Du kannst von hier nicht in Raum ** ", wohingehen)
			angrenzend(raum, dimensionsArray, zufallszahl, slice)
		} else if swi == false {
			fmt.Println("Du fällst in ein bodenloses Loch und stirbst.")
			os.Exit(0)
		}
	}

	if wohingehen == number["fledermaus1"] || wohingehen == number["fledermaus2"] {
		swi := false
		for _, inhalt := range dimensionsArray[raum-1] {
			if inhalt != wohingehen {
				swi = true
			} else if inhalt == wohingehen {
				swi = false
				break
			}
		}
		if swi == true {
			fmt.Println("Du kannst von hier nicht in Raum ***", wohingehen)
			angrenzend(raum, dimensionsArray, zufallszahl, slice)

		} else if swi == false {
			fmt.Println("Eine riesige Fledermaus pack dich und trägt dich in einen anderen Raum.")
			raum, zufallszahl, slice = random(slice)
			startraum(slice, raum, zufallszahl, dimensionsArray)
		}
	}

	finder(raum, wohingehen, dimensionsArray, zufallszahl, slice)
}

func randomint(min, max int) int {
	return min + rand.Intn(max-min)
}
func finder(raum, wohingehen int, dimensional [20][3]int, zufallszahl int, slice []int) {
	indexOfDimensional := dimensional[raum-1]

	sw, _ := findin(indexOfDimensional[:], wohingehen)
	if sw == true {
		fmt.Println("----------------------------")
		fmt.Println("Zufallszahl*(", zufallszahl, ")", "Loch*(", slice[0], ",", slice[1], ")", "Fledermaus*(", slice[2], ",", slice[3], ")")

		raum = wohingehen
		angrenzend(raum, dimensional, zufallszahl, slice)
	} else {
		fmt.Println("Du kannst von hier nicht in Raum ", wohingehen)
		angrenzend(raum, dimensional, zufallszahl, slice)
	}
}

func findin(arrayinhalt []int, wohingehen int) (bool, int) {
	switchKey := false
	for _, inhalt := range arrayinhalt {
		if inhalt == wohingehen {
			switchKey = true
		}
	}
	return switchKey, 0
}

func startraum(slice []int, raum int, zufallszahl int, dimensionsArray [20][3]int) {

	n := map[string]int{"l1": slice[0], "l2": slice[1], "f1": slice[2], "f2": slice[3], "z": zufallszahl,
		"d0": dimensionsArray[raum-1][0], "d1": dimensionsArray[raum-1][1], "d2": dimensionsArray[raum-1][2]}
	switchKey := false

	if n["d0"] == n["l1"] || n["d0"] == n["l2"] || n["d0"] == n["f1"] || n["d0"] == n["f2"] ||
		n["d1"] == n["l1"] || n["d1"] == n["l2"] || n["d1"] == n["f1"] || n["d1"] == n["f2"] ||
		n["d2"] == n["l1"] || n["d2"] == n["l2"] || n["d2"] == n["f1"] || n["d2"] == n["f2"] ||
		n["l1"] == n["f1"] || n["l1"] == n["f2"] || n["l2"] == n["f1"] || n["l2"] == n["f2"] ||
		n["f1"] == n["f2"] || n["l1"] == n["l2"] ||
		n["z"] == n["l1"] || n["z"] == n["l2"] || n["z"] == n["f1"] || n["z"] == n["f2"] ||
		n["z"] == n["d0"] || n["z"] == n["d1"] || n["z"] == n["d2"] {

		raum, zufallszahl, slice = random(slice)
		switchKey = true

	}
	if switchKey == true {
		startraum(slice, raum, zufallszahl, dimensionsArray)

	} else {
		//switchKey = false
		fmt.Println("Zufallszahl(", zufallszahl, ")", "Loch(", slice[0], ",", slice[1], ")", "Fledermaus(", slice[2], ",", slice[3], ")")
		angrenzend(raum, dimensionsArray, zufallszahl, slice)
	}
}

func random(slice []int) (int, int, []int) {
	rand.Seed(time.Now().UnixNano())
	raum := randomint(1, 20)
	zufallszahl := randomint(1, 20)
	slice[0] = randomint(1, 20)
	slice[1] = randomint(1, 20)
	slice[2] = randomint(1, 20)
	slice[3] = randomint(1, 20)
	return raum, zufallszahl, slice
}
