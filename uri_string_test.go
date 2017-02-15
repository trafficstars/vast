package vast

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("URIString", func() {

	DescribeTable("marshal",
		func(d URIString, exp string) {
			b, err := d.MarshalText()
			Expect(err).NotTo(HaveOccurred())
			Expect(string(b)).To(Equal(exp))
		},
		Entry("", URIString(""), ""),
		Entry("http://example.com", URIString("http://example.com"), "http://example.com"),
	)

	DescribeTable("unmarshal",
		func(s string, exp URIString) {
			d := new(URIString)
			Expect(d.UnmarshalText([]byte(s))).To(Succeed())
			Expect(*d).To(Equal(exp))
		},
		Entry("Blank", "", URIString("")),
		Entry("Whitespace only", "\n\t ", URIString("")),
		Entry("Ideal Example", "http://example.com", URIString("http://example.com")),
		Entry("Real-world Example", "\n\t\t\t http://example.com \n\t\t\t", URIString("http://example.com")),
	)

})
