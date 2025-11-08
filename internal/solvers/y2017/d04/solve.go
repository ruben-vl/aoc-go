package d04

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/ruben-vl/aoc-go/internal/solvers"
	"github.com/ruben-vl/aoc-go/internal/utils/input"
	"github.com/ruben-vl/aoc-go/internal/utils/set"
	"github.com/ruben-vl/aoc-go/internal/utils/stringconv"
)

func init() {
	solvers.Register("y2017d04", Solve)
}

func Solve(part int, r io.Reader) (string, error) {

	lines, err := input.ReadLines(r)
	if err != nil {
		return "", err
	}
	ppws := stringconv.WordsFromStrings(lines)

	switch part {
	case 1:
		return fmt.Sprintf("The number of valid passphrases is %d", numValidPassphrases(ppws)), nil
	case 2:
		return fmt.Sprintf("The number of valid passphrases is %d", numValidPassphrasesExtended(ppws)), nil
	default:
		return "", fmt.Errorf("unknown part %d", part)
	}
}

func numValidPassphrases(ppws [][]string) int {
	n := 0
	for _, ppw := range ppws {
		if !hasDuplicate(ppw) {
			n += 1
		}
	}
	return n
}

func hasDuplicate(ss []string) bool {
	elements := set.NewSet()
	for _, s := range ss {
		if elements.Contains(s) {
			return true
		}
		elements.Add(s)
	}
	return false
}

func numValidPassphrasesExtended(ppws [][]string) int {
	n := 0
	for _, ppw := range ppws {
		sortedStrings := make([]string, 0, len(ppw))
		for _, s := range ppw {
			sortedStrings = append(sortedStrings, sortChars(s))
		}
		if !hasDuplicate(sortedStrings) {
			n += 1
		}
	}
	return n
}

func sortChars(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
