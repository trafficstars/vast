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
					Extensions: &Extensions{
						Extensions: []Extension{
							{
								Type: "geo",
								Data: []byte(`
			<Geo>
				<Country>US</Country>
				<State>CA</State>
			</Geo>
		`),
							},
						},
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
						Error:       []Error{{URI: "http://myErrorURL/error"}},
						Impressions: []Impression{{URI: "http://myTrackingURL/impression"}},
						Creatives: []Creative{
							{
								AdID: "602678-NonLinear",
								NonLinearAds: &NonLinearAds{
									TrackingEvents: []Tracking{
										{Event: "creativeView", URI: "http://myTrackingURL/nonlinear/creativeView"},
										{Event: "expand", URI: "http://myTrackingURL/nonlinear/expand"},
										{Event: "collapse", URI: "http://myTrackingURL/nonlinear/collapse"},
										{Event: "acceptInvitation", URI: "http://myTrackingURL/nonlinear/acceptInvitation"},
										{Event: "close", URI: "http://myTrackingURL/nonlinear/close"},
									},
									NonLinears: []NonLinear{
										{
											Height:               50,
											Width:                300,
											MinSuggestedDuration: durationPtr(15 * time.Second),
											StaticResource: &StaticResource{
												CreativeType: "image/jpeg",
												URI:          "http://demo.tremormedia.com/proddev/vast/50x300_static.jpg",
											},
											NonLinearClickThrough: "http://www.tremormedia.com",
										},
										{
											Height:               50,
											Width:                450,
											MinSuggestedDuration: durationPtr(20 * time.Second),
											StaticResource: &StaticResource{
												CreativeType: "image/jpeg",
												URI:          "http://demo.tremormedia.com/proddev/vast/50x450_static.jpg",
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
												URI:          "http://demo.tremormedia.com/proddev/vast/300x250_companion_1.swf",
											},
											CompanionClickThrough: &CompanionClickThrough{URI: "http://www.tremormedia.com"},
										},
										{
											Width:  728,
											Height: 90,
											StaticResource: &StaticResource{
												CreativeType: "image/jpeg",
												URI:          "http://demo.tremormedia.com/proddev/vast/728x90_banner1.jpg",
											},
											TrackingEvents: []Tracking{
												{Event: "creativeView", URI: "http://myTrackingURL/secondCompanion"},
											},
											CompanionClickThrough: &CompanionClickThrough{URI: "http://www.tremormedia.com"},
										},
									},
								},
							},
						},
					},
				},
			},
		}),

		Entry("wrapper linear", "testdata/vast_wrapper_linear_1.xml", &VAST{
			Version: "2.0",
			Ads: []Ad{
				{
					ID: "602833",
					Wrapper: &Wrapper{
						AdSystem:     &AdSystem{Name: "Acudeo Compatible"},
						VASTAdTagURI: "http://demo.tremormedia.com/proddev/vast/vast_inline_linear.xml",
						Error:        []Error{{URI: "http://myErrorURL/wrapper/error"}},
						Impressions:  []Impression{{URI: "http://myTrackingURL/wrapper/impression"}},
						Creatives: []CreativeWrapper{
							{
								AdID: "602833",
								Linear: &LinearWrapper{
									TrackingEvents: []Tracking{
										{Event: "creativeView", URI: "http://myTrackingURL/wrapper/creativeView"},
										{Event: "start", URI: "http://myTrackingURL/wrapper/start"},
										{Event: "midpoint", URI: "http://myTrackingURL/wrapper/midpoint"},
										{Event: "firstQuartile", URI: "http://myTrackingURL/wrapper/firstQuartile"},
										{Event: "thirdQuartile", URI: "http://myTrackingURL/wrapper/thirdQuartile"},
										{Event: "complete", URI: "http://myTrackingURL/wrapper/complete"},
										{Event: "mute", URI: "http://myTrackingURL/wrapper/mute"},
										{Event: "unmute", URI: "http://myTrackingURL/wrapper/unmute"},
										{Event: "pause", URI: "http://myTrackingURL/wrapper/pause"},
										{Event: "resume", URI: "http://myTrackingURL/wrapper/resume"},
										{Event: "fullscreen", URI: "http://myTrackingURL/wrapper/fullscreen"},
									},
								},
							},
							{
								Linear: &LinearWrapper{
									VideoClicks: &VideoClicks{
										ClickTrackings: []VideoClick{{URI: "http://myTrackingURL/wrapper/click"}},
									},
								},
							},
							{
								AdID: "602833-NonLinearTracking",
								NonLinearAds: &NonLinearAdsWrapper{
									TrackingEvents: []Tracking{
										{Event: "creativeView", URI: "http://myTrackingURL/wrapper/creativeView"},
									},
								},
							},
						},
					},
				},
			},
		}),

		Entry("wrapper nonlinear", "testdata/vast_wrapper_nonlinear_1.xml", &VAST{
			Version: "2.0",
			Ads: []Ad{
				{
					ID: "602867",
					Wrapper: &Wrapper{
						AdSystem:     &AdSystem{Name: "Acudeo Compatible"},
						VASTAdTagURI: "http://demo.tremormedia.com/proddev/vast/vast_inline_nonlinear2.xml",
						Error:        []Error{{URI: "http://myErrorURL/wrapper/error"}},
						Impressions:  []Impression{{URI: "http://myTrackingURL/wrapper/impression"}},
						Creatives: []CreativeWrapper{
							{
								AdID:   "602867",
								Linear: &LinearWrapper{},
							},
							{
								AdID: "602867-NonLinearTracking",
								NonLinearAds: &NonLinearAdsWrapper{
									TrackingEvents: []Tracking{
										{Event: "creativeView", URI: "http://myTrackingURL/wrapper/nonlinear/creativeView/creativeView"},
										{Event: "expand", URI: "http://myTrackingURL/wrapper/nonlinear/creativeView/expand"},
										{Event: "collapse", URI: "http://myTrackingURL/wrapper/nonlinear/creativeView/collapse"},
										{Event: "acceptInvitation", URI: "http://myTrackingURL/wrapper/nonlinear/creativeView/acceptInvitation"},
										{Event: "close", URI: "http://myTrackingURL/wrapper/nonlinear/creativeView/close"},
									},
								},
							},
						},
					},
				},
			},
		}),
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
