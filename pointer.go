package main

import "fmt"

type Address struct {
	region, district, address string
}

func changeRegion(address *Address) {
	address.region = "Singapore"
}

func (address *Address) changeDistrict() {
	address.district = "New " + address.district
}

func main() {
	myAddress := Address{
		region:   "Indonesia",
		district: "Surabaya",
		address:  "Ngagel_Mulyo",
	}

	changeRegion(&myAddress)
	fmt.Println(myAddress)

	// Pointer di method
	myAddress.changeDistrict()
	fmt.Println(myAddress)

}
