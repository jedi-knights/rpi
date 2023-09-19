package accumulators_test

import (
	"github.com/jedi-knights/rpi/pkg"
	"github.com/jedi-knights/rpi/pkg/accumulators"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Losses", func() {
	var matches []pkg.Match
	var builder *pkg.MatchBuilder
	var factory *pkg.MatchFactory
	var lossesAccumulator *accumulators.Losses

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
				lossesAccumulator = accumulators.NewLosses("")
			})

			AfterEach(func() {
				lossesAccumulator = nil
			})

			It("should return 0 when the team name is empty", func() {
				// Arrange
				teamName := ""

				// Act
				losses, _ := lossesAccumulator.Calculate(teamName, &matches)

				// Assert
				Expect(losses).To(Equal(0))
			})

			It("should return 0 when the team name is not found", func() {
				// Arrange
				teamName := "Ashland Blazer"

				// Act
				losses, _ := lossesAccumulator.Calculate(teamName, &matches)

				// Assert
				Expect(losses).To(Equal(0))
			})

			It("should return 1 when the team has a single loss", func() {
				// Arrange
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,0"))

				// Act
				losses, _ := lossesAccumulator.Calculate("Team B", &matches)

				// Assert
				Expect(losses).To(Equal(1))
			})

			It("should return 1 when the team has a single win and a single loss", func() {
				// Arrange
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,4"))
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team C,1,Team B,0"))

				// Act
				losses, _ := lossesAccumulator.Calculate("Team B", &matches)

				// Assert
				Expect(losses).To(Equal(1))
			})

			It("should return 2 when the team has two losses", func() {
				// Arrange
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,4"))
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team C,1,Team B,0"))
				matches = append(matches, *factory.CreateFromString("2023-09-19,Team D,1,Team B,0"))

				// Act
				losses, _ := lossesAccumulator.Calculate("Team B", &matches)

				// Assert
				Expect(losses).To(Equal(2))
			})
		})
	})

	Context("non-empty skip team name", func() {
		BeforeEach(func() {
			lossesAccumulator = accumulators.NewLosses("Team A")
		})

		AfterEach(func() {
			lossesAccumulator = nil
		})

		It("should return 0 when the team name is empty", func() {
			// Act
			losses, _ := lossesAccumulator.Calculate("", &matches)

			// Assert
			Expect(losses).To(Equal(0))
		})

		It("should return 0 when the team name is not found", func() {
			// Act
			losses, _ := lossesAccumulator.Calculate("Ashland Blazer", &matches)

			// Assert
			Expect(losses).To(Equal(0))
		})

		It("should skip over a match with the skipped team", func() {
			// Arrange
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,4"))

			// Act
			losses, _ := lossesAccumulator.Calculate("Team B", &matches)

			// Assert
			Expect(losses).To(Equal(0))
		})

		It("should return 1 when the team has a single loss and skips over a loss with the skipped team", func() {
			// Arrange
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team A,1,Team B,0"))
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team C,1,Team B,0"))
			matches = append(matches, *factory.CreateFromString("2023-09-19,Team D,1,Team E,0"))

			// Act
			losses, _ := lossesAccumulator.Calculate("Team B", &matches)

			// Assert
			Expect(losses).To(Equal(1))
		})
	})
})
