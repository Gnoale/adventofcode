package main

import (
	"fmt"

	"github.com/Gnoale/adventofcode/puzzlein"
)

var bagMap map[string][]map[string]int

func findAncestor(color string, ancestor map[string]int) {
	for parent, childs := range bagMap {
		for _, child := range childs {
			for k, v := range child {
				if k == color {
					ancestor[parent] = v
					findAncestor(parent, ancestor)
				}
			}
		}
	}
}

func findChildbag(color string, childbag map[string]int) {
	for _, child := range bagMap[color] {
		for c, v := range child {
			if c != "other bags." {
				childbag[c] += childbag[color] * v
				findChildbag(c, childbag)
			}
		}
	}
}

func countColors(color string, index, previous int, count *int) {

	for _, colormap := range bagMap[color] { // [ map1[c1:n1], map2[c2:n2]..]
		for next, n := range colormap { // c1,n1 c2,n2
			if next != "other bags." {
				fmt.Printf("color = %v, next = %v, n = %v, index = %v\n", color, next, n, index)
				val := index * n * previous
				*count += val
				fmt.Printf("count = %v, current val = %v\n\n", *count, val)
				countColors(next, n, previous, count)
			}
		}
	}
}

func main() {

	b, err := puzzlein.GetStr("./day7/inputest2")
	if err != nil {
		panic(err)
	}
	bagMap = puzzlein.BuildBagmap(b)

	// part1
	//	ancestors := make(map[string]int)
	//	findAncestor("shiny gold", ancestors)
	//	fmt.Println(len(ancestors))

	// part2

	root := "shiny gold"
	childbag := make(map[string]int)
	childbag[root] = 1
	findChildbag(root, childbag)

	total := 0
	for k, v := range childbag {
		if k != root {
			total += v
		}
	}
	fmt.Println(total)
	fmt.Printf("%v\n\n", childbag)
	//fmt.Println(visited)

	var count int
	countColors(root, 1, 1, &count)
	fmt.Println(count)

}
