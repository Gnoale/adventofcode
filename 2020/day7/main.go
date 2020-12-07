package main

import (
	"fmt"

	"github.com/Gnoale/adventofcode/puzzlein"
)

var bagMap map[string][]string

// could lead somewhere
func findColorNum(node []string, colorMap map[string]int) {
	for _, color := range node {
		if color != "other bags." {
			colorMap[color] = 1
			findColorNum(bagMap[color], colorMap)
		}
	}
}

func main() {
	b, err := puzzlein.GetStr("./day7/input")
	if err != nil {
		panic(err)
	}

	bagMap = puzzlein.BuildSimpleBagmap(b)
	// (wrong) part1
	root := bagMap["shiny gold"]
	colorMap := make(map[string]int)
	findColorNum(root, colorMap)
	fmt.Println(len(colorMap))

	// shiny gold [map[plaid orange:1] map[striped lavender:2] map[pale brown:4] map[wavy blue:5]]
}
