package pkg

import (
	"errors"
	"fmt"
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
	Matches []Match
}

func (c *RPICalculator) Init(matches []Match) {
	c.WPVAL = 0.35
	c.OWPVAL = 0.35
	c.OOWPVAL = 0.30
	c.Matches = matches
}

func NewRPICalculator(matches []Match) (*RPICalculator, error) {
	calculator := &RPICalculator{}

	calculator.Init(matches)

	return calculator, nil
}

/*
WP is the winning percentage of the team being considered.
- The WP is computed by taking the number of wins and dividing by the number of games played.
*/
func (c *RPICalculator) CalculateWP(teamName string, skipTeamName string) (float64, error) {
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
func (c *RPICalculator) CalculateWinsLossesDraws(teamName string, skipTeamName string) (int, int, int, error) {
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
			draws += 1
			continue
		}

		if match.IsWinner(teamName) {
			wins += 1
		} else {
			losses += 1
		}
	}

	return wins, losses, draws, nil
}

func (c *RPICalculator) CalculateMatchesPlayed(teamName string) (int, error) {
	count := 0
	for _, match := range c.Matches {
		if match.Contains(teamName) {
			count += 1
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

func (c *RPICalculator) Calculate(statistic string, teamName string) (float64, error) {
	if statistic == "WP" {
		wp, err := c.CalculateWP(teamName, "")

		if err != nil {
			return 0.0, err
		}

		return wp, nil
	} else if statistic == "OWP" {

	} else if statistic == "OOWP" {
	} else if statistic == "WINS" {

	} else {
		return 0.0, fmt.Errorf("the statistic <%s> is not supported", statistic)
	}

	return 0.0, nil
}
