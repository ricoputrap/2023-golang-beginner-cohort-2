package data

import "fmt"

func (p TPerson) PrintData() {
	fmt.Println("ID:", p.ID);
	fmt.Println("Nama:", p.Nama);
	fmt.Printf("Alamat: %s, No %s, %s, %s, %s\n",
		p.Alamat.Jalan,
		p.Alamat.No,
		p.Alamat.Kelurahan,
		p.Alamat.Kecamatan,
		p.Alamat.Kota,
	);
	fmt.Println("Alasan:", p.Alasan);
}