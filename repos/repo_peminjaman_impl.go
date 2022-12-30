package repos

import (
	"context"
	"errors"

	"dot.go/entities"
	"dot.go/helper"
	"dot.go/model"
	"gorm.io/gorm"
)

type RepoPeminjamanImpl struct {
	DB *gorm.DB
}

func NewRepoPeminjamanImpl(db *gorm.DB) RepoPeminjaman {
	return &RepoPeminjamanImpl{DB: db}
}

func (s RepoPeminjamanImpl) InsertPeminjaman(ctx context.Context, pinjam entities.Peminjaman) error {
	funcNow := "RepoPeminjamanImpl.InsertPeminjaman"

	var buku entities.Buku

	tx := s.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		helper.LogError(err.Error(), "func:"+funcNow, "script: tx error")
		return err
	}

	err := tx.WithContext(ctx).Raw("SELECT id, nama, status FROM bukus WHERE id = ? FOR UPDATE", pinjam.IdBuku).Scan(&buku)
	if err.Error != nil {
		tx.Rollback()
		helper.LogError(err.Error.Error(), "func:"+funcNow, "script: select buku")
		return err.Error
	}

	if buku.Status == 0 {
		tx.Rollback()
		return errors.New("Maaf, Buku Sedang Dipinjam !")
	}

	err = tx.WithContext(ctx).Table("peminjamans").Select("IdUser", "IdBuku", "Status", "Created", "Expired").Create(&pinjam)
	if err.Error != nil {
		tx.Rollback()
		helper.LogError(err.Error.Error(), "func:"+funcNow, "script: insert peminjaman")
		return err.Error
	}

	err = tx.WithContext(ctx).Table("bukus").Where("id = ?", buku.ID).Update("status", 0)
	if err.Error != nil {
		tx.Rollback()
		helper.LogError(err.Error.Error(), "func:"+funcNow, "script: update status buku")
		return err.Error
	}

	tx.Commit()
	return nil
}

func (s RepoPeminjamanImpl) ListPeminjaman(ctx context.Context) ([]model.ModelPeminjaman, error) {
	funcNow := "RepoPeminjamanImpl.ListPeminjaman"

	var dataList []model.ModelPeminjaman

	// err := s.DB.WithContext(ctx).Table("peminjamans").Select("peminjamans.id, users.name, bukus.nama, peminjamans.status, peminjamans.created, peminjamans.expired").Joins("left join bukus on bukus.id = peminjamans.id_buku").Joins("left join users on users.id = peminjamans.id_user").Scan(&dataList)
	err := s.DB.WithContext(ctx).Raw("SELECT peminjamans.id, users.name, bukus.nama, peminjamans.status, peminjamans.created, peminjamans.expired FROM peminjamans left join bukus on bukus.id = peminjamans.id_buku left join users on users.id=peminjamans.id_user").Scan(&dataList)
	if err.Error != nil {
		helper.LogError(err.Error.Error(), "func:"+funcNow, "script: select buku")
		return dataList, err.Error
	}

	return dataList, nil
}
