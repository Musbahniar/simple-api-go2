package repositories

import (
	"context"
	"database/sql"
	"simple-api-go2/api/models"
)

type TahunAjaranRepo struct {
	db *sql.DB
}

func NewTahunAjaranRepo(db *sql.DB) models.TahunAjaran {
	return &TahunAjaranRepo{
		db: db,
	}
}

func (g *TahunAjaranRepo) GetAll(ctx context.Context, limit int64, offset int64) ([]*models.TahunAjaran, error) {
	rows, err := g.db.Query(`SELECT c_TahunAjaran,c_IsDefault,c_awal,c_akhir, 
							FROM t_TahunAjaran WHERE c_Status = 'Aktif' LIMIT ?,?`, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.TahunAjaran, 0)
	for rows.Next() {
		data := new(models.TahunAjaran)
		err := rows.Scan(
			&data.TahunAjaran,
			&data.IsDefault,
			&data.Awal,
			&data.Akhir,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}
