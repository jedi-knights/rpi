package math_test

import (
	"github.com/jedi-knights/rpi/pkg/math"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Math", func() {
	Describe("Average", func() {
		It("should return the average of a list of numbers", func() {
			// Arrange
			numbers := []float64{1.0, 2.0, 3.0}

			// Act
			average, err := math.ComputeAverage[float64](numbers)

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(average).To(Equal(2.0))
		})

		It("should return 0 when the list is empty", func() {
			// Arrange
			numbers := []float64{}

			// Act
			average, err := math.ComputeAverage[float64](numbers)

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(average).To(Equal(0.0))
		})

		It("should return 0 when the list is nil", func() {
			// Arrange
			var numbers []float64

			// Act
			average, err := math.ComputeAverage[float64](numbers)

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(average).To(Equal(0.0))
		})

		It("should return the value when the list has one value", func() {
			// Arrange
			numbers := []float64{1.0}

			// Act
			average, err := math.ComputeAverage[float64](numbers)

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(average).To(Equal(1.0))
		})

		It("should return the average of a list of integers", func() {
			// Arrange
			numbers := []int{1, 2, 3}

			// Act
			average, err := math.ComputeAverage[int](numbers)

			// Assert
			Expect(err).ToNot(HaveOccurred())
			Expect(average).To(Equal(2.0))
		})
	})
})
