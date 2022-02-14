package main

import "fmt"

//struct example begin
func structEx() {

	type StructName struct {
		ad    string
		soyad string
		// structArr []StructName
	}
	type structNameSub struct {
		StructName
		subStructName string
	}
	// myStruct := StructName{
	// 	ad:    "abuzer",
	// 	soyad: "kamil",
	// }

	//does't work.
	// arr, err := json.Marshal(myStruct)
	// fmt.Println(string(arr))
	// fmt.Println(myStruct)

	// myStruct2 := structNameSub{
	// 	StructName:    StructName{ad: "sutrctNameAd", soyad: "structNameSoyad"},
	// 	subStructName: "subStructName",
	// }
	// fmt.Println(myStruct2)

	// myStruct3 := structNameSub{}
	// myStruct3.ad = "structAd"
	// myStruct3.soyad = "structSoyad"
	// myStruct3.subStructName = "subStructName"

	// fmt.Println(myStruct3)

	// myStruct4 := structNameSub{}
	// myStruct4.StructName.ad = "structAd"
	// myStruct4.StructName.soyad = "structSoyad"
	// myStruct4.subStructName = "subStructName"

	// fmt.Println(myStruct4)

	// ////////////Anonim struct//////////
	// myStruct5 := struct {
	// 	name    string
	// 	surname string
	// }{name: "struct name", surname: "struct surname"}
	// fmt.Println(myStruct5)

	//////////////////////
	type nodeStruct struct {
		x int
		y int
	}
	obj := nodeStruct{x: 5, y: 10}
	fmt.Println(obj)
} //struct example end
