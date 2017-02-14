package vast

import (
	"encoding/xml"
	"os"
	"testing"
	"time"

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
			Expect(xml.NewDecoder(f).Decode(&v)).To(Succeed())
			Expect(v).To(Equal(*exp))
		},

		Entry("inline linear", "testdata/vast_inline_linear.xml", &VAST{
			Version: "2.0",
			Ads: []Ad{{
				ID: "601364",
				InLine: &InLine{
					AdSystem: &AdSystem{Version: "1.0", Name: "Acudeo Compatible"},
					AdTitle:  &AdTitle{Name: "VAST 2.0 Instream Test 1"},
					Impressions: []Impression{
						{URI: "http://myTrackingURL/impression"},
						{ID: "foo", URI: "http://myTrackingURL/impression2"},
					},
					Creatives: []Creative{
						{
							AdID: "601364",
							Linear: &Linear{
								Duration: durationPtr(30 * time.Second),
								TrackingEvents: []Tracking{
									{Event: "creativeView", URI: "http://myTrackingURL/creativeView"},
									{Event: "start", URI: "http://myTrackingURL/start"},
									{Event: "midpoint", URI: "http://myTrackingURL/midpoint"},
									{Event: "firstQuartile", URI: "http://myTrackingURL/firstQuartile"},
									{Event: "thirdQuartile", URI: "http://myTrackingURL/thirdQuartile"},
									{Event: "complete", URI: "http://myTrackingURL/complete"},
								},
								VideoClicks: &VideoClicks{
									ClickThroughs:  []VideoClick{{URI: "http://www.tremormedia.com"}},
									ClickTrackings: []VideoClick{{URI: "http://myTrackingURL/click"}},
								},
								MediaFiles: []MediaFile{
									{
										Delivery:            "progressive",
										Type:                "video/x-flv",
										Bitrate:             500,
										Width:               400,
										Height:              300,
										Scalable:            true,
										MaintainAspectRatio: true,
										URI:                 "http://cdnp.tremormedia.com/video/acudeo/Carrot_400x300_500kb.flv",
									},
								},
							},
						},
						{
							AdID: "601364-Companion",
							CompanionAds: &CompanionAds{
								Required: "all",
								Companions: []Companion{
									{
										Width:                 300,
										Height:                250,
										CompanionClickThrough: &CompanionClickThrough{URI: "http://www.tremormedia.com"},
										StaticResource:        &StaticResource{CreativeType: "image/jpeg", URI: "http://demo.tremormedia.com/proddev/vast/Blistex1.jpg"},
										TrackingEvents: []Tracking{
											{Event: "creativeView", URI: "http://myTrackingURL/firstCompanionCreativeView"},
										},
									},
									{
										Width:                 728,
										Height:                90,
										CompanionClickThrough: &CompanionClickThrough{URI: "http://www.tremormedia.com"},
										StaticResource:        &StaticResource{CreativeType: "image/jpeg", URI: "http://demo.tremormedia.com/proddev/vast/728x90_banner1.jpg"},
									},
								},
							},
						},
					},
					Description: "VAST 2.0 Instream Test 1",
					Error: []Error{
						{URI: "http://myErrorURL/error"},
						{URI: "http://myErrorURL/error2"},
					},
				},
			}},
		}),

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

func durationPtr(d time.Duration) *Duration {
	v := Duration(d)
	return &v
}
