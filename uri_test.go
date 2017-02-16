package vast

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("URI", func() {

	DescribeTable("marshal",
		func(d URI, exp string) {
			b, err := d.MarshalText()
			Expect(err).NotTo(HaveOccurred())
			Expect(string(b)).To(Equal(exp))
		},
		Entry("", URI(""), ""),
		Entry("http://example.com", URI("http://example.com"), "http://example.com"),
	)

	DescribeTable("unmarshal",
		func(s string, exp URI) {
			d := new(URI)
			Expect(d.UnmarshalText([]byte(s))).To(Succeed())
			Expect(*d).To(Equal(exp))
		},
		Entry("Blank", "", URI("")),
		Entry("Whitespace only", "\n\t ", URI("")),
		Entry("Ideal Example", "http://example.com", URI("http://example.com")),
		Entry("Real-world Example", "\n\t\t\t http://example.com \n\t\t\t", URI("http://example.com")),
	)

	DescribeTable("stringer",
		func(d URI, exp string) {
			Expect(d.String()).To(Equal(exp))
		},
		Entry("Blank", URI(""), ""),
		Entry("http://example.com", URI("http://example.com"), "http://example.com"),
	)

})
