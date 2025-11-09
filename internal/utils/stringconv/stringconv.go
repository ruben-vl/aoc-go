package stringconv

import (
	"fmt"
	"strconv"
	"strings"
)

func DigitsFromString(s string) []int {
	out := make([]int, 0, len(s))
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			out = append(out, int(ch-'0'))
		}
	}
	return out
}

func NumbersFromString(s string) ([]int, error) {
	fields := strings.Fields(s)
	out := make([]int, 0, len(fields))
	for _, f := range fields {
		n, err := strconv.Atoi(f)
		if err != nil {
			return nil, fmt.Errorf("invalid number %q: %w", f, err)
		}
		out = append(out, n)
	}
	return out, nil
}

func NumbersFromStrings(ss []string) ([][]int, error) {
	out := make([][]int, 0)
	for _, s := range ss {
		nums, err := NumbersFromString(s)
		if err != nil {
			return nil, err
		}
		out = append(out, nums)
	}
	return out, nil
}

func WordsFromString(s string) []string {
	return strings.Fields(s)
}

func WordsFromStrings(ss []string) [][]string {
	ws := make([][]string, 0, len(ss))
	for _, s := range ss {
		ws = append(ws, strings.Fields(s))
	}
	return ws
}
