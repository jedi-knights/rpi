package math

import (
	"fmt"
	"github.com/jedi-knights/rpi/pkg/errors"
	"math"
)

func ComputeAverage[T int | int64 | float64](numbers []T) (float64, error) {
	var total = float64(0.0)

	if len(numbers) == 0 {
		return 0.0, nil
	}

	for _, number := range numbers {
		float64Number := float64(number)

		if math.IsInf(float64Number, 1) {
			return 0.0, fmt.Errorf(errors.ErrPositiveInfinity)
		}

		if math.IsInf(float64Number, -1) {
			return 0.0, fmt.Errorf(errors.ErrNegativeInfinity)
		}

		if math.IsNaN(float64Number) {
			return 0.0, fmt.Errorf(errors.ErrNotANumber)
		}

		total += float64Number
	}

	return total / float64(len(numbers)), nil
}
