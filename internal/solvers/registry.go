package solvers

import "io"

type Solver func(part int, r io.Reader) (string, error)

var Registry = map[string]Solver{}

func Register(key string, s Solver) {
	if _, exists := Registry[key]; exists {
		panic("solver already registered: " + key)
	}
	Registry[key] = s
}
