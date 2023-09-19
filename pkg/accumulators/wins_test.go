package accumulators_test

import (
	"github.com/jedi-knights/rpi/pkg"
	"github.com/jedi-knights/rpi/pkg/accumulators"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Wins", func() {
	var matches []pkg.Match
	var builder *pkg.MatchBuilder
	var factory *pkg.MatchFactory
	var winsAccumulator *accumulators.Wins

	BeforeEach(func() {
		builder = pkg.NewMatchBuilder()
		factory = pkg.NewMatchFactory(builder)
		matches = []pkg.Match{}
	})

	AfterEach(func() {
		matches = nil
		factory = nil
		builder = nil
	})

	Context("empty skip team name", func() {
		Describe("Calculate", func() {
			BeforeEach(func() {
				winsAccumulator = accumulators.NewWins("")
			})

			AfterEach(func() {
				winsAccumulator = nil
			})

			It("should return 0 when the team name is empty", func() {
				// Arrange
				teamName := ""

				// Act
				wins, _ := winsAccumulator.Calculate(teamName, &matches)

				// Assert
				Expect(wins).To(Equal(0))
			})

			It("should return 0 when the team name is not found", func() {
				// Arrange
				teamName := "Ashland Blazer"

				// Act
				wins, _ := winsAccumulator.Calculate(teamName, &matches)

				// Assert
				Expect(wins).To(Equal(0))
			})

			It("should return 1 when the team has a single win", func() {
				// Arrange
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,0"))

				// Act
				wins, _ := winsAccumulator.Calculate("Team A", &matches)

				// Assert
				Expect(wins).To(Equal(1))
			})

			It("should return 1 when the team has a single win and a single loss", func() {
				// Arrange
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,0"))
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team C,1,Team A,0"))

				// Act
				wins, _ := winsAccumulator.Calculate("Team A", &matches)

				// Assert
				Expect(wins).To(Equal(1))
			})

			It("should return 2 when the team has two wins and a single loss", func() {
				// Arrange
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,0"))
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team C,1,Team A,0"))
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team D,0"))

				// Act
				wins, _ := winsAccumulator.Calculate("Team A", &matches)

				// Assert
				Expect(wins).To(Equal(2))
			})
		})
	})

	Context("non-empty skip team name", func() {
		Describe("Calculate", func() {
			const skippedTeamName = "Team A"

			BeforeEach(func() {
				winsAccumulator = accumulators.NewWins(skippedTeamName)
			})

			AfterEach(func() {
				winsAccumulator = nil
			})

			It("should return 0 when the team name is empty", func() {
				// Arrange
				teamName := ""

				// Act
				wins, _ := winsAccumulator.Calculate(teamName, &matches)

				// Assert
				Expect(wins).To(Equal(0))
			})

			It("should return 0 when the team name is not found", func() {
				// Arrange
				teamName := "Ashland Blazer"

				// Act
				wins, _ := winsAccumulator.Calculate(teamName, &matches)

				// Assert
				Expect(wins).To(Equal(0))
			})

			It("should skip over a match with the skipped team", func() {
				// Arrange
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,4"))

				// Act
				wins, _ := winsAccumulator.Calculate("Team B", &matches)

				// Assert
				Expect(wins).To(Equal(0))
			})

			It("should return 1 when the team has a single win and skips over a win with the skipped team", func() {
				// Arrange
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team C,2"))
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team C,1,Team D,0"))

				// Act
				wins, _ := winsAccumulator.Calculate("Team C", &matches)

				// Assert
				Expect(wins).To(Equal(1))
			})

			It("should return 2 when the team has two wins and skips over a win with the skipped team", func() {
				// Arrange
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,4"))
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team C,1,Team B,3"))
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team B,1,Team D,0"))
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team E,1,Team F,0"))

				// Act
				wins, _ := winsAccumulator.Calculate("Team B", &matches)

				// Assert
				Expect(wins).To(Equal(2))
			})
		})
	})
})
