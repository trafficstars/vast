package vast

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Offset", func() {
	zero := Duration(0)

	DescribeTable("marshal",
		func(o *Offset, exp string) {
			b, err := o.MarshalText()
			Expect(err).NotTo(HaveOccurred())
			Expect(string(b)).To(Equal(exp))
		},
		Entry("0%", &Offset{}, "0%"),
		Entry("10%", &Offset{Percent: 0.1}, "10%"),
		Entry("00:00:00", &Offset{Duration: &zero}, "00:00:00"),
	)

	DescribeTable("unmarshal",
		func(s string, pc float64, dur *Duration) {
			o := Offset{}
			Expect(o.UnmarshalText([]byte(s))).To(Succeed())
			Expect(o.Percent).To(BeNumerically("~", pc, 0.001))
			Expect(o.Duration).To(Equal(dur))
		},
		Entry("0%", "0%", 0.0, nil),
		Entry("10%", "10%", 0.1, nil),
		Entry("00:00:00", "00:00:00", 0.0, &zero),
	)

	It("should fail to unmarshal bad inputs", func() {
		o := new(Offset)
		Expect(o.UnmarshalText([]byte("abc%"))).To(MatchError("invalid offset: abc%"))
	})

})
