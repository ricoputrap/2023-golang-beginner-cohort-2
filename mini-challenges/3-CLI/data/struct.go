package data

type TAlamat struct {
	Jalan			string
	No				string
	Kelurahan	string
	Kecamatan	string
	Kota			string
}

type TPerson struct {
	ID			int
	Nama		string
	Alamat	TAlamat
	Alasan	string
}