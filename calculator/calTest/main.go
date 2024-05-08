package main

import (
	"math"
)

func CalRes(n float64) bool {
	a := 4
	return n == math.Pi*(float64(a))

}
func main() {
	CalRes(12.56636)
}
