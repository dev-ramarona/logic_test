package soal1

import (
	"fmt"
	"math"
	"time"
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
	for date := range dateSystem {
		strDatenw := time.Now().Format("060102")
		fmtDatenw, _ := time.Parse("060102", strDatenw)
		difFinald := 0
		for idx, yearvl := range []int{-1, 0, 1} {
			strYearnw := time.Now().AddDate(yearvl, 0, 0).Format("06")
			fmtSbrenw, _ := time.Parse("02Jan06", date+strYearnw)
			difDatenw := fmtDatenw.Sub(fmtSbrenw)
			difAbslte := int(math.Abs(difDatenw.Hours() / 24))
			// fmt.Println(difAbslte, fmtDatenw, fmtSbrenw)
			if idx == 0 {
				difFinald = difAbslte
				strDatenw = fmtSbrenw.Format("060102")
			} else if difFinald > difAbslte {
				difFinald = difAbslte
				strDatenw = fmtSbrenw.Format("060102")
			}
		}
		dateSystem[date] = strDatenw
	}
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
