package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/ruben-vl/aoc-go/internal/solvers"
	_ "github.com/ruben-vl/aoc-go/internal/solvers/y2017/d01"
	_ "github.com/ruben-vl/aoc-go/internal/solvers/y2017/d02"
	_ "github.com/ruben-vl/aoc-go/internal/solvers/y2017/d04"
	_ "github.com/ruben-vl/aoc-go/internal/solvers/y2017/d05"
)

func main() {
	var year, day, part int
	flag.IntVar(&year, "year", 2015, "year in range [2015..2024]")
	flag.IntVar(&day, "day", 1, "day in range [1..25]")
	flag.IntVar(&part, "part", 1, "part in range [1..2]")
	flag.Parse()

	key := fmt.Sprintf("y%04dd%02d", year, day)

	solver, ok := solvers.Registry[key]
	if !ok {
		log.Fatalf("no solver registered for %s", key)
	}

	inputPath := filepath.Join("inputs", fmt.Sprintf("y%04d", year), fmt.Sprintf("d%02d.txt", day))
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("failed to open input file %q: %v", inputPath, err)
	}
	defer file.Close()

	runSolver(key, part, file, solver)
}

func runSolver(key string, part int, input io.Reader, solver solvers.Solver) {
	start := time.Now()
	answer, err := solver(part, input)
	elapsed := time.Since(start)

	if err != nil {
		log.Fatalf("solver failed: %v", err)
	}

	fmt.Printf("[%s] part %d: %s\n", key, part, answer)
	fmt.Fprintf(os.Stderr, "(elapsed: %s)\n", elapsed.Round(time.Millisecond))
}
