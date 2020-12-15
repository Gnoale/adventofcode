package main

import (
	"fmt"

	"github.com/Gnoale/adventofcode/puzzlein"
)

type vertex struct {
	color string
	count int
	previous *node
	right *node
	left *node
}

type graph struct {
	[]*node
}


func (g *graph)populate (bags []string) {
	for _, line := range bags {
		sl := strings.Split(line, ",")  // sl = ["muted white bags contain 4 dark orange bags" "3 bright white bags."]
		s := strings.Split(sl[0], " ")  // s = ["muted" "white" "bags" "contain" ...]
		key := strings.Join(s[:2], " ") // key = "muted white"
		num, _ := strconv.Atoi(s[4])
		n := &node{key, }


		v := map[string]int{strings.Join(s[5:7], " "): num} // v = map["dark orange":4]
		m[key] = append(m[key], v)                          // m["mute dwhite"] = [map["dark orange"]:4]

		for i := 1; i < len(sl); i++ {
			s = strings.Split(sl[i], " ") // s = ["3", "bright", "white"...]
			num, _ := strconv.Atoi(s[1])
			v = map[string]int{strings.Join(s[2:4], " "): num}
			m[key] = append(m[key], v) // m["muted white"] = [map["dark orange"]:4 map["bright white"]:3]
		}
	}
	return m
}
func (g *graph)traverse (start string) {
	


}




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

func main() {

	b, err := puzzlein.GetStr("./day7/input")
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

}
