package repos

import (
	"context"

	"dot.go/entities"
	"dot.go/helper"
	"gorm.io/gorm"
)

type RepoBukuImpl struct {
	DB *gorm.DB
}

func NewRepoBukuImpl(db *gorm.DB) RepoBuku {
	return &RepoBukuImpl{DB: db}
}

func (s RepoBukuImpl) ListBuku(ctx context.Context) ([]entities.Buku, error) {
	funcNow := "RepoBukuImpl.ListBuku"

	var dataList []entities.Buku

	err := s.DB.WithContext(ctx).Table("bukus").Find(&dataList)
	if err.Error != nil {
		helper.LogError(err.Error.Error(), "func:"+funcNow, "script: select buku")
		return dataList, err.Error
	}

	return dataList, nil
}

func (s RepoBukuImpl) GetBuku(ctx context.Context, id_buku int) (entities.Buku, error) {
	funcNow := "RepoBukuImpl.GetBuku"

	var d entities.Buku

	err := s.DB.WithContext(ctx).Table("bukus").Find(&d, "id = ?", id_buku)
	if err.Error != nil {
		helper.LogError(err.Error.Error(), "func:"+funcNow, "script: select buku by id")
		return d, err.Error
	}

	return d, nil
}

func (s RepoBukuImpl) InsertBuku(ctx context.Context, buku entities.Buku) error {
	funcNow := "RepoBukuImpl.InsertBuku"

	err := s.DB.WithContext(ctx).Table("bukus").Select("Nama", "Status", "CreatedAt").Create(&buku)
	if err.Error != nil {
		helper.LogError(err.Error.Error(), "func:"+funcNow, "script: insert buku")
		return err.Error
	}

	return nil
}

func (s RepoBukuImpl) UpdateBuku(ctx context.Context, buku entities.Buku) error {
	funcNow := "RepoBukuImpl.UpdateBuku"
	err := s.DB.WithContext(ctx).Table("bukus").Where("id = ?", buku.ID).Updates(map[string]interface{}{"nama": buku.Nama, "status": buku.Status})
	if err.Error != nil {
		helper.LogError(err.Error.Error(), "func:"+funcNow, "script: update buku")
		return err.Error
	}

	return nil
}

func (s RepoBukuImpl) HapusBuku(ctx context.Context, buku entities.Buku) error {
	funcNow := "RepoBukuImpl.HapusBuku"

	err := s.DB.WithContext(ctx).Table("bukus").Delete(&buku)
	if err.Error != nil {
		helper.LogError(err.Error.Error(), "func:"+funcNow, "script: delete buku")
		return err.Error
	}

	return nil
}
