package calculators

import "github.com/jedi-knights/rpi/pkg/match"

type RPICalculator struct {
	WPCalculator   *WPCalculator
	OWPCalculator  *OWPCalculator
	OOWPCalculator *OOWPCalculator
}

func NewRPICalculator() *RPICalculator {
	return &RPICalculator{
		WPCalculator:   NewWPCalculator(""),
		OWPCalculator:  NewOWPCalculator(),
		OOWPCalculator: NewOOWPCalculator(),
	}
}

func (r *RPICalculator) Calculate(targetTeamName string, matches *[]match.Match) (float64, error) {
	var err error
	var wp, owp, oowp float64

	// get the team's winning percentage
	if wp, err = r.WPCalculator.Calculate(targetTeamName, matches); err != nil {
		return 0.0, err
	}

	// get the team's opponents' winning percentage
	if owp, err = r.OWPCalculator.Calculate(targetTeamName, matches); err != nil {
		return 0.0, err
	}

	// get the team's opponents' opponents' winning percentage
	if oowp, err = r.OOWPCalculator.Calculate(targetTeamName, matches); err != nil {
		return 0.0, err
	}

	// calculate the team's RPI
	rpi := 0.25*wp + 0.50*owp + 0.25*oowp

	return rpi, nil
}
