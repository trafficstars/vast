package vast

import (
	"encoding/xml"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("VAST", func() {

	DescribeTable("parse",
		func(fixture string, exp *VAST) {
			f, err := os.Open(fixture)
			Expect(err).NotTo(HaveOccurred())
			defer f.Close()

			var v VAST
			err = xml.NewDecoder(f).Decode(&v)
			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal(*exp))
		},
		PEntry("inline linear", "testdata/vast_inline_linear.xml", &VAST{}),
		PEntry("inline nonlinear", "testdata/vast_inline_nonlinear.xml", &VAST{}),
		PEntry("wrapper linear", "testdata/vast_wrapper_linear_1.xml", &VAST{}),
		PEntry("wrapper nonlinear", "testdata/vast_wrapper_nonlinear_1.xml", &VAST{}),
	)

})

// --------------------------------------------------------------------

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "vast")
}
