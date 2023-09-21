package match

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Match struct {
	Date time.Time
	Home Status
	Away Status
}

func NewMatch() *Match {
	return &Match{
		Date: time.Now(),
		Home: Status{
			Name:  "",
			Score: 0,
		},
		Away: Status{
			Name:  "",
			Score: 0,
		},
	}
}

func NewMatchFromString(matchString string) *Match {
	var err error
	tokens := strings.Split(matchString, ",")

	if len(tokens) == 5 {
		newMatch := NewMatch()

		newMatch.Home.Name = tokens[1]
		newMatch.Away.Name = tokens[3]

		if newMatch.Date, err = time.Parse("2006-01-02", tokens[0]); err != nil {
			return nil
		}
		if newMatch.Home.Score, err = strconv.Atoi(tokens[2]); err != nil {
			return nil
		}
		if newMatch.Away.Score, err = strconv.Atoi(tokens[4]); err != nil {
			return nil
		}

		return newMatch
	} else if len(tokens) == 4 {
		newMatch := NewMatch()

		newMatch.Date = time.Now()

		newMatch.Home.Name = tokens[0]
		newMatch.Away.Name = tokens[2]

		if newMatch.Home.Score, err = strconv.Atoi(tokens[1]); err != nil {
			return nil
		}
		if newMatch.Away.Score, err = strconv.Atoi(tokens[3]); err != nil {
			return nil
		}

		return newMatch
	}

	return nil
}

func (m *Match) IsHomeTeam(teamName string) bool {
	return teamName == m.Home.Name
}

func (m *Match) IsAwayTeam(teamName string) bool {
	return teamName == m.Away.Name
}

func (m *Match) Contains(teamName string) bool {
	return m.IsHomeTeam(teamName) || m.IsAwayTeam(teamName)
}

func (m *Match) IsDraw() bool {
	return m.Home.Score == m.Away.Score
}

func (m *Match) IsWinner(teamName string) bool {
	if !m.Contains(teamName) {
		return false
	}

	if m.IsDraw() {
		return false
	}

	answer := false

	if m.Home.Name == teamName {
		answer = m.Home.Score > m.Away.Score
	} else if m.Away.Name == teamName {
		answer = m.Away.Score > m.Home.Score
	}

	return answer
}

func (m *Match) IsLoser(teamName string) bool {
	if !m.Contains(teamName) {
		return false
	}

	if m.IsDraw() {
		return false
	}

	answer := false

	if m.Home.Name == teamName {
		answer = m.Home.Score < m.Away.Score
	} else if m.Away.Name == teamName {
		answer = m.Away.Score < m.Home.Score
	}

	return answer
}

func (m *Match) WinValue(teamName string) float64 {
	if !m.Contains(teamName) {
		return 0.0
	}

	if m.IsDraw() {
		return 0.5
	}

	if m.IsWinner(teamName) {
		return 1.0
	}

	return 0.0
}

func (m *Match) GetOpponent(teamName string) (string, error) {
	if teamName == "" {
		return "", fmt.Errorf("the specified team name is empty")
	}

	if !m.Contains(teamName) {
		return "", fmt.Errorf("the match doesn't contain the team <%s>", teamName)
	}

	if m.Home.Name == teamName {
		return m.Away.Name, nil
	}

	return m.Home.Name, nil
}

func (m *Match) ToString() string {
	return fmt.Sprintf("%s,%d,%s,%d", m.Home.Name, m.Home.Score, m.Away.Name, m.Away.Score)
}

func (m *Match) ToFullString() string {
	return fmt.Sprintf("[%s] %d, [%s] %d", m.Home.Name, m.Home.Score, m.Away.Name, m.Away.Score)
}

// GetOpponents returns the opponents for a given team.
func GetOpponents(matches *[]Match, teamName string) ([]string, error) {
	var opponents []string

	if teamName == "" {
		return nil, fmt.Errorf("the specified team name is empty")
	}

	for _, match := range *matches {
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

// GetMatchesPlayedBy returns the matches played by a given team.
func GetMatchesPlayedBy(matches *[]Match, teamName string) (*[]Match, error) {
	var matchesPlayed []Match

	if teamName == "" {
		return nil, fmt.Errorf("the specified team name is empty")
	}

	for _, match := range *matches {
		if match.Contains(teamName) {
			matchesPlayed = append(matchesPlayed, match)
		}
	}

	return &matchesPlayed, nil
}

// GetMatchesBetween returns the matches played between two teams.
func GetMatchesBetween(matches *[]Match, teamA string, teamB string) (*[]Match, error) {
	var matchesPlayed []Match

	if teamA == "" {
		return nil, fmt.Errorf("the first specified team name is empty")
	}

	if teamB == "" {
		return nil, fmt.Errorf("the second specified team name is empty")
	}

	matchesPlayed = make([]Match, 0)

	for _, currentMatch := range *matches {
		if !currentMatch.Contains(teamA) {
			continue
		}

		if !currentMatch.Contains(teamB) {
			continue
		}

		matchesPlayed = append(matchesPlayed, currentMatch)
	}

	return &matchesPlayed, nil
}

// GetMeetingCount returns the number of meetings between two teams.
func GetMeetingCount(matches *[]Match, teamA string, teamB string) (int, error) {
	var meetingCount int

	if teamA == "" {
		return 0, fmt.Errorf("the first specified team name is empty")
	}

	if teamB == "" {
		return 0, fmt.Errorf("the second specified team name is empty")
	}

	for _, currentMatch := range *matches {
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

// GetMatchesPlayedByOpponents returns the matches played by the opponents of a given team.
// This excludes all matches played by the given team.
func GetMatchesPlayedByOpponents(matches *[]Match, teamName string) (*[]Match, error) {
	var matchesPlayed []Match

	if teamName == "" {
		return nil, fmt.Errorf("the specified team name is empty")
	}

	opponents, err := GetOpponents(matches, teamName)
	if err != nil {
		return nil, err
	}

	if len(opponents) == 0 {
		return nil, fmt.Errorf("the specified opponents slice is empty")
	}

	for _, match := range *matches {
		if match.Contains(teamName) {
			continue
		}

		if !slices.Contains(opponents, match.Home.Name) && !slices.Contains(opponents, match.Away.Name) {
			continue
		}

		matchesPlayed = append(matchesPlayed, match)
	}

	return &matchesPlayed, nil
}
