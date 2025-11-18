package soal3

import (
	"fmt"
)

func Soal3() {

	// Isi output dengan jawaban anda
	output := ""

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

	//////////////////// Start Write your code answer in here

	//////////////////// End Write your code answer in here

	// Check output
	fmt.Println(aroute)
	fmt.Printf("Ekpetasi:CGK-UPG-DJJ-MKQ-SOQ \nOutput:%s\n", output)
	if output == "CGK-UPG-DJJ-MKQ-SOQ" {
		fmt.Println("BERHASIL")
	}
}
