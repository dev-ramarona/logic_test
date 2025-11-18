package soal2

import (
	"fmt"
)

func Soal2() {

	// Parameter
	routeFull := "CGK-UPG-DJJ-MKQ-SOQ"

	// Output
	routeBuy := map[string]string{
		"CGK-UPG": "",
		"CGK-DJJ": "",
		"CGK-MKQ": "",
		"CGK-SOQ": "",
		"UPG-DJJ": "",
		"UPG-MKQ": "",
		"UPG-SOQ": "",
		"DJJ-MKQ": "",
		"DJJ-SOQ": "",
		"MKQ-SOQ": "",
	}

	//////////////////// Start Write your code answer in here

	//////////////////// End Write your code answer in here

	// Check output
	routeAnswerkey := map[string]string{
		"CGK-UPG": "CGK-UPG",
		"CGK-DJJ": "CGK-UPG-DJJ",
		"CGK-MKQ": "CGK-UPG-DJJ-MKQ",
		"CGK-SOQ": "CGK-UPG-DJJ-MKQ-SOQ",
		"UPG-DJJ": "UPG-DJJ",
		"UPG-MKQ": "UPG-DJJ-MKQ",
		"UPG-SOQ": "UPG-DJJ-MKQ-SOQ",
		"DJJ-MKQ": "DJJ-MKQ",
		"DJJ-SOQ": "DJJ-MKQ-SOQ",
		"MKQ-SOQ": "MKQ-SOQ",
	}
	fmt.Println(routeFull)
	for k, v := range routeBuy {
		fmt.Printf("Route buy:%v | Route VCR:%v\n", k, v)
		if routeAnswerkey[k] != v {
			fmt.Println("GAGAL")
			return
		}
	}
	fmt.Println("BERHASIL")
}
