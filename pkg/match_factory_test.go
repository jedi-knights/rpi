package pkg_test

import (
	"github.com/jedi-knights/rpi/pkg"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MatchFactory", func() {
	var builder *pkg.MatchBuilder
	var factory *pkg.MatchFactory

	BeforeEach(func() {
		builder = pkg.NewMatchBuilder()
		factory = pkg.NewMatchFactory(builder)
	})

	AfterEach(func() {
		factory = nil
		builder = nil
	})

	Describe("CreateWithRandomDate", func() {
		It("should return a match with a random date", func() {
			// Act
			match := factory.CreateWithRandomDate("Team A", 1, "Team B", 0)

			// Assert
			Expect(match.Date).ToNot(BeNil())
			Expect(match.Home.Name).To(Equal("Team A"))
			Expect(match.Home.Score).To(Equal(1))
			Expect(match.Away.Name).To(Equal("Team B"))
			Expect(match.Away.Score).To(Equal(0))
		})
	})
})
