package soal1

import (
	"fmt"
)

func Soal1() {

	// Output
	dateSystem := map[string]string{
		"01JAN": "",
		"11DEC": "",
		"10JUL": "",
		"30OCT": "",
		"15NOV": "",
	}

	//////////////////// Start Write your code answer in here

	//////////////////// End Write your code answer in here

	// Check output
	dateAnswerkey := map[string]string{
		"01JAN": "260101",
		"11DEC": "251211",
		"10JUL": "250710",
		"30OCT": "251030",
		"15NOV": "251115",
		"17MAR": "260317",
	}
	for k, v := range dateSystem {
		fmt.Printf("Date buy:%v | Date VCR:%v\n", k, v)
		if dateAnswerkey[k] != v {
			fmt.Println("GAGAL")
			return
		}
	}
	fmt.Println("BERHASIL")
}
