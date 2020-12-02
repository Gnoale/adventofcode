package main

import (
	"fmt"

	"github.com/Gnoale/adventofcode/puzzlein"
)

func findProductNaive(s []int) int {
	//fmt.Println(s)
	for i, t := range s {
		for j := i; s[j] < (2020-t) && j < len(s)-1; j++ {
			u := s[j]
			for k := j; s[k] <= (2020-t-u) && k < len(s)-1; k++ {
				if (t + u + s[k]) == 2020 {
					fmt.Println(t, u, s[k])
					return (t * u * s[k])
				}
			}
		}
	}
	return 0
}

func findProductRecurse(s []int) {
	l := len(s)
	if l == 1 {
		return
	}
	L := s[:l/2]
	R := s[l/2:]
	//fmt.Println(L)
	//fmt.Println(R)
	if r := findProductNaive(L); r != 0 {
		return
	}
	if r := findProductNaive(R); r != 0 {
		return
	}

	//fmt.Println(L, R)
	findProductRecurse(L)
	findProductRecurse(R)
}

func main() {

	codes, err := puzzlein.GetInt("./day1/input")
	if err != nil {
		panic(err)
	}

	//r := findProductNaive(codes)
	//fmt.Println(r)

	findProductRecurse(codes)

}
