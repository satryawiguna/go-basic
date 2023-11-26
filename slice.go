package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Made"
	name[1] = "Satrya"
	name[2] = "Wiguna"

	slice := name[0:2]
	var sliceDua []string = name[1:]
	var sliceTiga []string = name[:1]
	sliceEmpat := name[:]

	fmt.Println(slice)
	fmt.Println(sliceDua)
	fmt.Println(sliceTiga)
	fmt.Println(sliceEmpat)

	var days = [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	fmt.Println(days)

	daySliceSatu := days[5:]

	fmt.Println(daySliceSatu)

	daySliceSatu[0] = "_Sabtu_"
	daySliceSatu[1] = "_Minggu_"

	fmt.Println(days)

	daySliceDua := append(daySliceSatu, "Libur")

	fmt.Println(daySliceDua)

	daySliceDua[0] = "_Libur_"

	fmt.Println(daySliceDua)
	fmt.Println(days)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(toSlice)
}
