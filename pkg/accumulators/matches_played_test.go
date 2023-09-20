package accumulators_test

import (
	"github.com/jedi-knights/rpi/pkg/accumulators"
	"github.com/jedi-knights/rpi/pkg/match"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MatchesPlayed", func() {
	var matches []match.Match
	var builder *match.Builder
	var factory *match.Factory
	var matchesPlayedAccumulator *accumulators.MatchesPlayed

	BeforeEach(func() {
		builder = match.NewBuilder()
		factory = match.NewFactory(builder)
		matches = []match.Match{}
	})

	AfterEach(func() {
		matches = nil
		factory = nil
		builder = nil
	})

	Context("empty skip team name", func() {
		BeforeEach(func() {
			matchesPlayedAccumulator = accumulators.NewMatchesPlayed("")
		})

		AfterEach(func() {
			matchesPlayedAccumulator = nil
		})

		It("should return 0 when the team name is empty", func() {
			// Arrange
			teamName := ""

			// Act
			matchesPlayed, _ := matchesPlayedAccumulator.Calculate(teamName, &matches)

			// Assert
			Expect(matchesPlayed).To(Equal(0))
		})

		It("should return 0 when the team name is not found", func() {
			// Act
			matchesPlayed, _ := matchesPlayedAccumulator.Calculate("something", &matches)

			// Assert
			Expect(matchesPlayed).To(Equal(0))
		})

		It("should return 1 when the team has a single match", func() {
			// Arrange
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,0"))

			// Act
			matchesPlayed, _ := matchesPlayedAccumulator.Calculate("Team A", &matches)

			// Assert
			Expect(matchesPlayed).To(Equal(1))
		})

		It("should return 2 when the team has two matches", func() {
			// Arrange
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,0"))
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team C,0"))

			// Act
			matchesPlayed, _ := matchesPlayedAccumulator.Calculate("Team A", &matches)

			// Assert
			Expect(matchesPlayed).To(Equal(2))
		})

		It("should return 3 when the team has three matches", func() {
			// Arrange
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,0"))
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team C,0"))
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team D,1,Team A,0"))

			// Act
			matchesPlayed, _ := matchesPlayedAccumulator.Calculate("Team A", &matches)

			// Assert
			Expect(matchesPlayed).To(Equal(3))
		})
	})

	Context("non-empty skip team name", func() {
		BeforeEach(func() {
			matchesPlayedAccumulator = accumulators.NewMatchesPlayed("Team A")
		})

		AfterEach(func() {
			matchesPlayedAccumulator = nil
		})

		It("should return 0 when the team name is empty", func() {
			// Arrange
			teamName := ""

			// Act
			matchesPlayed, _ := matchesPlayedAccumulator.Calculate(teamName, &matches)

			// Assert
			Expect(matchesPlayed).To(Equal(0))
		})

		It("should return 0 when the team name is not found", func() {
			// Arrange
			teamName := "Ashland Blazer"

			// Act
			matchesPlayed, _ := matchesPlayedAccumulator.Calculate(teamName, &matches)

			// Assert
			Expect(matchesPlayed).To(Equal(0))
		})

		It("should return 0 when the skipped team has two matches", func() {
			// Arrange
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,0"))
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team C,0"))

			// Act
			matchesPlayed, _ := matchesPlayedAccumulator.Calculate("Team B", &matches)

			// Assert
			Expect(matchesPlayed).To(Equal(0))
		})
	})
})
