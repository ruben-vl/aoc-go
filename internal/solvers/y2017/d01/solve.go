package d01

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ruben-vl/aoc-go/internal/solvers"
)

func init() {
	solvers.Register("y2017d01", Solve)
}

func Solve(part int, r io.Reader) (string, error) {

	line, err := readSingleLine(r)
	if err != nil {
		return "", err
	}
	nums := digitsFromString(line)

	switch part {
	case 1:
		return fmt.Sprintf("Sum of consecutive equal values: %d", sumConsecutiveEqualValues(nums)), nil
	case 2:
		return fmt.Sprintf("Sum of opposite equal values: %d", sumOppositeEqualValues(nums)), nil
	default:
		return "", fmt.Errorf("unknown part %d", part)
	}
}

func readSingleLine(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", fmt.Errorf("reading input: %w", err)
		}
		return "", fmt.Errorf("no input line found")
	}
	return scanner.Text(), nil
}

func digitsFromString(s string) []int {
	out := make([]int, 0, len(s))
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			out = append(out, int(ch-'0'))
		}
	}
	return out
}

func sumConsecutiveEqualValues(nums []int) int {
	sum := 0
	for i, n := range nums {
		if n == nums[(i+1)%len(nums)] {
			sum += n
		}
	}
	return sum
}

func sumOppositeEqualValues(nums []int) int {
	sum := 0
	for i, n := range nums {
		if n == nums[(i+len(nums)/2)%len(nums)] {
			sum += n
		}
	}
	return sum
}
