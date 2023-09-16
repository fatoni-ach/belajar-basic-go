package main

import (
	"errors"
	"fmt"
)

func bagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}

	return nilai / pembagi, nil
}
func main() {
	hasil, err := bagi(10, 2)
	if err != nil {
		fmt.Println("Hasilnya error : ", err.Error())
	} else {
		fmt.Println("Hasilnya : ", hasil)
	}

	hasil, err = bagi(10, 0)
	if err != nil {
		fmt.Println("Hasilnya error : ", err.Error())
	} else {
		fmt.Println("Hasilnya : ", hasil)
	}

}
