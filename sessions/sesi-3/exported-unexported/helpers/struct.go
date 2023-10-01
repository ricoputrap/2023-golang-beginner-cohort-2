package helpers

/*
Karena sudah define struct Person di sini,
maka kita tidak bisa define struct Person lagi di file lain
yg masih berada di package yg sama

Contoh:
Ga bisa lagi define struct Person di method.go
karena sudah define di struct.go
*/
type Person struct {
	Name  	string
	Age			int
	Addres	string
}