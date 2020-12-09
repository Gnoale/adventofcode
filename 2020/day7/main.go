package main

import (
	"fmt"

	"github.com/Gnoale/adventofcode/graph"
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

func findChildbag(color string, childbag map[string]int, graph map[string]int) {
	for _, child := range bagMap[color] {
		for c, v := range child {
			if c != "other bags." {
				if graph[color] < graph[c] {
					childbag[c] += childbag[color] * v
				}
				findChildbag(c, childbag, graph)
			}
		}
	}
}

func main() {

	b, err := puzzlein.GetStr("./day7/inputest")
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
	visited := graph.Bfs(root, bagMap)
	findChildbag(root, childbag, visited)

	total := 0
	for k, v := range childbag {
		if k != root {
			total += v
		}
	}
	fmt.Println(total)
	fmt.Printf("%v\n\n", childbag)

	fmt.Println(visited)

}
