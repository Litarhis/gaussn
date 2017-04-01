package main

// imports
import (
	"fmt"
	"gaussn/pkg/pic"
	"io/ioutil"
	"math/rand"
	"time"
)

// end imports

func picture(dx, dy int) [][]uint8 {
	rand.Seed(time.Now().UTC().UnixNano())
	var s = make([][]uint8, dy)
	for i := range s {
		s[i] = make([]uint8, dx)
		for j := range s[i] {
			s[i][j] = uint8(rand.Intn(255))
		}
	}
	return s
}

func main() {
	err := ioutil.WriteFile("gaussianNoise.png", pic.Show(picture), 0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("Filter created --> gaussianNoise.png")
	fmt.Println("\nPress any key...")
	fmt.Scanln()
}
