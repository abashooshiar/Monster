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
	fmt.Println("Du befindest dich in Raum : ", random)

	for p := range dimensionsArray {
		for i := 0; i < len(dimensionsArray); i++ {
			if random == p {
				fmt.Println("Angrenzende Räume sind : ", dimensionsArray[p]) //  :  [1,7,8]
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
	for h := range dimensional {
		for i := 0; i < len(dimensional); i++ {
			if random == h {
				//fmt.Println("Angrenzende Räume sind : ", dimensional[h]) //  :  [1,7,8]
				//fmt.Println("Du befindest dich in Raum:", j)             // j
				a := dimensional[h]
				t := FindIn(a[:], j)
				if t == 1 {
					fmt.Println("du darfst in Raum", j)
					return
				} else {
					fmt.Println("Du kannst von hier nicht in Raum ", j)
					return
				}

			}
		}
	}
}
func FindIn(a []int, j int) int {
	t := 0
	for _, content := range a {
		if content == j {
			t = 1
		}
	}
	return t
}
