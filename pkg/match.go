package pkg

import (
	"fmt"
	"time"
)

type Match struct {
	Date          time.Time
	HomeTeamName  string
	HomeTeamScore int
	AwayTeamName  string
	AwayTeamScore int
}

func (m *Match) Init(date time.Time, homeTeamName string, homeTeamScore int, awayTeamName string, awayTeamScore int) {
	m.Date = date
	m.HomeTeamName = homeTeamName
	m.HomeTeamScore = homeTeamScore
	m.AwayTeamName = awayTeamName
	m.AwayTeamScore = awayTeamScore
}

func (m *Match) IsHomeTeam(teamName string) bool {
	return teamName == m.HomeTeamName
}

func (m *Match) IsAwayTeam(teamName string) bool {
	return teamName == m.AwayTeamName
}

func (m *Match) Contains(teamName string) bool {
	return m.IsHomeTeam(teamName) || m.IsAwayTeam(teamName)
}

func (m *Match) IsDraw() bool {
	return m.HomeTeamScore == m.AwayTeamScore
}

func (m *Match) IsWinner(teamName string) bool {
	if !m.Contains(teamName) {
		return false
	}

	if m.IsDraw() {
		return false
	}

	answer := false

	if m.HomeTeamName == teamName {
		answer = m.HomeTeamScore > m.AwayTeamScore
	} else if m.AwayTeamName == teamName {
		answer = m.AwayTeamScore > m.HomeTeamScore
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

	if m.HomeTeamName == teamName {
		answer = m.HomeTeamScore < m.AwayTeamScore
	} else if m.AwayTeamName == teamName {
		answer = m.AwayTeamScore < m.HomeTeamScore
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
	if !m.Contains(teamName) {
		return "", fmt.Errorf("the match doesn't contain the team <%s>", teamName)
	}

	if m.HomeTeamName == teamName {
		return m.AwayTeamName, nil
	}

	return m.HomeTeamName, nil
}
