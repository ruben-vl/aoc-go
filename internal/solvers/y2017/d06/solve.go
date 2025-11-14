package d06

import (
	"fmt"
	"io"

	"github.com/ruben-vl/aoc-go/internal/solvers"
	"github.com/ruben-vl/aoc-go/internal/utils/input"
	"github.com/ruben-vl/aoc-go/internal/utils/set"
	"github.com/ruben-vl/aoc-go/internal/utils/stringconv"
)

func init() {
	solvers.Register("y2017d06", Solve)
}

func Solve(part int, r io.Reader) (string, error) {

	lines, err := input.ReadLines(r)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	nums, err := stringconv.NumbersFromString(lines[0])
	if err != nil {
		return "", fmt.Errorf("failed to get numbers from input: %w", err)
	}

	switch part {
	case 1:
		return fmt.Sprintf("Number of redistributions before repetition: %d", numRedistributions(nums, set.NewSet(), 0)), nil
	case 2:
		return fmt.Sprintf("Solution: %d", ""), nil
	default:
		return "", fmt.Errorf("unknown part %d", part)
	}
}

func numRedistributions(banks []int, seen *set.Set, cycle int) int {

	key := fmt.Sprint(banks)
	if seen.Contains(key) {
		return cycle
	}

	seen.Add(key)
	redis := redistribute(banks)
	return numRedistributions(redis, seen, cycle+1)
}

func redistribute(banks []int) []int {
	l := len(banks)
	redis := make([]int, l)
	idx, largest := largest(banks)
	banks[idx] = 0
	equalShare := largest / l
	for i := range l {
		redis[i] = banks[i] + equalShare
	}
	rest := largest % l
	for j := 1; j < rest+1; j++ {
		redis[(idx+j)%l] += 1
	}
	return redis
}

func largest(list []int) (idx, val int) {
	idx, val = 0, list[0]
	for i := 1; i < len(list); i++ {
		if list[i] > val {
			idx, val = i, list[i]
		}
	}
	return idx, val
}
