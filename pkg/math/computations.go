package math

import (
	"fmt"
	"github.com/jedi-knights/rpi/pkg/errors"
	"math"
)

func ComputeAverage(numbers []float64) (float64, error) {
	total := 0.0

	if len(numbers) == 0 {
		return 0.0, nil
	}

	for _, number := range numbers {
		if math.IsInf(number, 1) {
			return 0.0, fmt.Errorf(errors.ErrPositiveInfinity)
		}

		if math.IsInf(number, -1) {
			return 0.0, fmt.Errorf(errors.ErrNegativeInfinity)
		}

		if math.IsNaN(number) {
			return 0.0, fmt.Errorf(errors.ErrNotANumber)
		}

		total += number
	}

	return total / float64(len(numbers)), nil
}
