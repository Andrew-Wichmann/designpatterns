package main

type PatObservationGroup struct {
	impressions int
	spend       int
}

func (og *PatObservationGroup) getMetrics() Metrics {
	return Metrics{volumeField: og.impressions, spend: og.spend}
}
func (og *PatObservationGroup) getVendorName() string {
	return "PAT"
}

func newPatOg() *PatObservationGroup {
	return &PatObservationGroup{impressions: 1001}
}
