package calculators_test

import (
	"github.com/jedi-knights/rpi/pkg/calculators"
	"github.com/jedi-knights/rpi/pkg/match"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Winning Percentage Calculator", func() {
	var builder *match.Builder
	var factory *match.Factory
	var calculator *calculators.WPCalculator

	BeforeEach(func() {
		builder = match.NewBuilder()
		factory = match.NewFactory(builder)
	})

	AfterEach(func() {
		factory = nil
		builder = nil
	})

	Context("when the skipTeamName is empty", func() {
		BeforeEach(func() {
			calculator = calculators.NewWPCalculator("")
		})

		AfterEach(func() {
			calculator = nil
		})

		It("returns 0.5 for a team with an 8-8-4 record", func() {
			// Arrange
			matches := []match.Match{
				*factory.CreateFromString("2020-01-01,Team A,1,Team B,0"),
				*factory.CreateFromString("2020-01-02,Team A,1,Team C,0"),
				*factory.CreateFromString("2020-01-02,Team A,1,Team D,0"),
				*factory.CreateFromString("2020-01-02,Team A,1,Team E,0"),
				*factory.CreateFromString("2020-01-02,Team A,1,Team F,0"),
				*factory.CreateFromString("2020-01-02,Team A,1,Team G,0"),
				*factory.CreateFromString("2020-01-02,Team A,1,Team H,0"),
				*factory.CreateFromString("2020-01-02,Team A,1,Team I,0"),
				*factory.CreateFromString("2020-01-02,Team A,0,Team J,1"),
				*factory.CreateFromString("2020-01-02,Team A,0,Team K,1"),
				*factory.CreateFromString("2020-01-02,Team A,0,Team L,1"),
				*factory.CreateFromString("2020-01-02,Team A,0,Team M,1"),
				*factory.CreateFromString("2020-01-02,Team A,0,Team N,1"),
				*factory.CreateFromString("2020-01-02,Team A,0,Team O,1"),
				*factory.CreateFromString("2020-01-02,Team A,0,Team P,1"),
				*factory.CreateFromString("2020-01-02,Team A,0,Team Q,1"),
				*factory.CreateFromString("2020-01-02,Team A,0,Team R,0"),
				*factory.CreateFromString("2020-01-02,Team A,0,Team S,0"),
				*factory.CreateFromString("2020-01-02,Team A,0,Team T,0"),
				*factory.CreateFromString("2020-01-02,Team A,0,Team U,0"),
			}

			// Act
			wp, err := calculator.Calculate("Team A", &matches)

			// Assert
			Expect(err).NotTo(HaveOccurred())
			Expect(wp).To(Equal(0.5))
		})
	})

	// https://en.wikipedia.org/wiki/Rating_percentage_index
	Context("Wikipedia Example", func() {
		var matches []match.Match

		BeforeEach(func() {
			calculator = calculators.NewWPCalculator("")

			matches = []match.Match{
				*factory.CreateFromString("2020-01-01,UConn,64,Kansas,57"),
				*factory.CreateFromString("2020-01-01,UConn,82,Duke,68"),
				*factory.CreateFromString("2020-01-01,Wisconsin,71,UConn,72"),
				*factory.CreateFromString("2020-01-01,Kansas,69,UConn,62"),
				*factory.CreateFromString("2020-01-01,Duke,81,Wisconsin,70"),
				*factory.CreateFromString("2020-01-01,Wisconsin,52,Kansas,62"),
			}
		})

		AfterEach(func() {
			calculator = nil
		})

		It("should return the correct winning percentage for UConn", func() {
			// Act
			wp, err := calculator.Calculate("UConn", &matches)

			// Assert
			Expect(err).NotTo(HaveOccurred())
			Expect(wp).Should(BeNumerically("~", 0.7500, 0.0001))
		})

		It("should return the correct winning percentage for Kansas", func() {
			// Act
			wp, err := calculator.Calculate("Kansas", &matches)

			// Assert
			Expect(err).NotTo(HaveOccurred())
			Expect(wp).Should(BeNumerically("~", 0.6667, 0.0001))
		})

		It("should return the correct winning percentage for Duke", func() {
			// Act
			wp, err := calculator.Calculate("Duke", &matches)

			// Assert
			Expect(err).NotTo(HaveOccurred())
			Expect(wp).Should(BeNumerically("~", 0.5000, 0.0001))
		})

		It("should return the correct winning percentage for Wisconsin", func() {
			// Act
			wp, err := calculator.Calculate("Wisconsin", &matches)

			// Assert
			Expect(err).NotTo(HaveOccurred())
			Expect(wp).Should(BeNumerically("~", 0.0000, 0.0001))
		})
	})
})
