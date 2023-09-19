package calculators

type CalculatorInterface interface {
	Calculate(teamName string) (float64, error)
}
