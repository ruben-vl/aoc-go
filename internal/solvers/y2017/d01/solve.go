package d01

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/ruben-vl/aoc-go/internal/solvers"
)

func init() {
	solvers.Register("y2017d01", Solve)
}

func Solve(part int, r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	var lines []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("reading input: %w", err)
	}

	switch part {
	case 1:
		return fmt.Sprintf("part 1: %d lines", len(lines)), nil
	case 2:
		return fmt.Sprintf("part 2: reversed first line: %s", reverse(lines[0])), nil
	default:
		return "", fmt.Errorf("unknown part %d", part)
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
