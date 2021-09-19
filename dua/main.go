package main

import (
	"fmt"
)

func luas_segitiga(a float64, t float64) string {
	luas := 0.5*a*t
	return fmt.Sprintf("Luas Segitiga adalah %f",luas)
}

func keliling_segitiga(s float64) string {
	keliling := 3*s
	return fmt.Sprintf("Keliling segitiga adalah %f",keliling)
}

func luas_persegi(s float64) string {
	luas := s*s
	return fmt.Sprintf("Luas Persegi adalah %f",luas)
}

func keliling_persegi(s float64) string {
	keliling := 4*s
	return fmt.Sprintf("Keliling Persegi adalah %f",keliling)
}

func menu() {
	var a float64
	var t float64
	var s float64
	var i int

	fmt.Printf("Pilih:\n")
	fmt.Printf("1. Hitung Luas Segitiga:\n")
	fmt.Printf("2. Hitung Keliling Segitiga :\n")
	fmt.Printf("3. Hitung Luas Persegi :\n")
	fmt.Printf("4. Hitung Keliling Persegi :\n")
	fmt.Printf("5. Keluar \n")
	fmt.Printf("Pilih Menu : ")
	fmt.Scanf("%d",&i)

	

	switch i{
		case 1:
			fmt.Print("alas : ")
			fmt.Scan(&a)

			fmt.Print("tinggi : ")
			fmt.Scan(&t)
			fmt.Println(luas_segitiga(a,t))
			menu()
		case 2:
			fmt.Print("sisi : ")
			fmt.Scan(&s)
			fmt.Println(keliling_segitiga(s))
			menu()
		case 3:
			fmt.Print("sisi : ")
			fmt.Scan(&s)
			fmt.Println(luas_persegi(s))
			menu()
		case 4:
			fmt.Print("sisi : ")
			fmt.Scan(&s)
			fmt.Println(keliling_persegi(s))
			menu()
		case 5:
			fmt.Print("Bye")
		default:fmt.Println("Tidak ada")
	}
}

func main() {
	menu()
}

