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

//func ProcessSomething(targetTeamName, opponentName string, matches *[]match.Match) (float64, error) {
//	var err error
//	var OW, OL, OT int
//
//	winsAccumulator := accumulators.NewWins(targetTeamName)
//	lossesAccumulator := accumulators.NewLosses(targetTeamName)
//	tiesAccumulator := accumulators.NewTies(targetTeamName)
//
//	if OW, err = winsAccumulator.Calculate(opponentName, matches); err != nil {
//		return 0.0, err
//	}
//
//	if OL, err = lossesAccumulator.Calculate(opponentName, matches); err != nil {
//		return 0.0, err
//	}
//
//	if OT, err = tiesAccumulator.Calculate(opponentName, matches); err != nil {
//		return 0.0, err
//	}
//
//	// retrieve a subslice of matches played between the target team and the current opponent
//	pSubslice, err := match.GetMatchesBetween(matches, targetTeamName, opponentName)
//	if err != nil {
//		return 0.0, err
//	}
//
//	for _, meeting := range *pSubslice {
//		var portion float64
//
//		if meeting.IsWinner(targetTeamName) {
//			// the target team won the match
//			portion = (float64(OW) + 0.5*float64(OT)) / float64(OW+(OL-1)+OT)
//		} else if meeting.IsDraw() {
//			// the match was a draw
//			portion = (float64(OW) + 0.5*float64(OT-1)) / float64(OW+OL+(OT-1))
//		} else {
//			// the target team lost the match
//			portion = (float64(OW-1) + 0.5*float64(OT)) / float64((OW-1)+OL+OT)
//		}
//
//		// portions = append(portions, portion)
//	}
//}

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
