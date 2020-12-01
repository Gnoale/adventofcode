package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	f, _ := os.Open("input")
	s := bufio.NewScanner(f)

	codes := make([]int, 0)
	//s.Split(bufio.ScanLines)
	for s.Scan() {
		c, _ := strconv.Atoi(s.Text())
		codes = append(codes, c)
	}

	sort.Ints(codes)

	for i, _ := range codes {
		t := codes[i]
		for j := i; codes[j] < (2020 - t); j++ {
			u := codes[j]
			for k := j; codes[k] <= (2020 - t - u); k++ {
				if (t + u + codes[k]) == 2020 {
					fmt.Println(t, u, codes[k])
					fmt.Println(t * u * codes[k])
				}
			}
		}
	}

}
