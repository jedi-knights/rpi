package match_test

import (
	"github.com/jedi-knights/rpi/pkg/match"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Match", func() {
	var match *match.Match

	BeforeEach(func() {
		match = match.NewMatch()
	})

	AfterEach(func() {
		match = nil
	})

	Describe("IsHomeTeam", func() {
		It("returns false when the team name is empty", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Away.Name = "Team B"

			// Act
			answer := match.IsHomeTeam("")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns true when the team name matches the home team name", func() {
			// Arrange
			match.Home.Name = "Team A"

			// Act
			answer := match.IsHomeTeam("Team A")

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns false when the team name does not match the home team name", func() {
			// Arrange
			match.Home.Name = "Team A"

			// Act
			answer := match.IsHomeTeam("Team B")

			// Assert
			Expect(answer).To(BeFalse())
		})
	})

	Describe("IsAwayTeam", func() {
		It("returns false when the team name is empty", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Away.Name = "Team B"

			// Act
			answer := match.IsAwayTeam("")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns true when the team name matches the away team name", func() {
			// Arrange
			match.Away.Name = "Team B"

			// Act
			answer := match.IsAwayTeam("Team B")

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns false when the team name does not match the away team name", func() {
			// Arrange
			match.Away.Name = "Team B"

			// Act
			answer := match.IsAwayTeam("Team A")

			// Assert
			Expect(answer).To(BeFalse())
		})
	})

	Describe("Contains", func() {
		It("returns false when the team name is empty", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Away.Name = "Team B"

			// Act
			answer := match.Contains("")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns true when the team name matches the home team name", func() {
			// Arrange
			match.Home.Name = "Team A"

			// Act
			answer := match.Contains("Team A")

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns true when the team name matches the away team name", func() {
			// Arrange
			match.Away.Name = "Team B"

			// Act
			answer := match.Contains("Team B")

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns false when the team name does not match the home or away team name", func() {
			match.Home.Name = "Team A"
			match.Away.Name = "Team B"
			Expect(match.Contains("Team C")).To(BeFalse())
		})
	})

	Describe("IsDraw", func() {
		It("returns true when the home score equals the away score", func() {
			// Arrange
			match.Home.Score = 1
			match.Away.Score = 1

			// Act
			answer := match.IsDraw()

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns false when the home score does not equal the away score", func() {
			// Arrange
			match.Home.Score = 1
			match.Away.Score = 2

			// Act
			answer := match.IsDraw()

			// Assert
			Expect(answer).To(BeFalse())
		})
	})

	Describe("IsWinner", func() {
		It("returns false when the specified team is empty", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 1
			match.Away.Name = "Team B"
			match.Away.Score = 2

			// Act
			answer := match.IsWinner("")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns false when the match is a draw", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 1
			match.Away.Name = "Team B"
			match.Away.Score = 1

			// Act
			answer := match.IsWinner("Team A")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns false when the specified team is not in the match", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 1
			match.Away.Name = "Team B"
			match.Away.Score = 2

			// Act
			answer := match.IsWinner("Team C")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns true when the specified home team is the winner", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 2
			match.Away.Name = "Team B"
			match.Away.Score = 1

			// Act
			answer := match.IsWinner("Team A")

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns true when the specified away team is the winner", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 1
			match.Away.Name = "Team B"
			match.Away.Score = 2

			// Act
			answer := match.IsWinner("Team B")

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns false when the specified home team is not the winner", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 1
			match.Away.Name = "Team B"
			match.Away.Score = 2

			// Act
			answer := match.IsWinner("Team A")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns false when the specified away team is not the winner", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 2
			match.Away.Name = "Team B"
			match.Away.Score = 1

			// Act
			answer := match.IsWinner("Team B")

			// Assert
			Expect(answer).To(BeFalse())
		})
	})

	Describe("IsLoser", func() {
		It("returns false when the specified team is empty", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 1
			match.Away.Name = "Team B"
			match.Away.Score = 2

			// Act
			answer := match.IsLoser("")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("with the specified team is not empty and the match is a draw", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 1
			match.Away.Name = "Team B"
			match.Away.Score = 1

			// Act
			answer := match.IsLoser("Team A")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns true when the away team is the loser", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 2
			match.Away.Name = "Team B"
			match.Away.Score = 1

			// Act
			answer := match.IsLoser("Team B")

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns false when the match is a draw", func() {
			// Arrange
			match.Home.Score = 1
			match.Away.Score = 1

			// Act
			answer := match.IsLoser("Team A")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns false when the specified team is not in the match", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 1
			match.Away.Name = "Team B"
			match.Away.Score = 2

			// Act
			answer := match.IsLoser("Team C")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns true when the specified team is the loser", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 1
			match.Away.Name = "Team B"
			match.Away.Score = 2

			// Act
			answer := match.IsLoser("Team A")

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns false when the specified team is not the loser", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 2
			match.Away.Name = "Team B"
			match.Away.Score = 1

			// Act
			answer := match.IsLoser("Team A")

			// Assert
			Expect(answer).To(BeFalse())
		})
	})

	Describe("WinValue", func() {
		It("returns 0.0 when the specified team is the loser", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 2
			match.Away.Name = "Team B"
			match.Away.Score = 1

			// Act
			answer := match.WinValue("Team B")

			// Assert
			Expect(answer).To(Equal(0.0))
		})

		It("returns 0.0 when the specified team is empty", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 1
			match.Away.Name = "Team B"
			match.Away.Score = 2

			// Act
			answer := match.WinValue("")

			// Assert
			Expect(answer).To(Equal(0.0))
		})

		It("returns 0.5 when the match is a draw", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 1
			match.Away.Name = "Team B"
			match.Away.Score = 1

			// Act
			answer := match.WinValue("Team A")

			// Assert
			Expect(answer).To(Equal(0.5))
		})

		It("returns 0.0 when the specified team is not in the match", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 1
			match.Away.Name = "Team B"
			match.Away.Score = 2

			// Act
			answer := match.WinValue("Team C")

			// Assert
			Expect(answer).To(Equal(0.0))
		})

		It("returns 1.0 when the specified team is the winner", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Home.Score = 2
			match.Away.Name = "Team B"
			match.Away.Score = 1

			// Act
			answer := match.WinValue("Team A")

			// Assert
			Expect(answer).To(Equal(1.0))
		})
	})

	Describe("GetOpponent", func() {
		It("returns an error when the specified team is empty", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Away.Name = "Team B"

			// Act
			_, err := match.GetOpponent("")

			// Assert
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the specified team name is empty"))
		})

		It("returns an error when the specified team is not in the match", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Away.Name = "Team B"

			// Act
			_, err := match.GetOpponent("Team C")

			// Assert
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the match doesn't contain the team <Team C>"))
		})

		It("returns the opponent name when the specified team is the home team", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Away.Name = "Team B"

			// Act
			answer, err := match.GetOpponent("Team A")

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(answer).To(Equal("Team B"))
		})

		It("returns the opponent name when the specified team is the away team", func() {
			// Arrange
			match.Home.Name = "Team A"
			match.Away.Name = "Team B"

			// Act
			answer, err := match.GetOpponent("Team B")

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(answer).To(Equal("Team A"))
		})
	})

	Describe("GetOpponents", func() {
		It("returns an error when the specified team is empty", func() {
			// Arrange
			matches := []*match.Match{
				match.NewMatchBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
				match.NewMatchBuilder().
					BuildHomeName("Team B").
					BuildAwayName("Team C").
					GetInstance(),
			}

			// Act
			teams, err := match.GetOpponents(matches, "")

			// Assert
			Expect(teams).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the specified team name is empty"))
		})

		It("returns an empty slice when the specified team is not in any of the matches", func() {
			// Arrange
			matches := []*match.Match{
				match.NewMatchBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
				match.NewMatchBuilder().
					BuildHomeName("Team B").
					BuildAwayName("Team C").
					GetInstance(),
			}

			// Act
			answer, err := match.GetOpponents(matches, "Team D")

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(answer).To(BeEmpty())
		})

		It("returns the opponent names when the specified team is in all matches", func() {
			// Arrange
			matches := []*match.Match{
				match.NewMatchBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
				match.NewMatchBuilder().
					BuildHomeName("Team B").
					BuildAwayName("Team C").
					GetInstance(),
			}

			// Act
			answer, err := match.GetOpponents(matches, "Team B")

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(answer).To(ConsistOf("Team A", "Team C"))
		})

		It("returns the opponent names when the specified team is in some of the matches", func() {
			// Arrange
			matches := []*match.Match{
				match.NewMatchBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
				match.NewMatchBuilder().
					BuildHomeName("Team B").
					BuildAwayName("Team C").
					GetInstance(),
			}

			// Act
			answer, err := match.GetOpponents(matches, "Team A")

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(answer).To(ConsistOf("Team B"))
		})
	})
})
