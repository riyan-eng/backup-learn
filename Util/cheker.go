package main

import (
	"fmt"
)

func Checker() {
	fmt.Println("riyan")

	datang := []string{"b", "c", "d"}
	database := []string{"b", "c", "d", "f", "g"}
    add, rm:= getAddRm(datang, database)
    fmt.Println("add: ", add)
    fmt.Println("rm: ", rm)
}

func tambah(z string) {
	fmt.Println("tambah: ", z)
}

func hapus(z string) {
	fmt.Println("hapus: ", z)
}

func getAddRm(new, old []string) (add []string, rm []string) {
    diff := getDifference(new, old)
    for _, d := range diff{
        carry := false
        for _, e:= range old{
            if d == e{
                rm = append(rm, d)
                carry = true
                break
            }
        }
        if !carry{
            add = append(add, d)
        }
    }
    return
}

func getDifference(slice1 []string, slice2 []string) []string {
    var diff []string
    for i := 0; i < 2; i++ {
        for _, s1 := range slice1 {
            found := false
            for _, s2 := range slice2 {
                if s1 == s2 {
                    found = true
                    break
                }
            }
            if !found {
                diff = append(diff, s1)
            }
        }
        if i == 0 {
            slice1, slice2 = slice2, slice1
        }
    }
    return diff
}