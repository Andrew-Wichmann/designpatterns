package main

import "fmt"

func main() {
	println("Enter a vendor name")
	var vendorName string
	_, err := fmt.Scanln(&vendorName)
	if err != nil {
		panic(err)
	}
	og := createOG(vendorName)
	if og != nil {
		if og.getMetrics().volumeField > 100 {
			println(fmt.Sprintf("Replicating %s OG to the website", og.getVendorName()))
		} else {
			println("This OG is too small to replicate to the website")
		}
	}
}
