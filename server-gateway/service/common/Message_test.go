package common_test

import (
	. "github.com/a-trium/gipeline/server-gateway/service/common"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NumberRepository", func() {
	Describe("Submission", func() {
		Describe("NewSubmission", func() {
			It("should return valid submission instance", func() {
				s := NewSubmission("1ambda", 3)
				Expect(s).NotTo(Equal(nil))
			})
		})
	})
})
