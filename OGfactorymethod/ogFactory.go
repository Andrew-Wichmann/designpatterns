package main

func createOG(vendor string) iObservationGroup {
	if vendor == "PAT" {
		return newPatOg()
	}
	if vendor == "EDS" {
		return newEDSOg()
	}
	println("Unsupported vendor")
	return nil
}
