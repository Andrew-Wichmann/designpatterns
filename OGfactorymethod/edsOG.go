package main

type EDSObservationGroup struct {
	volume int
	spend  int
}

func (og *EDSObservationGroup) getMetrics() Metrics {
	return Metrics{volumeField: og.volume, spend: og.spend}
}
func (og *EDSObservationGroup) getVendorName() string {
	return "EDS"
}

func newEDSOg() *EDSObservationGroup {
	return &EDSObservationGroup{volume: 10}
}
