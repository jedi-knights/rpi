package match_test

import (
	"github.com/jedi-knights/rpi/pkg/match"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Match", func() {
	var myMatch *match.Match
	var builder *match.Builder
	var factory *match.Factory

	BeforeEach(func() {
		myMatch = match.NewMatch()
		builder = match.NewBuilder()
		factory = match.NewFactory(builder)
	})

	AfterEach(func() {
		factory = nil
		builder = nil
		myMatch = nil
	})

	Describe("IsHomeTeam", func() {
		It("returns false when the team name is empty", func() {
			// Arrange
			myMatch.Home.Name = "Team U"
			myMatch.Away.Name = "Team T"

			// Act
			answer := myMatch.IsHomeTeam("")

			// Assert
			Expect(answer).To(BeFalse())
		})

		It("returns true when the team name matches the home team name", func() {
			// Arrange
			myMatch.Home.Name = "Team Z"

			// Act
			answer := myMatch.IsHomeTeam("Team Z")

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
			myMatch.Home.Name = "Team Cool"
			myMatch.Away.Name = "Team Dork"

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
				*match.NewBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
				*match.NewBuilder().
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
				*match.NewBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
				*match.NewBuilder().
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
				*match.NewBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
				*match.NewBuilder().
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
				*match.NewBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
				*match.NewBuilder().
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

		It("returns opponent names without duplicates when the specified team is in some of the matches", func() {
			// Arrange
			matches := []match.Match{
				*match.NewBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
				*match.NewBuilder().
					BuildHomeName("Team B").
					BuildAwayName("Team A").
					GetInstance(),
			}

			// Act
			answer, err := match.GetOpponents(&matches, "Team A")

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(answer).To(ConsistOf("Team B"))
			Expect(len(answer)).To(Equal(1))
		})
	})

	Describe("GetMatchesPlayedBy", func() {
		It("returns an error when the specified team is empty", func() {
			// Arrange
			matches := []match.Match{
				*match.NewBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
				*match.NewBuilder().
					BuildHomeName("Team B").
					BuildAwayName("Team C").
					GetInstance(),
			}

			// Act
			pSubslice, err := match.GetMatchesPlayedBy(&matches, "")

			// Assert
			Expect(pSubslice).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the specified team name is empty"))
		})

		It("returns an empty slice when the specified team is not in any of the matches", func() {
			// Arrange
			matches := []match.Match{
				*match.NewBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
				*match.NewBuilder().
					BuildHomeName("Team B").
					BuildAwayName("Team C").
					GetInstance(),
			}

			// Act
			pSubslice, err := match.GetMatchesPlayedBy(&matches, "Team D")

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(*pSubslice).To(BeEmpty())
		})

		It("returns the matches when the specified team is in all matches", func() {
			// Arrange
			matches := []match.Match{
				*match.NewBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
				*match.NewBuilder().
					BuildHomeName("Team B").
					BuildAwayName("Team C").
					GetInstance(),
			}

			// Act
			pSubslice, err := match.GetMatchesPlayedBy(&matches, "Team B")

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(*pSubslice).To(ConsistOf(matches[0], matches[1]))
		})

		It("returns a subset of matches when some of the matches contain the specified team", func() {
			// Arrange
			matches := []match.Match{
				*match.NewBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
				*match.NewBuilder().
					BuildHomeName("Team B").
					BuildAwayName("Team C").
					GetInstance(),
			}

			// Act
			pSubslice, err := match.GetMatchesPlayedBy(&matches, "Team A")

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(*pSubslice).To(ConsistOf(matches[0]))
		})
	})

	Describe("GetMatchesBetween", func() {
		It("returns an error when the first specified team is empty", func() {
			// Arrange
			matches := []match.Match{
				*match.NewBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
			}

			// Act
			pSubslice, err := match.GetMatchesBetween(&matches, "", "Team B")

			// Assert
			Expect(pSubslice).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the first specified team name is empty"))
		})

		It("returns an error when the second specified team is empty", func() {
			// Arrange
			matches := []match.Match{
				*match.NewBuilder().
					BuildHomeName("Team A").
					BuildAwayName("Team B").
					GetInstance(),
			}

			// Act
			pSubslice, err := match.GetMatchesBetween(&matches, "Team A", "")

			// Assert
			Expect(pSubslice).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the second specified team name is empty"))
		})

		It("returns an empty slice when there are no matches between the specified teams", func() {
			// Arrange
			var matches []match.Match

			matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,0"))
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team C,1,Team D,0"))
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team E,1,Team F,0"))
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team B,1,Team G,0"))

			// Act
			pSubslice, err := match.GetMatchesBetween(&matches, "Team A", "Team C")

			// Assert
			Expect(pSubslice).NotTo(BeNil())
			Expect(*pSubslice).To(BeEmpty())
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns a slice with one match when there is one match between the specified teams", func() {
			// Arrange
			var matches []match.Match

			matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,0"))
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team C,1,Team D,0"))
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team E,1,Team F,0"))
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team B,1,Team G,0"))

			// Act
			pSubslice, err := match.GetMatchesBetween(&matches, "Team A", "Team B")

			// Assert
			Expect(pSubslice).NotTo(BeNil())
			Expect(*pSubslice).To(HaveLen(1))
			Expect(err).NotTo(HaveOccurred())
			Expect((*pSubslice)[0]).To(Equal(matches[0]))
		})
	})

	Describe("GetMatchesPlayedByOpponents", func() {
		It("returns all matches played by opponents", func() {
			// Arrange
			var matches []match.Match

			matches = []match.Match{
				*factory.CreateFromString("2020-01-01,UConn,64,Kansas,57"),
				*factory.CreateFromString("2020-01-01,UConn,82,Duke,68"),
				*factory.CreateFromString("2020-01-01,Wisconsin,71,UConn,72"),
				*factory.CreateFromString("2020-01-01,Kansas,69,UConn,62"),
				*factory.CreateFromString("2020-01-01,Duke,81,Wisconsin,70"),
				*factory.CreateFromString("2020-01-01,Wisconsin,52,Kansas,62"),
			}

			// Act
			pSubslice, err := match.GetMatchesPlayedByOpponents(&matches, "UConn")

			// Assert
			Expect(pSubslice).NotTo(BeNil())
			Expect(*pSubslice).To(HaveLen(2))
			Expect(err).NotTo(HaveOccurred())
			Expect((*pSubslice)[0]).To(Equal(matches[4]))
			Expect((*pSubslice)[1]).To(Equal(matches[5]))
		})
	})
})
