package main

// imports
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/litarhis/gaussn/pic"
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
	var width, height int
	fmt.Print("Enter width of the image: ")
	fmt.Scanln(&width)

	fmt.Print("Enter height of the image: ")
	fmt.Scanln(&height)
	err := ioutil.WriteFile("gaussianNoise.png", pic.Show(picture, width, height), 0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("Filter created --> gaussianNoise.png")
	fmt.Println("\nPress any key...")
	fmt.Scanln()
}
