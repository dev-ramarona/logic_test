package soal3

import (
	"fmt"
	"slices"
	"sort"
	"strings"
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

	// Main logic
	type keyval struct {
		Key   string
		Value int
	}
	arrdep := map[string]int{}
	arrarv := map[string]int{}
	for _, route1 := range aroute {
		route1split := strings.Split(route1, "-")
		if _, ist := arrdep[route1split[0]]; !ist {
			arrdep[route1split[0]] = 1
		} else {
			arrdep[route1split[0]] += 1
		}
		if _, ist := arrarv[route1split[1]]; !ist {
			arrarv[route1split[1]] = 1
		} else {
			arrarv[route1split[1]] += 1
		}
	}

	// Sort data
	arrdepobj := []keyval{}
	for key, val := range arrdep {
		arrdepobj = append(arrdepobj, keyval{key, val})
	}
	sort.Slice(arrdepobj, func(i, j int) bool {
		return arrdepobj[i].Value > arrdepobj[j].Value
	})
	arrarvobj := []keyval{}
	for key, val := range arrarv {
		arrarvobj = append(arrarvobj, keyval{key, val})
	}
	sort.Slice(arrarvobj, func(i, j int) bool {
		return arrarvobj[i].Value < arrarvobj[j].Value
	})

	// Push final data
	arrdepobj = append(arrdepobj, arrarvobj...)
	arrfinal := []string{}
	for _, val := range arrdepobj {
		if !slices.Contains(arrfinal, val.Key) {
			arrfinal = append(arrfinal, val.Key)
		}
	}
	output = strings.Join(arrfinal, "-")

	//////////////////// End Write your code answer in here

	// Check output
	fmt.Printf("Ekpetasi:CGK-UPG-DJJ-MKQ-SOQ \nOutput:%s\n", output)
	if output == "CGK-UPG-DJJ-MKQ-SOQ" {
		fmt.Println("BERHASIL")
	}
}
