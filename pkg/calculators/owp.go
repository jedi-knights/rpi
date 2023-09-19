package calculators

import (
	"fmt"
	"github.com/jedi-knights/rpi/pkg"
)

// OWPCalculator is a calculator that calculates the OWP for a given team.
type OWPCalculator struct{}

// NewOWPCalculator returns a new OWP calculator.
func NewOWPCalculator() *OWPCalculator {
	return &OWPCalculator{}
}

// Calculate returns the OWP for the given team.
func (w *OWPCalculator) Calculate(teamName string, matches *[]pkg.Match) (float64, error) {
	var err error
	var opponents []string

	if teamName == "" {
		return 0.0, fmt.Errorf(pkg.ErrTeamNameRequired)
	}

	if opponents, err = pkg.GetOpponents(matches, teamName); err != nil {
		return 0.0, err
	}

	wpCalculator := NewWPCalculator(teamName)

	accumulator := float64(0.0)
	for _, opponentName := range opponents {
		var wp float64

		if wp, err = wpCalculator.Calculate(opponentName, matches); err != nil {
			return 0.0, err
		}

		accumulator += wp
	}

	owp := accumulator / float64(len(opponents))

	return owp, nil
}
