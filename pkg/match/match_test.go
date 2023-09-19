package match_test

import (
	"github.com/jedi-knights/rpi/pkg/match"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Match", func() {
	var myMatch *match.Match

	BeforeEach(func() {
		myMatch = match.NewMatch()
	})

	AfterEach(func() {
		myMatch = nil
	})

	Describe("IsHomeTeam", func() {
		It("returns false when the team name is empty", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Away.Name = "Team B"

			// Act
			answer := myMatch.IsHomeTeam("")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns true when the team name matches the home team name", func() {
			// Arrange
			myMatch.Home.Name = "Team A"

			// Act
			answer := myMatch.IsHomeTeam("Team A")

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns false when the team name does not myMatch the home team name", func() {
			// Arrange
			myMatch.Home.Name = "Team A"

			// Act
			answer := myMatch.IsHomeTeam("Team B")

			// Assert
			Expect(answer).To(BeFalse())
		})
	})

	Describe("IsAwayTeam", func() {
		It("returns false when the team name is empty", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Away.Name = "Team B"

			// Act
			answer := myMatch.IsAwayTeam("")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns true when the team name matches the away team name", func() {
			// Arrange
			myMatch.Away.Name = "Team B"

			// Act
			answer := myMatch.IsAwayTeam("Team B")

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns false when the team name does not myMatch the away team name", func() {
			// Arrange
			myMatch.Away.Name = "Team B"

			// Act
			answer := myMatch.IsAwayTeam("Team A")

			// Assert
			Expect(answer).To(BeFalse())
		})
	})

	Describe("Contains", func() {
		It("returns false when the team name is empty", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Away.Name = "Team B"

			// Act
			answer := myMatch.Contains("")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns true when the team name matches the home team name", func() {
			// Arrange
			myMatch.Home.Name = "Team A"

			// Act
			answer := myMatch.Contains("Team A")

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns true when the team name matches the away team name", func() {
			// Arrange
			myMatch.Away.Name = "Team B"

			// Act
			answer := myMatch.Contains("Team B")

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns false when the team name does not myMatch the home or away team name", func() {
			myMatch.Home.Name = "Team A"
			myMatch.Away.Name = "Team B"
			Expect(myMatch.Contains("Team C")).To(BeFalse())
		})
	})

	Describe("IsDraw", func() {
		It("returns true when the home score equals the away score", func() {
			// Arrange
			myMatch.Home.Score = 1
			myMatch.Away.Score = 1

			// Act
			answer := myMatch.IsDraw()

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns false when the home score does not equal the away score", func() {
			// Arrange
			myMatch.Home.Score = 1
			myMatch.Away.Score = 2

			// Act
			answer := myMatch.IsDraw()

			// Assert
			Expect(answer).To(BeFalse())
		})
	})

	Describe("IsWinner", func() {
		It("returns false when the specified team is empty", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 1
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 2

			// Act
			answer := myMatch.IsWinner("")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns false when the myMatch is a draw", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 1
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 1

			// Act
			answer := myMatch.IsWinner("Team A")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns false when the specified team is not in the myMatch", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 1
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 2

			// Act
			answer := myMatch.IsWinner("Team C")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns true when the specified home team is the winner", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 2
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 1

			// Act
			answer := myMatch.IsWinner("Team A")

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns true when the specified away team is the winner", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 1
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 2

			// Act
			answer := myMatch.IsWinner("Team B")

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns false when the specified home team is not the winner", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 1
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 2

			// Act
			answer := myMatch.IsWinner("Team A")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns false when the specified away team is not the winner", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 2
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 1

			// Act
			answer := myMatch.IsWinner("Team B")

			// Assert
			Expect(answer).To(BeFalse())
		})
	})

	Describe("IsLoser", func() {
		It("returns false when the specified team is empty", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 1
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 2

			// Act
			answer := myMatch.IsLoser("")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("with the specified team is not empty and the myMatch is a draw", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 1
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 1

			// Act
			answer := myMatch.IsLoser("Team A")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns true when the away team is the loser", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 2
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 1

			// Act
			answer := myMatch.IsLoser("Team B")

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns false when the myMatch is a draw", func() {
			// Arrange
			myMatch.Home.Score = 1
			myMatch.Away.Score = 1

			// Act
			answer := myMatch.IsLoser("Team A")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns false when the specified team is not in the myMatch", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 1
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 2

			// Act
			answer := myMatch.IsLoser("Team C")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns true when the specified team is the loser", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 1
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 2

			// Act
			answer := myMatch.IsLoser("Team A")

			// Assert
			Expect(answer).To(BeTrue())
		})

		It("returns false when the specified team is not the loser", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 2
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 1

			// Act
			answer := myMatch.IsLoser("Team A")

			// Assert
			Expect(answer).To(BeFalse())
		})
	})

	Describe("WinValue", func() {
		It("returns 0.0 when the specified team is the loser", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 2
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 1

			// Act
			answer := myMatch.WinValue("Team B")

			// Assert
			Expect(answer).To(Equal(0.0))
		})

		It("returns 0.0 when the specified team is empty", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 1
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 2

			// Act
			answer := myMatch.WinValue("")

			// Assert
			Expect(answer).To(Equal(0.0))
		})

		It("returns 0.5 when the myMatch is a draw", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 1
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 1

			// Act
			answer := myMatch.WinValue("Team A")

			// Assert
			Expect(answer).To(Equal(0.5))
		})

		It("returns 0.0 when the specified team is not in the myMatch", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 1
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 2

			// Act
			answer := myMatch.WinValue("Team C")

			// Assert
			Expect(answer).To(Equal(0.0))
		})

		It("returns 1.0 when the specified team is the winner", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Home.Score = 2
			myMatch.Away.Name = "Team B"
			myMatch.Away.Score = 1

			// Act
			answer := myMatch.WinValue("Team A")

			// Assert
			Expect(answer).To(Equal(1.0))
		})
	})

	Describe("GetOpponent", func() {
		It("returns an error when the specified team is empty", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Away.Name = "Team B"

			// Act
			_, err := myMatch.GetOpponent("")

			// Assert
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the specified team name is empty"))
		})

		It("returns an error when the specified team is not in the myMatch", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Away.Name = "Team B"

			// Act
			_, err := myMatch.GetOpponent("Team C")

			// Assert
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the match doesn't contain the team <Team C>"))
		})

		It("returns the opponent name when the specified team is the home team", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Away.Name = "Team B"

			// Act
			answer, err := myMatch.GetOpponent("Team A")

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(answer).To(Equal("Team B"))
		})

		It("returns the opponent name when the specified team is the away team", func() {
			// Arrange
			myMatch.Home.Name = "Team A"
			myMatch.Away.Name = "Team B"

			// Act
			answer, err := myMatch.GetOpponent("Team B")

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(answer).To(Equal("Team A"))
		})
	})

	Describe("GetOpponents", func() {
		It("returns an error when the specified team is empty", func() {
			// Arrange
			matches := []match.Match{
				*match.NewMatchBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
				*match.NewMatchBuilder().
					BuildHomeName("Team B").
					BuildAwayName("Team C").
					GetInstance(),
			}

			// Act
			teams, err := match.GetOpponents(&matches, "")

			// Assert
			Expect(teams).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the specified team name is empty"))
		})

		It("returns an empty slice when the specified team is not in any of the matches", func() {
			// Arrange
			matches := []match.Match{
				*match.NewMatchBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
				*match.NewMatchBuilder().
					BuildHomeName("Team B").
					BuildAwayName("Team C").
					GetInstance(),
			}

			// Act
			answer, err := match.GetOpponents(&matches, "Team D")

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(answer).To(BeEmpty())
		})

		It("returns the opponent names when the specified team is in all matches", func() {
			// Arrange
			matches := []match.Match{
				*match.NewMatchBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
				*match.NewMatchBuilder().
					BuildHomeName("Team B").
					BuildAwayName("Team C").
					GetInstance(),
			}

			// Act
			answer, err := match.GetOpponents(&matches, "Team B")

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(answer).To(ConsistOf("Team A", "Team C"))
		})

		It("returns the opponent names when the specified team is in some of the matches", func() {
			// Arrange
			matches := []match.Match{
				*match.NewMatchBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
				*match.NewMatchBuilder().
					BuildHomeName("Team B").
					BuildAwayName("Team C").
					GetInstance(),
			}

			// Act
			answer, err := match.GetOpponents(&matches, "Team A")

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(answer).To(ConsistOf("Team B"))
		})
	})
})
