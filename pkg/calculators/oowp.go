package calculators

import (
	"github.com/jedi-knights/rpi/pkg/match"
)

type OOWPCalculator struct{}

func NewOOWPCalculator() *OOWPCalculator {
	return &OOWPCalculator{}
}

func (w *OOWPCalculator) Calculate(targetTeamName string, matches *[]match.Match) (float64, error) {
	opponentsNames, err := match.GetOpponents(matches, targetTeamName)
	if err != nil {
		return 0.0, err
	}

	owpMap := make(map[string]float64)

	for _, opponentName := range opponentsNames {
		// get the opponent's winning percentage
		owpCalculator := NewOWPCalculator()
		owp, err := owpCalculator.Calculate(opponentName, matches)
		if err != nil {
			return 0.0, err
		}

		owpMap[opponentName] = owp
	}

	// calculate the average of the opponent's winning percentages
	var oowp float64

	for _, owp := range owpMap {
		oowp += owp
	}

	oowp /= float64(len(owpMap))

	return oowp, nil
}
