package main

import "fmt"

func main() {
	var car = map[string]string{}
	car["name"] = "BWM"
	car["color"] = "Black"

	// buat 2 buah fungsi :
	// 1 => fungsi yang mengembalikan sebuah string
	// pada fungsi ini terjadi pengolahan kata sehingga menghasilkan kata : Mobil BMW berwarna Black
	message := arrToString(car)

	// 2 => fungsi yang menampilkan hasil dari kembalian string
	// fungsi ini hanya bertugas untuk menampilkan kata
	showMessage(message)

	// alur
	// simpan hasil dari return function kedalam sebuah variable message
	// tampilkan hasil dari variable message

	// output => Mobil BMW berwarna Black

}

func arrToString(car map[string]string) string {
	return fmt.Sprintf("Mobil %s berwarna %s", car["name"], car["color"])
}

func showMessage(message string) {
	fmt.Println(message)
}
