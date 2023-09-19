package calculators

import (
	"github.com/jedi-knights/rpi/pkg/match"
)

type OOWPCalculator struct{}

func NewOOWPCalculator() *OOWPCalculator {
	return &OOWPCalculator{}
}

func (w *OOWPCalculator) Calculate(teamName string, matches *[]match.Match) (float64, error) {
	return 0.0, nil
}
