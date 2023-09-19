package accumulators_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAccumulators(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Accumulators Suite")
}
