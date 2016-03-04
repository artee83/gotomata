package main

import (
	"fmt"
	"image"
	"image/png"
	"math/rand"
	"os"
	"time"
)

const (
	// The width of the level map
	width int = 20

	// The heigh of the level map
	height int = 20
)

//
var level [width][height]int

var seed int64
var fillPercent = 40 //rand.Intn(100)

func main() {
	GenerateMap()
}

func GenerateMap() {
	RandomFillMap()
	for i := 0; i < 5; i++ {
		SmoothMap()
		GenerateImage(i)
	}

}

func RandomFillMap() {
	seed = (int64)(time.Now().Nanosecond() % 100)
	rand.Seed(seed)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if isBoundary(x, y) {
				level[x][y] = 1
			} else {

				if rand.Intn(100) < fillPercent {
					level[x][y] = 1
				} else {
					level[x][y] = 0
				}
			}
		}
	}
}

func isBoundary(x, y int) bool {
	if x == 0 || x == width-1 || y == 0 || y == height-1 {
		return true
	}

	return false
}

func SmoothMap() {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			neighbours := neighbourCount(x, y)
			if neighbours > 4 {
				level[x][y] = 1
			} else if neighbours < 4 {
				level[x][y] = 0
			}
		}
	}
}

func neighbourCount(x, y int) int {
	var neighbours int = 0
	for nx := (x - 1); nx <= (x + 1); nx++ {
		for ny := (y - 1); ny <= (y + 1); ny++ {
			if isBoundary(nx, ny) {
				neighbours++
			} else {
				if nx != x || ny != y {
					neighbours += level[x][y]
				}
			}

		}
	}

	return neighbours
}

func GenerateImage(i int) {
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if level[x][y] == 1 {
				m.Set(x, y, image.Black)
			}
		}
	}

	fileName := fmt.Sprintf("output%d.png", i)
	out, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = png.Encode(out, m)
	if err != nil {
		fmt.Println(err)
	}
}
