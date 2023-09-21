package schedule

import (
	"fmt"
	. "github.com/jedi-knights/rpi/pkg/match"
	"slices"
)

type ISchedule interface {
	AddMatch(match *Match)
	GetMatches() []*Match
	GetMatchesForTeam(teamName string) []*Match
	GetOpponents(teamName string) ([]string, error)
	GetOpponentsForTeam(teamName string) ([]string, error)
	GetMatchesPlayedBy(teamName string) ([]*Match, error)
	GetWinsForTeam(teamName, skipTeamName string) (int, error)
	GetLossesForTeam(teamName, skipTeamName string) (int, error)
	GetTiesForTeam(teamName, skipTeamName string) (int, error)
	GetTotalMatchesPlayedForTeam(teamName string) (int, error)
	GetTotalMatchesPlayed() int
	CalculateWP(teamName string) (float64, error)
	CalculateOWP(teamName string) (float64, error)
	CalculateOOWP(teamName string) (float64, error)
	CalculateRPI(teamName string) (float64, error)
}

type Schedule struct {
	matches []*Match
}

func NewSchedule() *Schedule {
	return &Schedule{
		matches: make([]*Match, 0),
	}
}

func (s *Schedule) AddMatch(match *Match) {
	s.matches = append(s.matches, match)
}

func (s *Schedule) AddMatchFromString(matchString string) {
	s.AddMatch(NewMatchFromString(matchString))
}

func (s *Schedule) GetMatches() []*Match {
	return s.matches
}

func (s *Schedule) GetMatchesForTeam(teamName string) []*Match {
	var matches []*Match

	for _, match := range s.matches {
		if match.Contains(teamName) {
			matches = append(matches, match)
		}
	}

	return matches
}

func (s *Schedule) GetOpponents(teamName string) ([]string, error) {
	var opponents []string

	if teamName == "" {
		return nil, fmt.Errorf("the specified team name is empty")
	}

	for _, match := range s.matches {
		var err error
		var opponentName string

		if opponentName, err = match.GetOpponent(teamName); err != nil {
			continue
		}

		if !slices.Contains(opponents, opponentName) {
			opponents = append(opponents, opponentName)
		}
	}

	return opponents, nil
}

func (s *Schedule) GetMatchesPlayedBy(teamName string) ([]*Match, error) {
	var matchesPlayed []*Match

	if teamName == "" {
		return nil, fmt.Errorf("the specified team name is empty")
	}

	for _, match := range s.matches {
		if match.Contains(teamName) {
			matchesPlayed = append(matchesPlayed, match)
		}
	}

	if len(matchesPlayed) == 0 {
		return nil, fmt.Errorf("no matches found for team %s", teamName)
	}

	return matchesPlayed, nil
}

func (s *Schedule) Contains(teamName string) bool {
	for _, match := range s.matches {
		if match.Contains(teamName) {
			return true
		}
	}

	return false
}

func (s *Schedule) GetWinsForTeam(teamName, skipTeamName string) (int, error) {
	var total int

	if teamName == "" {
		return 0, fmt.Errorf("the specified team name is empty")
	}

	if !s.Contains(teamName) {
		return 0, fmt.Errorf("no matches found for team %s", teamName)
	}

	for _, match := range s.matches {
		if !match.Contains(teamName) {
			continue
		}

		if len(skipTeamName) > 0 && match.Contains(skipTeamName) {
			continue
		}

		if match.IsWinner(teamName) {
			total++
		}
	}

	return total, nil
}

func (s *Schedule) GetLossesForTeam(teamName, skipTeamName string) (int, error) {
	var total int

	if teamName == "" {
		return 0, fmt.Errorf("the specified team name is empty")
	}

	if !s.Contains(teamName) {
		return 0, fmt.Errorf("no matches found for team %s", teamName)
	}

	for _, match := range s.matches {
		if !match.Contains(teamName) {
			continue
		}

		if len(skipTeamName) > 0 && match.Contains(skipTeamName) {
			continue
		}

		if match.IsLoser(teamName) {
			total++
		}
	}

	return total, nil
}

func (s *Schedule) GetTiesForTeam(teamName, skipTeamName string) (int, error) {
	var total int

	if teamName == "" {
		return 0, fmt.Errorf("the specified team name is empty")
	}

	if !s.Contains(teamName) {
		return 0, fmt.Errorf("no matches found for team %s", teamName)
	}

	for _, match := range s.matches {
		if !match.Contains(teamName) {
			continue
		}

		if len(skipTeamName) > 0 && match.Contains(skipTeamName) {
			continue
		}

		if match.IsDraw() {
			total++
		}
	}

	return total, nil
}

