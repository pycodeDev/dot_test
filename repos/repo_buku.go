package repos

import (
	"context"

	"dot.go/entities"
)

type RepoBuku interface {
	ListBuku(ctx context.Context) ([]entities.Buku, error)
	GetBuku(ctx context.Context, id_buku int) (entities.Buku, error)
	InsertBuku(ctx context.Context, buku entities.Buku) error
	UpdateBuku(ctx context.Context, buku entities.Buku) error
	HapusBuku(ctx context.Context, buku entities.Buku) error
}
