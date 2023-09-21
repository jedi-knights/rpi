package schedule_test

import (
	"github.com/jedi-knights/rpi/pkg/schedule"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Schedule", func() {
	var pSchedule *schedule.Schedule

	BeforeEach(func() {
		pSchedule = schedule.NewSchedule()

		pSchedule.AddMatchFromString("UConn,64,Kansas,57")
		pSchedule.AddMatchFromString("UConn,82,Duke,68")
		pSchedule.AddMatchFromString("Wisconsin,71,UConn,72")
		pSchedule.AddMatchFromString("Kansas,69,UConn,62")
		pSchedule.AddMatchFromString("Duke,81,Wisconsin,70")
		pSchedule.AddMatchFromString("Wisconsin,52,Kansas,62")
	})

	AfterEach(func() {
		pSchedule = nil
	})

	Describe("GetMatchesForTeam", func() {
		It("should return all matches for a team - case #1", func() {
			// Act
			matches := pSchedule.GetMatchesForTeam("UConn")

			// Assert
			Expect(len(matches)).To(Equal(4))
			Expect(matches[0].ToString()).To(Equal("UConn,64,Kansas,57"))
			Expect(matches[1].ToString()).To(Equal("UConn,82,Duke,68"))
			Expect(matches[2].ToString()).To(Equal("Wisconsin,71,UConn,72"))
			Expect(matches[3].ToString()).To(Equal("Kansas,69,UConn,62"))
		})

		It("should return all matches for a team - case #2", func() {
			// Act
			matches := pSchedule.GetMatchesForTeam("Wisconsin")

			// Assert
			Expect(len(matches)).To(Equal(3))
			Expect(matches[0].ToString()).To(Equal("Wisconsin,71,UConn,72"))
			Expect(matches[1].ToString()).To(Equal("Duke,81,Wisconsin,70"))
			Expect(matches[2].ToString()).To(Equal("Wisconsin,52,Kansas,62"))
		})

		It("should return all matches for a team - case #3", func() {
			// Act
			matches := pSchedule.GetMatchesForTeam("Duke")

			// Assert
			Expect(len(matches)).To(Equal(2))
			Expect(matches[0].ToString()).To(Equal("UConn,82,Duke,68"))
			Expect(matches[1].ToString()).To(Equal("Duke,81,Wisconsin,70"))
		})

		It("should return all matches for a team - case #4", func() {
			// Act
			matches := pSchedule.GetMatchesForTeam("Kansas")

			// Assert
			Expect(len(matches)).To(Equal(3))
			Expect(matches[0].ToString()).To(Equal("UConn,64,Kansas,57"))
			Expect(matches[1].ToString()).To(Equal("Kansas,69,UConn,62"))
			Expect(matches[2].ToString()).To(Equal("Wisconsin,52,Kansas,62"))
		})
	})

	Describe("GetOpponents", func() {
		It("should return all opponents for UConn", func() {
			// Act
			opponents, err := pSchedule.GetOpponents("UConn")

			// Assert
			Expect(err).NotTo(HaveOccurred())

			Expect(len(opponents)).To(Equal(3))
			Expect(opponents[0]).To(Equal("Kansas"))
			Expect(opponents[1]).To(Equal("Duke"))
			Expect(opponents[2]).To(Equal("Wisconsin"))
		})

		It("should return all opponents for Kansas", func() {
			// Act
			opponents, err := pSchedule.GetOpponents("Kansas")

			// Assert
			Expect(err).NotTo(HaveOccurred())

			Expect(len(opponents)).To(Equal(2))
			Expect(opponents[0]).To(Equal("UConn"))
			Expect(opponents[1]).To(Equal("Wisconsin"))
		})

		It("should return all opponents for Duke", func() {
			// Act
			opponents, err := pSchedule.GetOpponents("Duke")

			// Assert
			Expect(err).NotTo(HaveOccurred())

			Expect(len(opponents)).To(Equal(2))
			Expect(opponents[0]).To(Equal("UConn"))
			Expect(opponents[1]).To(Equal("Wisconsin"))
		})

		It("should return all opponents for Wisconsin", func() {
			// Act
			opponents, err := pSchedule.GetOpponents("Wisconsin")

			// Assert
			Expect(err).NotTo(HaveOccurred())

			Expect(len(opponents)).To(Equal(3))
			Expect(opponents[0]).To(Equal("UConn"))
			Expect(opponents[1]).To(Equal("Duke"))
			Expect(opponents[2]).To(Equal("Kansas"))
		})

		It("should return an error for an empty team name", func() {
			// Act
			_, err := pSchedule.GetOpponents("")

			// Assert
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the specified team name is empty"))
		})
	})

	Describe("GetMatchesPlayedBy", func() {
		It("should return all matches for UConn", func() {
			// Act
			matches, err := pSchedule.GetMatchesPlayedBy("UConn")

			// Assert
			Expect(err).NotTo(HaveOccurred())

			Expect(len(matches)).To(Equal(4))
			Expect(matches[0].ToString()).To(Equal("UConn,64,Kansas,57"))
			Expect(matches[1].ToString()).To(Equal("UConn,82,Duke,68"))
			Expect(matches[2].ToString()).To(Equal("Wisconsin,71,UConn,72"))
			Expect(matches[3].ToString()).To(Equal("Kansas,69,UConn,62"))
		})

		It("should return all matches for Kansas", func() {
			// Act
			matches, err := pSchedule.GetMatchesPlayedBy("Kansas")

			// Assert
			Expect(err).NotTo(HaveOccurred())

			Expect(len(matches)).To(Equal(3))
			Expect(matches[0].ToString()).To(Equal("UConn,64,Kansas,57"))
			Expect(matches[1].ToString()).To(Equal("Kansas,69,UConn,62"))
			Expect(matches[2].ToString()).To(Equal("Wisconsin,52,Kansas,62"))
		})

		It("should return all matches for Duke", func() {
			// Act
			matches, err := pSchedule.GetMatchesPlayedBy("Duke")

			// Assert
			Expect(err).NotTo(HaveOccurred())

			Expect(len(matches)).To(Equal(2))
			Expect(matches[0].ToString()).To(Equal("UConn,82,Duke,68"))
			Expect(matches[1].ToString()).To(Equal("Duke,81,Wisconsin,70"))
		})

		It("should return all matches for Wisconsin", func() {
			// Act
			matches, err := pSchedule.GetMatchesPlayedBy("Wisconsin")

			// Assert
			Expect(err).NotTo(HaveOccurred())

			Expect(len(matches)).To(Equal(3))
			Expect(matches[0].ToString()).To(Equal("Wisconsin,71,UConn,72"))
			Expect(matches[1].ToString()).To(Equal("Duke,81,Wisconsin,70"))
			Expect(matches[2].ToString()).To(Equal("Wisconsin,52,Kansas,62"))
		})

		It("should return an error for an empty team name", func() {
			// Act
			matchesPlayed, err := pSchedule.GetMatchesPlayedBy("")

			// Assert
			Expect(matchesPlayed).To(BeNil())

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the specified team name is empty"))
		})

		It("should return an error for a team that doesn't exist", func() {
			// Act
			matchesPlayed, err := pSchedule.GetMatchesPlayedBy("Foo")

			// Assert
			Expect(matchesPlayed).To(BeNil())

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("no matches found for team Foo"))
		})
	})

	Describe("GetWinsForTeam", func() {
		It("should return the correct number of wins for UConn", func() {
			// Act
			wins, err := pSchedule.GetWinsForTeam("UConn", "")

			// Assert
			Expect(wins).To(Equal(3))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct number of wins for Kansas", func() {
			// Act
			wins, err := pSchedule.GetWinsForTeam("Kansas", "")

			// Assert
			Expect(wins).To(Equal(2))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct number of wins for Duke", func() {
			// Act
			wins, err := pSchedule.GetWinsForTeam("Duke", "")

			// Assert
			Expect(wins).To(Equal(1))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct number of wins for Wisconsin", func() {
			// Act
			wins, err := pSchedule.GetWinsForTeam("Wisconsin", "")

			// Assert
			Expect(wins).To(Equal(0))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return an error for an empty team name", func() {
			// Act
			wins, err := pSchedule.GetWinsForTeam("", "")

			// Assert
			Expect(wins).To(Equal(0))

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the specified team name is empty"))
		})

		It("should return an error for a team that doesn't exist", func() {
			// Act
			wins, err := pSchedule.GetWinsForTeam("Foo", "")

			// Assert
			Expect(wins).To(Equal(0))

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("no matches found for team Foo"))
		})
	})

	Describe("GetLossesForTeam", func() {
		It("should return the correct number of losses for UConn", func() {
			// Act
			losses, err := pSchedule.GetLossesForTeam("UConn", "")

			// Assert
			Expect(losses).To(Equal(1))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct number of losses for Kansas", func() {
			// Act
			losses, err := pSchedule.GetLossesForTeam("Kansas", "")

			// Assert
			Expect(losses).To(Equal(1))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct number of losses for Duke", func() {
			// Act
			losses, err := pSchedule.GetLossesForTeam("Duke", "")

			// Assert
			Expect(losses).To(Equal(1))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct number of losses for Wisconsin", func() {
			// Act
			losses, err := pSchedule.GetLossesForTeam("Wisconsin", "")

			// Assert
			Expect(losses).To(Equal(3))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return an error for an empty team name", func() {
			// Act
			losses, err := pSchedule.GetLossesForTeam("", "")

			// Assert
			Expect(losses).To(Equal(0))

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the specified team name is empty"))
		})

		It("should return an error for a team that doesn't exist", func() {
			// Act
			losses, err := pSchedule.GetLossesForTeam("Foo", "")

			// Assert
			Expect(losses).To(Equal(0))

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("no matches found for team Foo"))
		})
	})

	Describe("GetTiesForTeam", func() {
		It("should return the correct number of ties for UConn", func() {
			// Act
			ties, err := pSchedule.GetTiesForTeam("UConn", "")

			// Assert
			Expect(ties).To(Equal(0))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct number of ties for Kansas", func() {
			// Act
			ties, err := pSchedule.GetTiesForTeam("Kansas", "")

			// Assert
			Expect(ties).To(Equal(0))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct number of ties for Duke", func() {
			// Act
			ties, err := pSchedule.GetTiesForTeam("Duke", "")

			// Assert
			Expect(ties).To(Equal(0))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct number of ties for Wisconsin", func() {
			// Act
			ties, err := pSchedule.GetTiesForTeam("Wisconsin", "")

			// Assert
			Expect(ties).To(Equal(0))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return an error for an empty team name", func() {
			// Act
			ties, err := pSchedule.GetTiesForTeam("", "")

			// Assert
			Expect(ties).To(Equal(0))

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the specified team name is empty"))
		})

		It("should return an error for a team that doesn't exist", func() {
			// Act
			ties, err := pSchedule.GetTiesForTeam("Foo", "")

			// Assert
			Expect(ties).To(Equal(0))

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("no matches found for team Foo"))
		})
	})

	Describe("GetTotalMatchesPlayedForTeam", func() {
		It("should return the correct number of matches played for UConn", func() {
			// Act
			totalMatchesPlayed, err := pSchedule.GetTotalMatchesPlayedForTeam("UConn")

			// Assert
			Expect(totalMatchesPlayed).To(Equal(4))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct number of matches played for Kansas", func() {
			// Act
			totalMatchesPlayed, err := pSchedule.GetTotalMatchesPlayedForTeam("Kansas")

			// Assert
			Expect(totalMatchesPlayed).To(Equal(3))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct number of matches played for Duke", func() {
			// Act
			totalMatchesPlayed, err := pSchedule.GetTotalMatchesPlayedForTeam("Duke")

			// Assert
			Expect(totalMatchesPlayed).To(Equal(2))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct number of matches played for Wisconsin", func() {
			// Act
			totalMatchesPlayed, err := pSchedule.GetTotalMatchesPlayedForTeam("Wisconsin")

			// Assert
			Expect(totalMatchesPlayed).To(Equal(3))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return an error for an empty team name", func() {
			// Act
			totalMatchesPlayed, err := pSchedule.GetTotalMatchesPlayedForTeam("")

			// Assert
			Expect(totalMatchesPlayed).To(Equal(0))

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the specified team name is empty"))
		})

		It("should return an error for a team that doesn't exist", func() {
			// Act
			totalMatchesPlayed, err := pSchedule.GetTotalMatchesPlayedForTeam("Foo")

			// Assert
			Expect(totalMatchesPlayed).To(Equal(0))

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("no matches found for team Foo"))
		})
	})

	Describe("GetTotalMatchesPlayed", func() {
		It("should return the correct number of matches played", func() {
			// Act
			totalMatchesPlayed := pSchedule.GetTotalMatchesPlayed()

			// Assert
			Expect(totalMatchesPlayed).To(Equal(6))
		})
	})

	Describe("CalculateWP", func() {
		It("should return the correct WP for UConn", func() {
			// Act
			wp, err := pSchedule.CalculateWP("UConn", "")

			// Assert
			Expect(wp).Should(BeNumerically("~", 0.7500, 0.0001))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct WP for Kansas", func() {
			// Act
			wp, err := pSchedule.CalculateWP("Kansas", "")

			// Assert
			Expect(wp).Should(BeNumerically("~", 0.6667, 0.0001))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct WP for Duke", func() {
			// Act
			wp, err := pSchedule.CalculateWP("Duke", "")

			// Assert
			Expect(wp).Should(BeNumerically("~", 0.5000, 0.0001))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct WP for Wisconsin", func() {
			// Act
			wp, err := pSchedule.CalculateWP("Wisconsin", "")

			// Assert
			Expect(wp).Should(BeNumerically("~", 0.0000, 0.0001))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return an error for an empty team name", func() {
			// Act
			wp, err := pSchedule.CalculateWP("", "")

			// Assert
			Expect(wp).To(Equal(0.0))

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the specified team name is empty"))
		})

		It("should return an error for a team that doesn't exist", func() {
			// Act
			wp, err := pSchedule.CalculateWP("Foo", "")

			// Assert
			Expect(wp).To(Equal(0.0))

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("no matches found for team Foo"))
		})
	})

	Describe("CalculateOWP", func() {
		It("should return the correct OWP for UConn", func() {
			// Act
			owp, err := pSchedule.CalculateOWP("UConn")

			// Assert
			Expect(owp).Should(BeNumerically("~", 0.7500, 0.0001))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct OWP for Kansas", func() {
			// Act
			owp, err := pSchedule.CalculateOWP("Kansas")

			// Assert
			Expect(owp).Should(BeNumerically("~", 0.6667, 0.0001))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct OWP for Duke", func() {
			// Act
			owp, err := pSchedule.CalculateOWP("Duke")

			// Assert
			Expect(owp).Should(BeNumerically("~", 0.3333, 0.0001))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct OWP for Wisconsin", func() {
			// Act
			owp, err := pSchedule.CalculateOWP("Wisconsin")

			// Assert
			Expect(owp).Should(BeNumerically("~", 0.3889, 0.0001))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return an error for an empty team name", func() {
			// Act
			owp, err := pSchedule.CalculateOWP("")

			// Assert
			Expect(owp).To(Equal(0.0))

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the specified team name is empty"))
		})

		It("should return an error for a team that doesn't exist", func() {
			// Act
			owp, err := pSchedule.CalculateOWP("Foo")

			// Assert
			Expect(owp).To(Equal(0.0))

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("no matches found for team Foo"))
		})
	})

	Describe("CalculateOOWP", func() {
		It("should return the correct OOWP for UConn", func() {
			// Act
			oowp, err := pSchedule.CalculateOOWP("UConn")

			// Assert
			Expect(oowp).Should(BeNumerically("~", 0.5139, 0.0001))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct OOWP for Kansas", func() {
			// Act
			oowp, err := pSchedule.CalculateOOWP("Kansas")

			// Assert
			Expect(oowp).Should(BeNumerically("~", 0.6296, 0.0001))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct OOWP for Duke", func() {
			// Act
			oowp, err := pSchedule.CalculateOOWP("Duke")

			// Assert
			Expect(oowp).Should(BeNumerically("~", 0.5694, 0.0001))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct OOWP for Wisconsin", func() {
			// Act
			oowp, err := pSchedule.CalculateOOWP("Wisconsin")

			// Assert
			Expect(oowp).Should(BeNumerically("~", 0.5833, 0.0001))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return an error for an empty team name", func() {
			// Act
			oowp, err := pSchedule.CalculateOOWP("")

			// Assert
			Expect(oowp).To(Equal(0.0))

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the specified team name is empty"))
		})

		It("should return an error for a team that doesn't exist", func() {
			// Act
			oowp, err := pSchedule.CalculateOOWP("Foo")

			// Assert
			Expect(oowp).To(Equal(0.0))

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("no matches found for team Foo"))
		})
	})

	Context("UConn", func() {
		It("should calculate the correct winning percentage for UConn", func() {
			// Act
			wp, err := pSchedule.CalculateWP("UConn", "")

			// Assert
			Expect(wp).Should(BeNumerically("~", 0.7500, 0.0001))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should calculate the correct opponents' winning percentage for UConn", func() {
			// Act
			owp, err := pSchedule.CalculateOWP("UConn")

			// Assert
			Expect(owp).Should(BeNumerically("~", 0.7500, 0.0001))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should calculate the correct opponents' opponents' winning percentage for UConn", func() {
			// Act
			oowp, err := pSchedule.CalculateOOWP("UConn")

			// Assert
			Expect(oowp).Should(BeNumerically("~", 0.5139, 0.0001))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should calculate the correct RPI for UConn", func() {
			// Act
			rpi, err := pSchedule.CalculateRPI("UConn")

			// Assert
			Expect(rpi).Should(BeNumerically("~", 0.7066, 0.1))

			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("CalculateRPI", func() {
		It("should return the correct RPI for UConn", func() {
			// Act
			rpi, err := pSchedule.CalculateRPI("UConn")

			// Assert
			Expect(rpi).Should(BeNumerically("~", 0.7066, 0.1))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct RPI for Kansas", func() {
			// Act
			rpi, err := pSchedule.CalculateRPI("Kansas")

			// Assert
			Expect(rpi).Should(BeNumerically("~", 0.6830, 0.1))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct RPI for Duke", func() {
			// Act
			rpi, err := pSchedule.CalculateRPI("Duke")

			// Assert
			Expect(rpi).Should(BeNumerically("~", 0.4340, 0.1))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return the correct RPI for Wisconsin", func() {
			// Act
			rpi, err := pSchedule.CalculateRPI("Wisconsin")

			// Assert
			Expect(rpi).Should(BeNumerically("~", 0.3403, 0.1))

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return an error for an empty team name", func() {
			// Act
			rpi, err := pSchedule.CalculateRPI("")

			// Assert
			Expect(rpi).To(Equal(0.0))

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("the specified team name is empty"))
		})

		It("should return an error for a team that doesn't exist", func() {
			// Act
			rpi, err := pSchedule.CalculateRPI("Foo")

			// Assert
			Expect(rpi).To(Equal(0.0))

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("no matches found for team Foo"))
		})
	})
})
