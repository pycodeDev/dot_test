package entities

type Peminjaman struct {
	ID      int32  `json:"id" gorm:"primaryKey"`
	IdUser  int32  `json:"id_user"`
	IdBuku  int32  `json:"id_buku"`
	Status  int32  `json:"status"`
	Created string `json:"created"`
	Expired string `json:"expired"`
}
