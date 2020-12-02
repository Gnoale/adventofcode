package main

import (
	"github.com/Gnoale/adventofcode/puzzlein"
)

func findProductNaive(s []int) int {
	for i, t := range s {
		for j := i; s[j] < (2020-t) && j < len(s)-1; j++ {
			u := s[j]
			for k := j; s[k] <= (2020-t-u) && k < len(s)-1; k++ {
				if (t + u + s[k]) == 2020 {
					return (t * u * s[k])
				}
			}
		}
	}
	return 0
}

func findProductRecurse(s []int) int {
	l := len(s)
	if l == 1 {
		return 0
	}
	L := s[:l/2]
	R := s[l/2:]
	if r := findProductNaive(L); r != 0 {
		return r
	}
	if r := findProductNaive(R); r != 0 {
		return r
	}
	findProductRecurse(L)
	findProductRecurse(R)
	return 0
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
