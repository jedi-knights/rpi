package match_test

import (
	"github.com/jedi-knights/rpi/pkg/match"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Factory", func() {
	var builder *match.Builder
	var factory *match.Factory

	BeforeEach(func() {
		builder = match.NewBuilder()
		factory = match.NewFactory(builder)
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
