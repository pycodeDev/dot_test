package entities

type Buku struct {
	ID        int32  `json:"id" gorm:"primaryKey"`
	Nama      string `json:"nama"`
	Status    int32  `json:"status"`
	CreatedAt string `json:"created_at"`
}
