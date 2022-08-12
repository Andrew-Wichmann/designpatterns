package main

type Metrics struct {
	volumeField int
	spend       int
}

type iObservationGroup interface {
	getMetrics() Metrics
	getVendorName() string
}
