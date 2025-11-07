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
