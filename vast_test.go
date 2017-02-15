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
						{URI: URIString("http://myTrackingURL/impression")},
						{ID: "foo", URI: URIString("http://myTrackingURL/impression2")},
					},
					Creatives: []Creative{
						{
							AdID: "601364",
							Linear: &Linear{
								Duration: durationPtr(30 * time.Second),
								TrackingEvents: []Tracking{
									{Event: "creativeView", URI: URIString("http://myTrackingURL/creativeView")},
									{Event: "start", URI: URIString("http://myTrackingURL/start")},
									{Event: "midpoint", URI: URIString("http://myTrackingURL/midpoint")},
									{Event: "firstQuartile", URI: URIString("http://myTrackingURL/firstQuartile")},
									{Event: "thirdQuartile", URI: URIString("http://myTrackingURL/thirdQuartile")},
									{Event: "complete", URI: URIString("http://myTrackingURL/complete")},
								},
								VideoClicks: &VideoClicks{
									ClickThroughs:  []VideoClick{{URI: URIString("http://www.tremormedia.com")}},
									ClickTrackings: []VideoClick{{URI: URIString("http://myTrackingURL/click")}},
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
										URI:                 URIString("http://cdnp.tremormedia.com/video/acudeo/Carrot_400x300_500kb.flv"),
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
										CompanionClickThrough: &CompanionClickThrough{URI: URIString("http://www.tremormedia.com")},
										StaticResource:        &StaticResource{CreativeType: "image/jpeg", URI: URIString("http://demo.tremormedia.com/proddev/vast/Blistex1.jpg")},
										TrackingEvents: []Tracking{
											{Event: "creativeView", URI: URIString("http://myTrackingURL/firstCompanionCreativeView")},
										},
									},
									{
										Width:                 728,
										Height:                90,
										CompanionClickThrough: &CompanionClickThrough{URI: URIString("http://www.tremormedia.com")},
										StaticResource:        &StaticResource{CreativeType: "image/jpeg", URI: URIString("http://demo.tremormedia.com/proddev/vast/728x90_banner1.jpg")},
									},
								},
							},
						},
					},
					Description: "VAST 2.0 Instream Test 1",
					Error: []Error{
						{URI: URIString("http://myErrorURL/error")},
						{URI: URIString("http://myErrorURL/error2")},
					},
				},
			}},
		}),

		Entry("inline nonlinear", "testdata/vast_inline_nonlinear.xml", &VAST{
			Version: "2.0",
			Ads: []Ad{
				{
					ID: "602678",
					InLine: &InLine{
						AdSystem:    &AdSystem{Name: "Acudeo Compatible"},
						AdTitle:     &AdTitle{Name: "NonLinear Test Campaign 1"},
						Description: "NonLinear Test Campaign 1",
						Survey:      "http://mySurveyURL/survey",
						Error:       []Error{{URI: URIString("http://myErrorURL/error")}},
						Impressions: []Impression{{URI: URIString("http://myTrackingURL/impression")}},
						Creatives: []Creative{
							{
								AdID: "602678-NonLinear",
								NonLinearAds: &NonLinearAds{
									TrackingEvents: []Tracking{
										{Event: "creativeView", URI: URIString("http://myTrackingURL/nonlinear/creativeView")},
										{Event: "expand", URI: URIString("http://myTrackingURL/nonlinear/expand")},
										{Event: "collapse", URI: URIString("http://myTrackingURL/nonlinear/collapse")},
										{Event: "acceptInvitation", URI: URIString("http://myTrackingURL/nonlinear/acceptInvitation")},
										{Event: "close", URI: URIString("http://myTrackingURL/nonlinear/close")},
									},
									NonLinears: []NonLinear{
										{
											Height:               50,
											Width:                300,
											MinSuggestedDuration: durationPtr(15 * time.Second),
											StaticResource: &StaticResource{
												CreativeType: "image/jpeg",
												URI:          URIString("http://demo.tremormedia.com/proddev/vast/50x300_static.jpg"),
											},
											NonLinearClickThrough: "http://www.tremormedia.com",
										},
										{
											Height:               50,
											Width:                450,
											MinSuggestedDuration: durationPtr(20 * time.Second),
											StaticResource: &StaticResource{
												CreativeType: "image/jpeg",
												URI:          URIString("http://demo.tremormedia.com/proddev/vast/50x450_static.jpg"),
											},
											NonLinearClickThrough: "http://www.tremormedia.com",
										},
									},
								},
							},
							{
								AdID: "602678-Companion",
								CompanionAds: &CompanionAds{
									Companions: []Companion{
										{
											Width:  300,
											Height: 250,
											StaticResource: &StaticResource{
												CreativeType: "application/x-shockwave-flash",
												URI:          URIString("http://demo.tremormedia.com/proddev/vast/300x250_companion_1.swf"),
											},
											CompanionClickThrough: &CompanionClickThrough{URI: URIString("http://www.tremormedia.com")},
										},
										{
											Width:  728,
											Height: 90,
											StaticResource: &StaticResource{
												CreativeType: "image/jpeg",
												URI:          URIString("http://demo.tremormedia.com/proddev/vast/728x90_banner1.jpg"),
											},
											TrackingEvents: []Tracking{
												{Event: "creativeView", URI: URIString("http://myTrackingURL/secondCompanion")},
											},
											CompanionClickThrough: &CompanionClickThrough{URI: URIString("http://www.tremormedia.com")},
										},
									},
								},
							},
						},
					},
				},
			},
		}),
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
