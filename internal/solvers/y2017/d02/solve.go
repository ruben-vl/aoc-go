package d02

import (
	"fmt"
	"io"
	"math"
	"sort"

	"github.com/ruben-vl/aoc-go/internal/solvers"
	"github.com/ruben-vl/aoc-go/internal/utils/input"
	"github.com/ruben-vl/aoc-go/internal/utils/stringconv"
)

func init() {
	solvers.Register("y2017d02", Solve)
}

func Solve(part int, r io.Reader) (string, error) {

	lines, err := input.ReadLines(r)
	if err != nil {
		return "", err
	}
	numRows := make([][]int, 0, len(lines))
	for _, line := range lines {
		nums, err := stringconv.NumbersFromString(line)
		if err != nil {
			return "", err
		}
		numRows = append(numRows, nums)
	}

	switch part {
	case 1:
		return fmt.Sprintf("The checksum for the spreadsheet is: %d", checksum(numRows)), nil
	case 2:
		return fmt.Sprintf("The sum of evenly divisible values is: %d", sumEvenlyDivisibleValues(numRows)), nil
	default:
		return "", fmt.Errorf("unknown part %d", part)
	}
}

func checksum(numRows [][]int) int {
	diffSum := 0
	for _, row := range numRows {
		r := append([]int(nil), row...) // copy to avoid mutating input
		sort.Ints(r)
		diffSum += r[len(r)-1] - r[0]
	}
	return diffSum
}

func sumEvenlyDivisibleValues(numRows [][]int) int {
	divSum := 0
	for _, row := range numRows {
		for i1, num1 := range row {
			found := false
			for i2, num2 := range row {
				if i1 != i2 {
					if math.Mod(float64(num1), float64(num2)) == 0.0 {
						divSum += num1 / num2
						found = true
						break
					}
				}
			}
			if found {
				break
			}
		}
	}
	return divSum
}
