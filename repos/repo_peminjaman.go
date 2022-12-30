package repos

import (
	"context"

	"dot.go/entities"
	"dot.go/model"
)

type RepoPeminjaman interface {
	InsertPeminjaman(ctx context.Context, pinjam entities.Peminjaman) error
	ListPeminjaman(ctx context.Context) ([]model.ModelPeminjaman, error)
}
