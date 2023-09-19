package pkg

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// MatchFactory is a factory for creating matches.
type MatchFactory struct {
	builder *MatchBuilder
}

// NewMatchFactory creates a new match factory.
func NewMatchFactory(builder *MatchBuilder) *MatchFactory {
	return &MatchFactory{
		builder: builder,
	}
}

// Create creates a new match with the given parameters.
func (m *MatchFactory) Create(date time.Time, homeName string, homeScore int, awayName string, awayScore int) *Match {
	return m.builder.
		BuildDate(date).
		BuildHomeName(homeName).
		BuildHomeScore(homeScore).
		BuildAwayName(awayName).
		BuildAwayScore(awayScore).
		GetInstance()
}

func randate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func (m *MatchFactory) CreateWithRandomDate(homeName string, homeScore int, awayName string, awayScore int) *Match {
	return m.Create(randate(), homeName, homeScore, awayName, awayScore)
}

func (m *MatchFactory) CreateFromString(input string) *Match {
	tokens := strings.Split(input, ",")

	date, err := time.Parse("2006-01-02", tokens[0])
	if err != nil {
		return nil
	}

	homeName := tokens[1]

	homeScore, err := strconv.Atoi(tokens[2])
	if err != nil {
		return nil
	}

	awayName := tokens[3]

	awayScore, err := strconv.Atoi(tokens[4])
	if err != nil {
		return nil
	}

	return m.Create(date, homeName, homeScore, awayName, awayScore)
}
