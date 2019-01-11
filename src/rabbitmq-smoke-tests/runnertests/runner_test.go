package runnertests_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"rabbitmq-smoke-tests/tests/helper"
)

var _ = Describe("Runner", func() {

	It("runs", func() {
		metrics := helper.RetrieveMetricsFromFirehose("fjan-preprovisioned")
		Expect(metrics).To(ContainSubstring("/p.rabbitmq/rabbitmq/erlang/reachable_nodes"))
		Expect(metrics).To(ContainSubstring("/p.rabbitmq/rabbitmq/queues/count"))
	})
})
