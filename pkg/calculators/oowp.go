package calculators

import "github.com/jedi-knights/rpi/pkg"

type OOWPCalculator struct {
	matches []*pkg.Match
}

func NewOOWPCalculator(matches []*pkg.Match) *OOWPCalculator {
	return &OOWPCalculator{
		matches: matches,
	}
}

func (w *OOWPCalculator) Calculate(teamName string) (float64, error) {
	return 0.0, nil
}
