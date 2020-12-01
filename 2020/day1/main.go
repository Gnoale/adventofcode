package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func findProductNaive(s []int) int {
	for i, _ := range s {
		t := s[i]
		for j := i; s[j] < (2020 - t); j++ {
			u := s[j]
			for k := j; s[k] <= (2020 - t - u); k++ {
				if (t + u + s[k]) == 2020 {
					fmt.Println(t, u, s[k])
					return (t * u * s[k])
				}
			}
		}
	}
	return 0
}

func getInput(f string) []int {
	fd, _ := os.Open(f)
	s := bufio.NewScanner(fd)

	codes := make([]int, 0)
	//s.Split(bufio.ScanLines)
	for s.Scan() {
		c, _ := strconv.Atoi(s.Text())
		codes = append(codes, c)
	}
	sort.Ints(codes)
	return codes

}

func main() {

	codes := getInput("input")

	r := findProductNaive(codes)
	fmt.Println(r)

}
