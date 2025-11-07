package d03

import (
	"fmt"
	"io"
	"math"
	"slices"

	"github.com/ruben-vl/aoc-go/internal/solvers"
)

func init() {
	solvers.Register("y2017d03", Solve)
}

func Solve(part int, r io.Reader) (string, error) {

	input := 361527

	switch part {
	case 1:
		circle := isWithinCircle(input)
		dist := distFromStraightLine(circle, input)
		return fmt.Sprintf("Steps required: %d", (circle-1)/2+dist), nil
	case 2:
		return "", nil
	default:
		return "", fmt.Errorf("unknown part %d", part)
	}
}

func isWithinCircle(i int) int {
	n := 1
	for n*n < i {
		n += 2
	}
	return n
}

func distFromStraightLine(c int, i int) int {
	fi := float64(i)
	distances := []float64{
		math.Abs(float64(c*c-(0*(c-1)+c/2)) - fi),
		math.Abs(float64(c*c-(1*(c-1)+c/2)) - fi),
		math.Abs(float64(c*c-(2*(c-1)+c/2)) - fi),
		math.Abs(float64(c*c-(3*(c-1)+c/2)) - fi),
	}
	return int(slices.Min(distances))
}
