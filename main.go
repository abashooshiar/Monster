package main

import (
	"fmt"
	"log"
	"math/rand"
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

	for p := range dimensionsArray {
		for i := 0; i < len(dimensionsArray); i++ {
			if random == p {
				fmt.Println("Angrenzende Räume sind : ", dimensionsArray[p])
				break
			}
		}
	}

	fmt.Print("Wohin möchtest Du gehen? : ")
	var j int
	if _, err := fmt.Scan(&j); err != nil {
		log.Print("  Scan Destination node ", err)
		return
	}

	Finder(random, j, dimensionsArray)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
func Finder(random, j int, dimensional [21][3]int) {
	for h, _ := range dimensional {
		if h == random {
			a := dimensional[h]
			sw := FindIn(a[:], j)
			if sw == true {
				random = j
				Angrenzend(random, dimensional)
				break
			} else {
				fmt.Println("Du kannst von hier nicht in Raum ", j)
				Angrenzend(random, dimensional)
				break
			}
		}
	}
}
func FindIn(a []int, j int) bool {
	sw := false
	for _, content := range a {
		if content == j {
			sw = true
		}
	}
	return sw
}