func (s *Schedule) GetTotalMatchesPlayedForTeam(teamName string) (int, error) {
	var totalMatchesPlayed int

	if teamName == "" {
		return 0, fmt.Errorf("the specified team name is empty")
	}

	found := false
	for _, match := range s.matches {
		found = found || match.Contains(teamName)
		if match.Contains(teamName) {
			totalMatchesPlayed++
		}
	}

	if !found {
		return 0, fmt.Errorf("no matches found for team %s", teamName)
	}

	return totalMatchesPlayed, nil
}

func (s *Schedule) GetTotalMatchesPlayed() int {
	return len(s.matches)
}

func (s *Schedule) CalculateWP(teamName, skipTeamName string) (float64, error) {
	var winningPercentage float64

	if teamName == "" {
		return 0.0, fmt.Errorf("the specified team name is empty")
	}

	wins, err := s.GetWinsForTeam(teamName, skipTeamName)
	if err != nil {
		return 0.0, err
	}

	losses, err := s.GetLossesForTeam(teamName, skipTeamName)
	if err != nil {
		return 0.0, err
	}

	ties, err := s.GetTiesForTeam(teamName, skipTeamName)
	if err != nil {
		return 0.0, err
	}

	totalMatchesPlayed := wins + losses + ties

	winningPercentage = float64(wins+(ties/2)) / float64(totalMatchesPlayed)

	return winningPercentage, nil
}

func (s *Schedule) GetMeetingCount(teamA, teamB string) (int, error) {
	var meetingCount int

	if teamA == "" {
		return 0, fmt.Errorf("the first specified team name is empty")
	}

	if teamB == "" {
		return 0, fmt.Errorf("the second specified team name is empty")
	}

	for _, currentMatch := range s.matches {
		if !currentMatch.Contains(teamA) {
			continue
		}

		if !currentMatch.Contains(teamB) {
			continue
		}

		meetingCount++
	}

	return meetingCount, nil
}

func (s *Schedule) CalculateOWP(targetTeamName string) (float64, error) {
	if targetTeamName == "" {
		return 0.0, fmt.Errorf("the specified team name is empty")
	}

	if !s.Contains(targetTeamName) {
		return 0.0, fmt.Errorf("no matches found for team %s", targetTeamName)
	}

	// for each fo Team A's opponents, compute the opponent's wins and ties as compared to the opponent's total
	// games played, excluding games against Team A
	opponentsNames, err := s.GetOpponents(targetTeamName)
	if err != nil {
		return 0.0, err
	}

	// for each opponent, compute the opponent's wins and ties as compared to the opponent's total games played,
	opponentWinningPercentageMap := make(map[string]float64)
	for _, opponentName := range opponentsNames {
		// compute the wins for this opponent
		opponentWins, err := s.GetWinsForTeam(opponentName, targetTeamName)
		if err != nil {
			return 0.0, err
		}

		// compute the losses for this opponent
		opponentLosses, err := s.GetLossesForTeam(opponentName, targetTeamName)
		if err != nil {
			return 0.0, err
		}

		// compute the ties for this opponent
		opponentTies, err := s.GetTiesForTeam(opponentName, targetTeamName)
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
		meetingCount, err := s.GetMeetingCount(targetTeamName, opponentName)
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

// CalculateOOWP calculates the opponent's opponent's winning percentage for the specified team.
// The opponent's opponent's winning percentage is the average of the winning percentages of all of the
// opponents of the specified team.
func (s *Schedule) CalculateOOWP(teamName string) (float64, error) {
	var err error
	var owp float64

	if teamName == "" {
		return 0.0, fmt.Errorf("the specified team name is empty")
	}

	if !s.Contains(teamName) {
		return 0.0, fmt.Errorf("no matches found for team %s", teamName)
	}

	opponents, err := s.GetOpponents(teamName)
	if err != nil {
		return 0.0, err
	}

	opponentsOWPMap := make(map[string]float64)
	for _, opponent := range opponents {
		if owp, err = s.CalculateOWP(opponent); err != nil {
			return 0.0, err
		}

		opponentsOWPMap[opponent] = owp
	}

	var accumulator float64
	var numberOfMatches int

	for opponent, owp := range opponentsOWPMap {
		var numberOfMeetings int
		if numberOfMeetings, err = s.GetMeetingCount(teamName, opponent); err != nil {
			return 0.0, err
		}

		numberOfMatches += numberOfMeetings
		accumulator += owp * float64(numberOfMeetings)
	}

	average := accumulator / float64(numberOfMatches)

	return average, nil
}

func (s *Schedule) CalculateRPI(teamName string) (float64, error) {
	var err error
	var wp, owp, oowp, rpi float64

	if wp, err = s.CalculateWP(teamName, ""); err != nil {
		return 0.0, err
	}

	if owp, err = s.CalculateOWP(teamName); err != nil {
		return 0.0, err
	}

	if oowp, err = s.CalculateOOWP(teamName); err != nil {
		return 0.0, err
	}

	rpi = (wp + (float64(2) * owp) + oowp) / float64(4)

	return rpi, nil
}
