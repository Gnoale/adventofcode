package puzzlein

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
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
	//sort.Ints(codes)
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

func GetStrEl(f string) ([]string, error) {
	fd, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	s := bufio.NewScanner(fd)

	codes := make([]string, 0)
	s.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
			// We have a full newline-terminated line.
			return i + 2, data[0:i], nil
		}
		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return len(data), data, nil
		}
		// Request more data.
		return 0, nil, nil
	})
	for s.Scan() {
		codes = append(codes, s.Text())
	}
	return codes, nil
}

// clean take a list of string and return a list of lists of strings
// splitted by \n and space
func Clean(p []string) [][]string {
	new := make([][]string, len(p))
	for i, q := range p {
		// f = [s1 s2, s3]
		f := strings.Split(q, "\n")
		for _, r := range f {
			// s = [s1, s2, s3]
			s := strings.Split(r, " ")
			for _, t := range s {
				if t != "" {
					new[i] = append(new[i], t)
				}
			}
		}
	}
	return new
}
