package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	// Compare()
	data := []string{"1", "2", "3"}
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		if i == 5 {
			data = []string{"1", "2", "3", "4", "5"}
		}
		fmt.Println(data)
		time.Sleep(time.Second * 2)
	}

}

var kelainanFisikMentalOrder = map[int]string{
	1: "A",
	2: "TA",
}

var kelainanFisikMental = map[string]string{
	"A":  "Ada",
	"TA": "Tidak Ada",
}

func KelainanFisikMentalListName() []string {
	var key []int
	for k := range kelainanFisikMentalOrder {
		key = append(key, k)
	}

	sort.Ints(key)

	name := make([]string, 0)
	for _, k := range key {
		name = append(name, kelainanFisikMental[kelainanFisikMentalOrder[k]])
	}
	return name
}

func KelainanFisikMentalCodeByName(name string) string {
	code := new(string)
	for k, v := range kelainanFisikMental {
		if v == name {
			code = &k
			break
		}
	}
	return *code
}

type Dusun struct {
	Name string
}

func Filter(data []Dusun) []string {
	dataMap := map[string]string{}
	for _, d := range data {
		dataMap[d.Name] = d.Name
	}

	var listString []string
	for d := range dataMap {
		listString = append(listString, d)
	}
	return listString
}
