package accumulators

import "github.com/jedi-knights/rpi/pkg"

// AccumulatorInterface is an interface that all accumulators must implement.
type AccumulatorInterface interface {
	Calculate(teamName string, matches *[]pkg.Match) (int, error)
}
