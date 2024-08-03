package main

import "fmt"

type DuSl struct {
	ID   string
	Type string
	Name string
}

func Compare() {
	data1 := []DuSl{
		{ID: "1", Type: "A", Name: "Rambak"},
		{ID: "2", Type: "A", Name: "Rambaks"},
		{ID: "3", Type: "A", Name: "Rambaksd"},
		{ID: "4", Type: "A", Name: "Rambaksdf"},
	}
	data2 := []DuSl{
		{ID: "1", Type: "B", Name: "Rambak"},
		{ID: "2", Type: "B", Name: "Rambaks"},
		{ID: "3", Type: "A", Name: "Rambaksd"},
	}

	// var data3 []DuSl

	current := compareCurrentPersetujuan(data1, data2)
	fmt.Println(current)

	// fmt.Println("riyan")
}

func compareCurrentPersetujuan(before []DuSl, after []DuSl) (merger []DuSl) {
	for _, d := range before {
		carry := false
		for _, e := range after {
			if d.ID == e.ID {
				merger = append(merger, e)
			}
			if d.ID == e.ID {
				carry = true
				break
			}
		}
		if !carry {
			merger = append(merger, d)
		}
	}
	return
}
