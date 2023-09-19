package accumulators_test

import (
	"github.com/jedi-knights/rpi/pkg/accumulators"
	"github.com/jedi-knights/rpi/pkg/match"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ties", func() {
	var matches []match.Match
	var builder *match.MatchBuilder
	var factory *match.MatchFactory
	var drawsAccumulator *accumulators.Ties

	BeforeEach(func() {
		builder = match.NewMatchBuilder()
		factory = match.NewMatchFactory(builder)
		matches = []match.Match{}
	})

	AfterEach(func() {
		matches = nil
		factory = nil
		builder = nil
	})

	Context("empty skip team name", func() {
		Describe("Calculate", func() {
			BeforeEach(func() {
				drawsAccumulator = accumulators.NewTies("Team A")
			})

			AfterEach(func() {
				drawsAccumulator = nil
			})

			It("should return 0 when the team name is empty", func() {
				// Arrange
				teamName := ""

				// Act
				draws, _ := drawsAccumulator.Calculate(teamName, &matches)

				// Assert
				Expect(draws).To(Equal(0))
			})

			It("should return 0 when the team name is not found", func() {
				// Arrange
				teamName := "Ashland Blazer"

				// Act
				draws, _ := drawsAccumulator.Calculate(teamName, &matches)

				// Assert
				Expect(draws).To(Equal(0))
			})
		})
	})

	Context("non-empty skip team name", func() {
		Describe("Calculate", func() {
			BeforeEach(func() {
				drawsAccumulator = accumulators.NewTies("Team A")
			})

			AfterEach(func() {
				drawsAccumulator = nil
			})

			It("should return 0 when the team name is empty", func() {
				// Arrange
				teamName := ""

				// Act
				draws, _ := drawsAccumulator.Calculate(teamName, &matches)

				// Assert
				Expect(draws).To(Equal(0))
			})

			It("should return 0 when the team name is not found", func() {
				// Arrange
				teamName := "Ashland Blazer"

				// Act
				draws, _ := drawsAccumulator.Calculate(teamName, &matches)

				// Assert
				Expect(draws).To(Equal(0))
			})

			It("should return 0 when the only match includes the skipped team", func() {
				// Arrange
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,1"))

				// Act
				draws, _ := drawsAccumulator.Calculate("Team B", &matches)

				// Assert
				Expect(draws).To(Equal(0))
			})

			It("shouold return 2 when the team has two draws skipping the skip team name", func() {
				// Arrange
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,1"))
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team C,1,Team B,1"))
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team B,1,Team D,1"))
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team E,1,Team F,1"))

				// Act
				draws, _ := drawsAccumulator.Calculate("Team B", &matches)

				// Assert
				Expect(draws).To(Equal(2))
			})
		})
	})
})
