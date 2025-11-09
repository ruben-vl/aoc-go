package d05

import (
	"fmt"
	"io"

	"github.com/ruben-vl/aoc-go/internal/solvers"
	"github.com/ruben-vl/aoc-go/internal/utils/input"
	"github.com/ruben-vl/aoc-go/internal/utils/stringconv"
)

func init() {
	solvers.Register("y2017d05", Solve)
}

func Solve(part int, r io.Reader) (string, error) {

	lines, err := input.ReadLines(r)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	nums, err := stringconv.NumbersFromStrings(lines)
	if err != nil {
		return "", fmt.Errorf("failed to parse numbers from file: %w", err)
	}
	flatNums := flatten(nums)

	switch part {
	case 1:
		return fmt.Sprintf("Number of steps to exit: %d", stepsToExit(flatNums)), nil
	case 2:
		return fmt.Sprintf("Number of steps to exit: %d", stepsToExitExtended(flatNums)), nil
	default:
		return "", fmt.Errorf("unknown part %d", part)
	}
}

func flatten(ii [][]int) []int {
	out := make([]int, 0, len(ii))
	for _, i := range ii {
		out = append(out, i...)
	}
	return out
}

func stepsToExit(nums []int) int {
	steps := 0
	pos := 0
	for pos >= 0 && pos < len(nums) {
		nums[pos] += 1
		pos += (nums[pos] - 1)
		steps += 1
	}
	return steps
}

func stepsToExitExtended(nums []int) int {
	steps := 0
	pos := 0
	for pos >= 0 && pos < len(nums) {
		if nums[pos] >= 3 {
			nums[pos] -= 1
			pos += (nums[pos] + 1)
			steps += 1
		} else {
			nums[pos] += 1
			pos += (nums[pos] - 1)
			steps += 1
		}
	}
	return steps
}
