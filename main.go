package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	random := randomInt(1, 20)
	fmt.Println("Dein aktuelles Zimmer ist : ", random)
	fmt.Print("Bitte geben Sie den Zielraum : ")
	var j int
	if _, err := fmt.Scan(&j); err != nil {
		log.Print("  Scan Destination node ", err)
		return
	}
	Switch(random, j)

	//fmt.Println(j)
}
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
func Switch(random, j int) {
	switch random {
	case 1:
		if j == 2 || j == 4 || j == 3 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 2:
		if j == 1 || j == 5 || j == 6 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 3:
		if j == 1 || j == 7 || j == 8 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 4:
		if j == 1 || j == 9 || j == 10 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 5:
		if j == 2 || j == 10 || j == 11 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 6:
		if j == 2 || j == 7 || j == 12 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 7:
		if j == 3 || j == 6 || j == 13 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 8:
		if j == 3 || j == 9 || j == 14 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 9:
		if j == 4 || j == 8 || j == 15 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 10:
		if j == 4 || j == 5 || j == 16 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 11:
		if j == 5 || j == 12 || j == 17 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 12:
		if j == 6 || j == 11 || j == 18 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 13:
		if j == 7 || j == 14 || j == 18 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 14:
		if j == 8 || j == 13 || j == 19 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 15:
		if j == 9 || j == 16 || j == 19 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 16:
		if j == 10 || j == 15 || j == 17 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 17:
		if j == 11 || j == 16 || j == 20 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 18:
		if j == 12 || j == 13 || j == 20 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 19:
		if j == 14 || j == 15 || j == 20 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	case 20:
		if j == 17 || j == 18 || j == 19 {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Raum "+s+" ist erlaubt.")
		} else {
			s := fmt.Sprintf("%d", j)
			io.WriteString(os.Stdout, "Sie dürfen nicht im Raum "+s+" rein.")
		}
		return
	}

}
