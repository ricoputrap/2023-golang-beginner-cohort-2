package main

import (
	data "3-CLI/data"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Selamat datang di FriendBook.");
	fmt.Println("Silahkan masukan nama atau nomor absen teman Anda.");
	fmt.Println("Contoh: 'go run main.go id 1' atau 'go run main.go nama Rico'");

	fmt.Println(strings.Repeat("-", 60))

	var people = []data.TPerson{
		{
			ID: 1,
			Nama: "Rico",
			Alamat: data.TAlamat{
				Jalan: "Jl HM Alif",
				No: "4A",
				Kelurahan: "Kukusan",
				Kecamatan: "Beji",
				Kota: "Depok",
			},
			Alasan: "Persiapan resign supaya lebih hireable",
		},
	};

	var args []string = os.Args[1:];
	
	if len(args) < 2 {
		fmt.Println("Mohon maaf perintah tidak valid.\nSilahkan dicoba kembali.");
		return;
	}

	var key string = args[0];

	if (key != "id") && (key != "nama") {
		fmt.Println("Mohon maaf perintah tidak valid.");
		fmt.Println("Silahkan masukkan 'id' atau 'nama' dalam perintah.")
		return;
	}

	var person *data.TPerson;
	var value string = args[1];

	if key == "id" {
		id, error := strconv.Atoi(value);

		if error != nil {
			fmt.Println("GAGAL Nilai 'id' tidak dapat dikonversi ke 'int'");
			return;
		}
		person = data.GetPersonByID(id, people);
	}

	if key == "nama" {
		person = data.GetPersonByNama(value, people);
	}

	if person == nil {
		fmt.Println("Data tidak tersedia.");
		return;
	}

	person.PrintData();
}