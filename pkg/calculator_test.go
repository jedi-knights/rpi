package pkg_test

import (
	"github.com/jedi-knights/rpi/pkg"
	"github.com/jedi-knights/rpi/pkg/match"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RPI", Ordered, func() {
	var builder *match.MatchBuilder
	var factory *match.MatchFactory
	var matchSlice1 []string

	BeforeAll(func() {
		builder = match.NewMatchBuilder()
		factory = match.NewMatchFactory(builder)

		matchSlice1 = []string{
			"2023-09-19,Team A,1,Team B,0",
			"2023-09-19,Team A,1,Team C,0",
			"2023-09-19,Team A,1,Team D,0",
			"2023-09-19,Team A,1,Team E,0",
			"2023-09-19,Team A,1,Team F,0",
			"2023-09-19,Team A,1,Team G,0",
			"2023-09-19,Team A,1,Team H,0",
			"2023-09-19,Team A,1,Team I,0",
			"2023-09-19,Team A,0,Team J,1",
			"2023-09-19,Team A,0,Team K,1",
			"2023-09-19,Team A,0,Team L,1",
			"2023-09-19,Team A,0,Team M,1",
			"2023-09-19,Team A,0,Team N,1",
			"2023-09-19,Team A,0,Team O,1",
			"2023-09-19,Team A,0,Team P,1",
			"2023-09-19,Team A,0,Team Q,1",
			"2023-09-19,Team A,0,Team R,0",
			"2023-09-19,Team A,0,Team S,0",
			"2023-09-19,Team A,0,Team T,0",
			"2023-09-19,Team A,0,Team U,0",
		}

	})

	AfterAll(func() {
		matchSlice1 = nil
		factory = nil
		builder = nil
	})

	Describe("Accumulate", func() {
		var matches []match.Match

		BeforeAll(func() {
			matches = []match.Match{}

			for _, matchString := range matchSlice1 {
				matches = append(matches, *factory.CreateFromString(matchString))
			}
		})

		AfterAll(func() {
			matches = nil
		})

		Context("with no skipped team", func() {
			Describe("Wins", func() {
				It("should return the number of wins for a team", func() {
					// Act
					wins, err := pkg.Accumulate("Wins", "Team A", "", &matches)

					// Assert
					Expect(wins).To(Equal(8))
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Describe("Losses", func() {
				It("should return the number of losses for a team", func() {
					// Act
					losses, err := pkg.Accumulate("Losses", "Team A", "", &matches)

					// Assert
					Expect(losses).To(Equal(8))
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Describe("Ties", func() {
				It("should return the number of ties for a team", func() {
					// Act
					ties, err := pkg.Accumulate("Ties", "Team A", "", &matches)

					// Assert
					Expect(ties).To(Equal(4))
					Expect(err).NotTo(HaveOccurred())
				})
			})
		})
	})

	Context("Element 1", func() {
		Describe("WP", func() {
			It("should return 0.5 for a team with a record of 8-8-4", func() {
				// Arrange
				matches := []match.Match{}

				for _, matchString := range matchSlice1 {
					matches = append(matches, *factory.CreateFromString(matchString))
				}

				// Act
				wp, err := pkg.Calculate("WP", "Team A", &matches)

				// Assert
				Expect(wp).Should(BeNumerically("~", 0.5, 0.001))
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})

//var _ = Describe("RPI", Ordered, func() {
//	var matches []pkg.Match
//
//	BeforeAll(func() {
//		matches = []pkg.Match{
//			{HomeTeamName: "Ashland Blazer", HomeTeamScore: 1, AwayTeamName: "Raceland", AwayTeamScore: 0},
//			{HomeTeamName: "Ashland Blazer", HomeTeamScore: 1, AwayTeamName: "Harlan County", AwayTeamScore: 0},
//			{HomeTeamName: "Ashland Blazer", HomeTeamScore: 1, AwayTeamName: "Greenup County", AwayTeamScore: 0},
//			{HomeTeamName: "Ashland Blazer", HomeTeamScore: 1, AwayTeamName: "Rowan County", AwayTeamScore: 0},
//			{HomeTeamName: "Ashland Blazer", HomeTeamScore: 0, AwayTeamName: "Johnson Central", AwayTeamScore: 4},
//			{HomeTeamName: "Ashland Blazer", HomeTeamScore: 1, AwayTeamName: "George Washington, WV", AwayTeamScore: 0},
//			{HomeTeamName: "Ashland Blazer", HomeTeamScore: 1, AwayTeamName: "Russell", AwayTeamScore: 0},
//			{HomeTeamName: "Ashland Blazer", HomeTeamScore: 1, AwayTeamName: "Ironton, OH", AwayTeamScore: 0},
//			{HomeTeamName: "Ashland Blazer", HomeTeamScore: 1, AwayTeamName: "Boyd County", AwayTeamScore: 0},
//			{HomeTeamName: "Ashland Blazer", HomeTeamScore: 1, AwayTeamName: "East Carter", AwayTeamScore: 0},
//		}
//	})
//
//	Context("NewRPICalculator", func() {
//		It("should initialize WPVAL to 0.35", func() {
//			calculator, err := pkg.NewRPICalculator([]pkg.Match{})
//
//			Expect(err).To(BeNil())
//			Expect(calculator).ToNot(BeNil())
//			Expect(calculator.WPVAL).To(Equal(0.35))
//		})
//
//		It("should initialize OWPVAL to 0.35", func() {
//			calculator, err := pkg.NewRPICalculator([]pkg.Match{})
//
//			Expect(err).To(BeNil())
//			Expect(calculator).ToNot(BeNil())
//			Expect(calculator.OWPVAL).To(Equal(0.35))
//		})
//
//		It("should initialize OOWPVAL to 0.30", func() {
//			calculator, err := pkg.NewRPICalculator([]pkg.Match{})
//
//			Expect(err).To(BeNil())
//			Expect(calculator).ToNot(BeNil())
//			Expect(calculator.OOWPVAL).To(Equal(0.30))
//		})
//
//		It("should initialize Matches to an empty slice", func() {
//			calculator, err := pkg.NewRPICalculator([]pkg.Match{})
//
//			Expect(err).To(BeNil())
//			Expect(calculator).ToNot(BeNil())
//			Expect(calculator.Matches).ToNot(BeNil())
//			Expect(len(calculator.Matches)).To(Equal(0))
//		})
//	})
//
//	Context("CalculateMatchesPlayed", func() {
//		It("should return 0 when there are no matches", func() {
//			// Arrange
//			calculator, _ := pkg.NewRPICalculator([]pkg.Match{})
//
//			// Act
//			count, _ := calculator.CalculateMatchesPlayed("something")
//
//			// Assert
//			Expect(count).To(Equal(0))
//		})
//
//		It("should return 1 when there is a single match [singular]", func() {
//			// Arrange
//			calculator, _ := pkg.NewRPICalculator([]pkg.Match{
//				{HomeTeamName: "team1", AwayTeamName: "team2"},
//			})
//
//			// Act
//			count, _ := calculator.CalculateMatchesPlayed("team2")
//
//			// Assert
//			Expect(count).To(Equal(1))
//		})
//
//		It("should return 1 when there is a single match [multiple]", func() {
//			// Arrange
//			calculator, _ := pkg.NewRPICalculator([]pkg.Match{
//				{HomeTeamName: "team1", AwayTeamName: "team2"},
//				{HomeTeamName: "team1", AwayTeamName: "team3"},
//			})
//
//			// Act
//			count, _ := calculator.CalculateMatchesPlayed("team2")
//
//			// Assert
//			Expect(count).To(Equal(1))
//		})
//
//		It("should return 2 when there is are two matches", func() {
//			// Arrange
//			calculator, _ := pkg.NewRPICalculator([]pkg.Match{
//				{HomeTeamName: "team1", AwayTeamName: "team2"},
//				{HomeTeamName: "team1", AwayTeamName: "team3"},
//				{HomeTeamName: "team4", AwayTeamName: "team3"},
//			})
//
//			// Act
//			count, _ := calculator.CalculateMatchesPlayed("team1")
//
//			// Assert
//			Expect(count).To(Equal(2))
//		})
//	})
//
//	Context("CalculateWinsLossesDraws", func() {
//		It("should return 0 when there are no matches", func() {
//			// Arrange
//			calculator, _ := pkg.NewRPICalculator([]pkg.Match{})
//
//			// Act
//			wins, losses, draws, _ := calculator.CalculateWinsLossesDraws("team1", "")
//
//			// Assert
//			Expect(wins).To(Equal(0))
//			Expect(losses).To(Equal(0))
//			Expect(draws).To(Equal(0))
//		})
//
//		It("should return the correct counts", func() {
//			// Arrange
//			calculator, _ := pkg.NewRPICalculator([]pkg.Match{
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 2},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 2},
//				{HomeTeamName: "team1", HomeTeamScore: 0, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 0, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 0, AwayTeamName: "team2", AwayTeamScore: 0},
//			})
//
//			// Act
//			wins, losses, draws, _ := calculator.CalculateWinsLossesDraws("team1", "")
//
//			// Assert
//			Expect(wins).To(Equal(1))
//			Expect(losses).To(Equal(2))
//			Expect(draws).To(Equal(3))
//		})
//	})
//
//	Context("Winning Percentage", func() {
//		It("should return 0.789 for a 15-4 team", func() {
//			// Arrange
//			calculator, _ := pkg.NewRPICalculator([]pkg.Match{
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 0, AwayTeamName: "team2", AwayTeamScore: 1},
//				{HomeTeamName: "team1", HomeTeamScore: 0, AwayTeamName: "team2", AwayTeamScore: 1},
//				{HomeTeamName: "team1", HomeTeamScore: 0, AwayTeamName: "team2", AwayTeamScore: 1},
//				{HomeTeamName: "team1", HomeTeamScore: 0, AwayTeamName: "team2", AwayTeamScore: 1},
//			})
//
//			// Act
//			wp, _ := calculator.CalculateWP("team1", "")
//
//			// Assert
//			Expect(wp).Should(BeNumerically("~", 0.789, 0.001))
//		})
//
//		It("should return 0.816 for a 15-3-1 team", func() {
//			// Arrange
//			calculator, _ := pkg.NewRPICalculator([]pkg.Match{
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 0},
//				{HomeTeamName: "team1", HomeTeamScore: 0, AwayTeamName: "team2", AwayTeamScore: 1},
//				{HomeTeamName: "team1", HomeTeamScore: 0, AwayTeamName: "team2", AwayTeamScore: 1},
//				{HomeTeamName: "team1", HomeTeamScore: 0, AwayTeamName: "team2", AwayTeamScore: 1},
//				{HomeTeamName: "team1", HomeTeamScore: 1, AwayTeamName: "team2", AwayTeamScore: 1},
//			})
//
//			// Act
//			wp, _ := calculator.CalculateWP("team1", "")
//
//			// Assert
//			Expect(wp).Should(BeNumerically("~", 0.816, 0.001))
//		})
//	})
//
//	Context("Opponent's Winning Percentage", func() {
//		BeforeEach(func() {
//			Expect(matches).ToNot(BeNil())
//		})
//		/*
//			OWP is the Opponent's Winning Percentage which is the average of the WP of each of the opponents
//			the team has played.
//
//			The WP for each of the opponents is calculated like the WP for the team with one big exception:
//			none of the games with the team who's RPI is being determined are included when calculating the
//			WP of the opponents.
//
//			So first we need to construct a slice containing all of the opponents of our target team.
//			Then one by one we need to calculate and sum up all of the WP's of the opponents
//			(minus any matches containing our target team).
//		*/
//	})
//})
