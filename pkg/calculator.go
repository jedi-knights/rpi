package pkg

import (
	"errors"
	"fmt"
	"github.com/jedi-knights/rpi/pkg/accumulators"
	"github.com/jedi-knights/rpi/pkg/calculators"
	"github.com/jedi-knights/rpi/pkg/match"
)

// https://sites.google.com/site/rpifordivisioniwomenssoccer/rpi-formula

/*
Ratings Percentage Index

measures a teams string relative to other teams based largely on the strength o ftheir schedules

The RPI is calculated from:
- the team's Winning Percentage (WP)
- the Opponent's Winning Percentage (OWP)
- the Opponent's Opponent's Winning Percentage (OOWP)

A win is assigned a value of 1.0 as a basis for comparison;
A tie is assigned a value of 0.5;
A loss is assigned a 0.0 value.
*/

type RPICalculator struct {
	WPVAL   float64
	OWPVAL  float64
	OOWPVAL float64
	Matches []match.Match
}

func (c *RPICalculator) Init(matches []match.Match) {
	c.WPVAL = 0.35
	c.OWPVAL = 0.35
	c.OOWPVAL = 0.30
	c.Matches = matches
}

func NewRPICalculator(matches []match.Match) (*RPICalculator, error) {
	calculator := &RPICalculator{}

	calculator.Init(matches)

	return calculator, nil
}

/*
WP is the winning percentage of the team being considered.
- The WP is computed by taking the number of wins and dividing by the number of games played.
*/
func (c *RPICalculator) CalculateWP(teamName, skipTeamName string) (float64, error) {
	wins, losses, draws, err := c.CalculateWinsLossesDraws(teamName, skipTeamName)
	if err != nil {
		return 0.0, err
	}

	matchesPlayed := wins + losses + draws
	if matchesPlayed == 0 {
		return 0.0, nil
	}

	wp := (float64(wins) + 0.5*float64(draws)) / float64(matchesPlayed)

	return wp, nil
}

// Returns wins, losses, draws, err
func (c *RPICalculator) CalculateWinsLossesDraws(teamName, skipTeamName string) (int, int, int, error) {
	wins := 0
	losses := 0
	draws := 0

	for _, match := range c.Matches {
		if !match.Contains(teamName) {
			continue
		}

		if len(skipTeamName) > 0 && match.Contains(skipTeamName) {
			continue
		}

		if match.IsDraw() {
			draws++
			continue
		}

		if match.IsWinner(teamName) {
			wins++
		} else {
			losses++
		}
	}

	return wins, losses, draws, nil
}

func (c *RPICalculator) CalculateMatchesPlayed(teamName string) (int, error) {
	count := 0
	for _, match := range c.Matches {
		if match.Contains(teamName) {
			count++
		}
	}

	return count, nil
}

func (c *RPICalculator) CalculateOWP(teamName string) (float64, error) {
	opponents := []string{}

	for _, match := range c.Matches {
		if !match.Contains(teamName) {
			continue
		}

		opponentName, err := match.GetOpponent(teamName)
		if err != nil {
			return 0.0, err
		}

		opponents = append(opponents, opponentName)
	}

	accumulator := float64(0.0)
	for _, opponentName := range opponents {
		wp, _ := c.CalculateWP(opponentName, teamName)

		accumulator += wp
	}

	owp := accumulator / float64(len(opponents))

	return owp, nil
}

func (c *RPICalculator) CalculateOOWP() (float32, error) {
	return 0.0, errors.New("todo")
}

func Accumulate(metric, teamName, skipTeamName string, matches *[]match.Match) (int, error) {
	switch metric {
	case "Wins":
		accumulator := accumulators.NewWins(skipTeamName)
		result, err := accumulator.Calculate(teamName, matches)
		if err != nil {
			return 0, err
		}

		return result, nil
	case "Losses":
		accumulator := accumulators.NewLosses(skipTeamName)
		result, err := accumulator.Calculate(teamName, matches)
		if err != nil {
			return 0, err
		}

		return result, nil
	case "Ties":
		accumulator := accumulators.NewTies(skipTeamName)
		result, err := accumulator.Calculate(teamName, matches)
		if err != nil {
			return 0, err
		}

		return result, nil
	default:
		return 0, fmt.Errorf("the metric <%s> is unsupported", metric)
	}
}

func Calculate(metric string, teamName string, matches *[]match.Match) (float64, error) {
	switch metric {
	case "WP":
		calculator := calculators.NewWPCalculator("")
		result, err := calculator.Calculate(teamName, matches)
		if err != nil {
			return 0.0, err
		}

		return result, nil
	case "OWP":
		calculator := calculators.NewOWPCalculator()
		result, err := calculator.Calculate(teamName, matches)
		if err != nil {
			return 0.0, err
		}

		return result, nil
	case "OOWP":
		calculator := calculators.NewOOWPCalculator()
		result, err := calculator.Calculate(teamName, matches)
		if err != nil {
			return 0.0, err
		}

		return result, nil
	case "RPI":
		return 0.0, errors.New("todo")
	default:
		return 0.0, fmt.Errorf("the metric <%s> is unsupported", metric)
	}
}
