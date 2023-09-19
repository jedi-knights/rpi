package pkg_test

import (
	"github.com/jedi-knights/rpi/pkg"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("MatchBuilder", func() {
	var builder *pkg.MatchBuilder

	BeforeEach(func() {
		builder = pkg.NewMatchBuilder()
	})

	AfterEach(func() {
		builder = nil
	})

	It("should be able to build a match", func() {
		// Arrange
		now := time.Now()
		homeName := "Ashland Blazer"
		homeScore := 1
		awayName := "Raceland"
		awayScore := 0

		// Act
		match := builder.
			BuildDate(now).
			BuildHomeName(homeName).
			BuildHomeScore(homeScore).
			BuildAwayName(awayName).
			BuildAwayScore(awayScore).
			GetInstance()

		// Assert
		Expect(match.Date).To(Equal(now))
		Expect(match.Home.Name).To(Equal(homeName))
		Expect(match.Home.Score).To(Equal(homeScore))
		Expect(match.Away.Name).To(Equal(awayName))
		Expect(match.Away.Score).To(Equal(awayScore))
	})
})
