package calculators

import "github.com/jedi-knights/rpi/pkg"

type OOWPCalculator struct{}

func NewOOWPCalculator() *OOWPCalculator {
	return &OOWPCalculator{}
}

func (w *OOWPCalculator) Calculate(teamName string, matches *[]pkg.Match) (float64, error) {
	return 0.0, nil
}
