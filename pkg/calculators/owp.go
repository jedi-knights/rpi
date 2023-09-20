package calculators

import (
	"github.com/jedi-knights/rpi/pkg/accumulators"
	"github.com/jedi-knights/rpi/pkg/match"
)

// OWPCalculator is a calculator that calculates the OWP for a given team.
type OWPCalculator struct{}

// NewOWPCalculator returns a new OWP calculator.
func NewOWPCalculator() *OWPCalculator {
	return &OWPCalculator{}
}

// Calculate returns the specified team's opponents' average winning percentage.
// To determine Team A's opponents' average winning percentage, the NCAA first computes, for each of Team A's
// opponents, the opponent's wins and ties as compared to the opponent's total games played, in the same way it
// does the calculation for Team A's Element 1.  The only difference is that the NCAA excludes the opponent's
// games against Team A itself.  Thus this first part of the computation determines each opponent's Element 1
// based on games played against teams other than Team A.
//
// The OWP is calculated by taking the average of the WP's for each of the team's opponentNames with the
// requirement that all games against the team in question are removed from the equation.
func (w *OWPCalculator) Calculate(targetTeamName string, matches *[]match.Match) (float64, error) {

	// for each fo Team A's opponents, compute the opponent's wins and ties as compared to the opponent's total
	// games played, excluding games against Team A
	opponentsNames, err := match.GetOpponents(matches, targetTeamName)
	if err != nil {
		return 0.0, err
	}

	// for each opponent, compute the opponent's wins and ties as compared to the opponent's total games played,
	opponentWinningPercentageMap := make(map[string]float64)
	for _, opponentName := range opponentsNames {
		// compute the wins for this opponent
		winsAccumulator := accumulators.NewWins(targetTeamName)
		opponentWins, err := winsAccumulator.Calculate(opponentName, matches)
		if err != nil {
			return 0.0, err
		}

		// compute the losses for this opponent
		lossesAccumulator := accumulators.NewLosses(targetTeamName)
		opponentLosses, err := lossesAccumulator.Calculate(opponentName, matches)
		if err != nil {
			return 0.0, err
		}

		// compute the ties for this opponent
		tiesAccumulator := accumulators.NewTies(targetTeamName)
		opponentTies, err := tiesAccumulator.Calculate(opponentName, matches)
		if err != nil {
			return 0.0, err
		}

		// determine the total games played by this opponent
		totalMatchesPlayed := opponentWins + opponentLosses + opponentTies

		// determine the opponent's winning percentage
		opponentWP := (float64(opponentWins) + 0.5*float64(opponentTies)) / float64(totalMatchesPlayed)

		// store the opponent's winning percentage
		opponentWinningPercentageMap[opponentName] = opponentWP
	}

	// at this point we should have the winning percentage for each of Team A's opponents
	// now we need to determine the average of these winning percentages

	// determine the sum of the winning percentages
	// Note: this calculation is a bit tricky because it depends on the number of meetings between the target team
	// and the opponent.  For example, if Team A has played Team B 3 times and Team C 2 times, then the sum of the
	// winning percentages is:
	//   (Team B's WP * 3) + (Team C's WP * 2)
	// but you also have to keep track of the total matches played with respect to each opponent.  For example, if
	// Team B has played 3 matches against Team A and 2 matches against Team C, then the total matches played is:
	//   3 + 2 = 5
	var sum float64
	numberOfMatches := 0
	for opponentName, wp := range opponentWinningPercentageMap {
		meetingCount, err := match.GetMeetingCount(matches, targetTeamName, opponentName)
		numberOfMatches += meetingCount
		if err != nil {
			return 0.0, err
		}

		sum += wp * float64(meetingCount)
	}

	// determine the average of the winning percentages
	average := sum / float64(numberOfMatches)

	return average, nil
}
