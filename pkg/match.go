package pkg

import (
	"fmt"
	"time"
)

type Match struct {
	Date time.Time
	Home MatchStatus
	Away MatchStatus
}

func NewMatch() *Match {
	return &Match{
		Date: time.Now(),
		Home: MatchStatus{
			Name:  "",
			Score: 0,
		},
		Away: MatchStatus{
			Name:  "",
			Score: 0,
		},
	}
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

		opponents = append(opponents, opponentName)
	}

	return opponents, nil
}
