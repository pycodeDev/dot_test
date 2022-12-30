package model

type ModelPeminjaman struct {
	ID      int    `json:"id"`
	Name    string `json:"nama_user"`
	Nama    string `json:"nama_buku"`
	Status  int    `json:"status"`
	Created string `json:"created"`
	Expired string `json:"expired"`
}
