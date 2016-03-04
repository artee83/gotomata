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
	width int = 100

	// The heigh of the level map
	height int = 100
)

//
var level [width][height]int

var seed int64
var fillPercent = 50 //rand.Intn(100)

func main() {

	GenerateMap()

}

func GenerateMap() {
	RandomFillMap()

	GenerateImage()
}

func RandomFillMap() {
	seed = (int64)(time.Now().Nanosecond() % 100)
	rand.Seed(seed)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if isBoundary(x, y) {
				level[x][y] = 1
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

func GenerateImage() {
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if level[x][y] == 1 {
				m.Set(x, y, image.Black)
			}
		}
	}

	out, err := os.Create("output.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = png.Encode(out, m)
	if err != nil {
		fmt.Println(err)
	}
}
