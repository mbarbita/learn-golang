package main

import "fmt"
import "github.com/mbarbita/learn-golang/golang-book/chapter11/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
