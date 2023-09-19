package calculators_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCalculators(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculators Suite")
}
