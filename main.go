package main

import (
	"fmt"
)

func main() {

	// Output yang di harapkan
	expect := "CGK-UPG-DJJ-MKQ-SOQ" // Do not use this variable
	output := ""                    // Fill this variable

	// Parameter
	aroute := map[string]string{
		"DJJ-SOQ": "DJJ-SOQ",
		"UPG-SOQ": "UPG-SOQ",
		"CGK-UPG": "CGK-UPG",
		"CGK-DJJ": "CGK-DJJ",
		"CGK-SOQ": "CGK-SOQ",
		"DJJ-MKQ": "DJJ-MKQ",
		"MKQ-SOQ": "MKQ-SOQ",
		"UPG-MKQ": "UPG-MKQ",
		"UPG-DJJ": "UPG-DJJ",
		"CGK-MKQ": "CGK-MKQ",
	}

	//////////////////// Write your code in here

	//////////////////// Write your code in here

	// Report
	fmt.Println(aroute)
	fmt.Printf("Ekpetasi:%s \nOutput:%s\n", expect, output)
	if expect == output {
		fmt.Println("BERHASIL")
		return
	}
	fmt.Println("GAGAL")
}
