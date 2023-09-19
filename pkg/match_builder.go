package pkg

import "time"

type MatchBuilder struct {
	date      time.Time
	homeName  string
	homeScore int
	awayName  string
	awayScore int
}

func NewMatchBuilder() *MatchBuilder {
	return &MatchBuilder{
		date:      time.Now(),
		homeName:  "",
		homeScore: 0,
		awayName:  "",
		awayScore: 0,
	}
}

func (m *MatchBuilder) BuildDate(date time.Time) *MatchBuilder {
	m.date = date
	return m
}

func (m *MatchBuilder) BuildHomeName(homeName string) *MatchBuilder {
	m.homeName = homeName
	return m
}

func (m *MatchBuilder) BuildHomeScore(homeScore int) *MatchBuilder {
	m.homeScore = homeScore
	return m
}

func (m *MatchBuilder) BuildAwayName(awayName string) *MatchBuilder {
	m.awayName = awayName
	return m
}

func (m *MatchBuilder) BuildAwayScore(awayScore int) *MatchBuilder {
	m.awayScore = awayScore
	return m
}

func (m *MatchBuilder) GetInstance() *Match {
	match := NewMatch()

	match.Date = m.date
	match.Home.Name = m.homeName
	match.Home.Score = m.homeScore
	match.Away.Name = m.awayName
	match.Away.Score = m.awayScore

	return match
}
