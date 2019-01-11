package runnertests_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRunnertests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Runnertests Suite")
}
