package d01

import (
	"fmt"
	"io"

	"github.com/ruben-vl/aoc-go/internal/solvers"
	"github.com/ruben-vl/aoc-go/internal/utils/input"
	"github.com/ruben-vl/aoc-go/internal/utils/stringconv"
)

func init() {
	solvers.Register("y2017d01", Solve)
}

func Solve(part int, r io.Reader) (string, error) {

	line, err := input.ReadSingleLine(r)
	if err != nil {
		return "", err
	}
	nums := stringconv.DigitsFromString(line)

	switch part {
	case 1:
		return fmt.Sprintf("Sum of consecutive equal values: %d", sumConsecutiveEqualValues(nums)), nil
	case 2:
		return fmt.Sprintf("Sum of opposite equal values: %d", sumOppositeEqualValues(nums)), nil
	default:
		return "", fmt.Errorf("unknown part %d", part)
	}
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
