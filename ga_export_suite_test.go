package ga

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGaExport(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GaExport Suite")
}
