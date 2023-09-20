package match

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Factory is a factory for creating matches.
type Factory struct {
	builder *Builder
}

// NewFactory creates a new match factory.
func NewFactory(builder *Builder) *Factory {
	return &Factory{
		builder: builder,
	}
}

// Create creates a new match with the given parameters.
func (m *Factory) Create(date time.Time, homeName string, homeScore int, awayName string, awayScore int) *Match {
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

func (m *Factory) CreateWithRandomDate(homeName string, homeScore int, awayName string, awayScore int) *Match {
	return m.Create(randate(), homeName, homeScore, awayName, awayScore)
}

func (m *Factory) CreateFromString(input string) *Match {
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
