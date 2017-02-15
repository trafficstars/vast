package vast

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Duration", func() {

	DescribeTable("marshal",
		func(d Duration, exp string) {
			b, err := d.MarshalText()
			Expect(err).NotTo(HaveOccurred())
			Expect(string(b)).To(Equal(exp))
		},
		Entry("00:00:00", Duration(0), "00:00:00"),
		Entry("00:00:00.002", Duration(2*time.Millisecond), "00:00:00.002"),
		Entry("00:00:02", Duration(2*time.Second), "00:00:02"),
		Entry("00:02:00", Duration(2*time.Minute), "00:02:00"),
		Entry("02:00:00", Duration(2*time.Hour), "02:00:00"),
	)

	DescribeTable("unmarshal",
		func(s string, exp Duration) {
			d := new(Duration)
			Expect(d.UnmarshalText([]byte(s))).To(Succeed())
			Expect(*d).To(Equal(exp))
		},
		Entry("00:00:00", "00:00:00", Duration(0)),
		Entry("00:00:00.002", "00:00:00.002", Duration(2*time.Millisecond)),
		Entry("00:00:02", "00:00:02", Duration(2*time.Second)),
		Entry("00:02:00", "00:02:00", Duration(2*time.Minute)),
		Entry("02:00:00", "02:00:00", Duration(2*time.Hour)),
	)

	It("should fail to unmarshal bad inputs", func() {
		d := new(Duration)
		Expect(d.UnmarshalText([]byte("00:00:60"))).To(MatchError("invalid duration: 00:00:60"))
		Expect(d.UnmarshalText([]byte("00:60:00"))).To(MatchError("invalid duration: 00:60:00"))
		Expect(d.UnmarshalText([]byte("00:00:00.-1"))).To(MatchError("invalid duration: 00:00:00.-1"))
		Expect(d.UnmarshalText([]byte("00:00:00.1000"))).To(MatchError("invalid duration: 00:00:00.1000"))
		Expect(d.UnmarshalText([]byte("00h01m"))).To(MatchError("invalid duration: 00h01m"))
	})

})
