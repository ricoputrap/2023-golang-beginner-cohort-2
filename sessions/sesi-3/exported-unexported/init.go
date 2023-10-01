package main

import "fmt"

// Will be executed first, before main()
// biasanya berisi utk inisiasi konfigurasi2
// spt konfig utk buka koneksi ke DB
func init() {
	fmt.Println("INITTT")
}