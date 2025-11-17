package d07

import (
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/ruben-vl/aoc-go/internal/solvers"
	"github.com/ruben-vl/aoc-go/internal/utils/input"
)

func init() {
	solvers.Register("y2017d07", Solve)
}

func Solve(part int, r io.Reader) (string, error) {

	lines, err := input.ReadLines(r)
	if err != nil {
		return "", err
	}

	tower := supports(lines)
	weights := weights(lines)

	switch part {
	case 1:
		return fmt.Sprintf("The bottom program is named %q", bottom(tower)), nil
	case 2:
		unbalancedSubTower(bottom(tower), 1, &tower, &weights)
		return fmt.Sprintln("The solution is:"), nil
	default:
		return "", fmt.Errorf("unknown part %d", part)
	}
}

func supports(input []string) (out map[string][]string) {

	out = make(map[string][]string, len(input))

	for _, line := range input {
		lineElements := strings.Split(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(
						strings.ReplaceAll(
							line,
							")", ""),
						"(", ""),
					" -> ", " "),
				",", ""),
			" ")
		name := lineElements[0]
		if len(lineElements) > 2 {
			out[name] = lineElements[2:]
		} else {
			out[name] = make([]string, 0)
		}
	}

	return
}

func bottom(tower map[string][]string) string {
	supported := make([]string, 0, len(tower))
	for _, above := range tower {
		supported = append(supported, above...)
	}
	for key := range tower {
		if !slices.Contains(supported, key) {
			return key
		}
	}
	return ""
}

func weights(input []string) (out map[string]int) {

	out = make(map[string]int, len(input))

	for _, line := range input {
		lineElements := strings.Split(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(
						strings.ReplaceAll(
							line,
							")", ""),
						"(", ""),
					" -> ", " "),
				",", ""),
			" ")
		weight, _ := strconv.Atoi(lineElements[1])
		out[lineElements[0]] = weight
	}

	return
}

func unbalancedSubTower(start string, depth int, tower *map[string][]string, weights *map[string]int) {
	subTowerWeights := make([]int, 0, len((*tower)[start]))
	for _, subTower := range (*tower)[start] {
		subTowerWeights = append(subTowerWeights, towerWeight(subTower, tower, weights))
	}
	if allEqual(subTowerWeights) {
		// fmt.Printf("%v: BALANCED found with weights %v\n", strings.Repeat("=", depth), subTowerWeights)
		return
	} else {
		fmt.Printf("%v: unbalanced towers found with weights %v\n", strings.Repeat("=", depth), subTowerWeights)
		for _, subTower := range (*tower)[start] {
			fmt.Printf("%v: %v\n", subTower, (*weights)[subTower])
		}
		for _, subTower := range (*tower)[start] {
			unbalancedSubTower(subTower, depth+1, tower, weights)
		}
	}
}

func towerWeight(start string, tower *map[string][]string, weights *map[string]int) (total int) {
	total = (*weights)[start]
	for _, subTower := range (*tower)[start] {
		total += towerWeight(subTower, tower, weights)
	}
	return total
}

func allEqual(ints []int) bool {
	if len(ints) < 2 {
		return true
	}
	for _, v := range ints {
		if v != ints[0] {
			return false
		}
	}
	return true
}
