package puzzlein

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func GetInt(f string) ([]int, error) {
	fd, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	s := bufio.NewScanner(fd)

	codes := make([]int, 0)
	//s.Split(bufio.ScanLines)
	for s.Scan() {
		c, _ := strconv.Atoi(s.Text())
		codes = append(codes, c)
	}
	sort.Ints(codes)
	return codes, nil
}

func GetStr(f string) ([]string, error) {
	fd, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	s := bufio.NewScanner(fd)

	codes := make([]string, 0)
	//s.Split(bufio.ScanLines)
	for s.Scan() {
		codes = append(codes, s.Text())
	}
	return codes, nil
}
