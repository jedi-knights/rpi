package calculators_test

import (
	"github.com/jedi-knights/rpi/pkg/calculators"
	"github.com/jedi-knights/rpi/pkg/match"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Opponents Winning Percentage", func() {
	var builder *match.Builder
	var factory *match.Factory
	var calculator *calculators.OWPCalculator

	BeforeEach(func() {
		builder = match.NewBuilder()
		factory = match.NewFactory(builder)
	})

	AfterEach(func() {
		factory = nil
		builder = nil
	})

	Context("Wikipedia Example", func() {
		var matches []match.Match

		BeforeEach(func() {
			calculator = calculators.NewOWPCalculator()

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

		It("should return the correct opponents winning percentage for UConn", func() {
			// Act
			owp, err := calculator.Calculate("UConn", &matches)

			// Assert
			Expect(err).NotTo(HaveOccurred())
			Expect(owp).Should(BeNumerically("~", 0.7500, 0.0001))
		})

		It("should return the correct opponents winning percentage for Kansas", func() {
			// Act
			owp, err := calculator.Calculate("Kansas", &matches)

			// Assert
			Expect(err).NotTo(HaveOccurred())
			Expect(owp).Should(BeNumerically("~", 0.6667, 0.0001))
		})

		It("should return the correct opponents winning percentage for Duke", func() {
			// Act
			owp, err := calculator.Calculate("Duke", &matches)

			// Assert
			Expect(err).NotTo(HaveOccurred())
			Expect(owp).Should(BeNumerically("~", 0.3333, 0.0001))
		})

		It("should return the correct opponents winning percentage for Wisconsin", func() {
			// Act
			owp, err := calculator.Calculate("Wisconsin", &matches)

			// Assert
			Expect(err).NotTo(HaveOccurred())
			Expect(owp).Should(BeNumerically("~", 0.3889, 0.0001))
		})
	})
})
