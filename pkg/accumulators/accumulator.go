package accumulators

import (
	"github.com/jedi-knights/rpi/pkg/match"
)

// AccumulatorInterface is an interface that all accumulators must implement.
type AccumulatorInterface interface {
	Calculate(teamName string, matches *[]match.Match) (int, error)
}
