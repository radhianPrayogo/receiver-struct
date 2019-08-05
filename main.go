package main

import (
	"fmt"
	"sort"
)

type TheNumbers struct {
	Numbers []float64
}

func (n TheNumbers) sortNumber() {
	sort.Float64s(n.Numbers)
	fmt.Println(n.Numbers)
}

func main() {
	var Number1 = TheNumbers{[]float64{4, 3, 5, 7, 2, 4, 1, 4}}
	Number1.sortNumber()
}
