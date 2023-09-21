package team

import "fmt"

type Team struct {
	Name   string
	Wins   int
	Losses int
	Ties   int
}

func NewTeam(name string) *Team {
	return &Team{
		Name:   name,
		Wins:   0,
		Losses: 0,
		Ties:   0,
	}
}

func (t *Team) UpdateRecord(teamScore, opponentScore int) {
	if teamScore > opponentScore {
		t.Wins++
	} else if teamScore < opponentScore {
		t.Losses++
	} else {
		t.Ties++
	}
}

func (t *Team) ToString() string {
	return fmt.Sprintf("%s (%d-%d-%d)", t.Name, t.Wins, t.Losses, t.Ties)
}
