package match

import "time"

type Builder struct {
	date      time.Time
	homeName  string
	homeScore int
	awayName  string
	awayScore int
}

func NewBuilder() *Builder {
	return &Builder{
		date:      time.Now(),
		homeName:  "",
		homeScore: 0,
		awayName:  "",
		awayScore: 0,
	}
}

func (m *Builder) BuildDate(date time.Time) *Builder {
	m.date = date
	return m
}

func (m *Builder) BuildHomeName(homeName string) *Builder {
	m.homeName = homeName
	return m
}

func (m *Builder) BuildHomeScore(homeScore int) *Builder {
	m.homeScore = homeScore
	return m
}

func (m *Builder) BuildAwayName(awayName string) *Builder {
	m.awayName = awayName
	return m
}

func (m *Builder) BuildAwayScore(awayScore int) *Builder {
	m.awayScore = awayScore
	return m
}

func (m *Builder) GetInstance() *Match {
	match := NewMatch()

	match.Date = m.date
	match.Home.Name = m.homeName
	match.Home.Score = m.homeScore
	match.Away.Name = m.awayName
	match.Away.Score = m.awayScore

	return match
}
