package repositories

import (
	"context"
	"database/sql"
	"simple-api-go2/api/models"
	"strconv"
)

type BidangRepo struct {
	db *sql.DB
}

func NewBidangRepo(db *sql.DB) models.BidangRepo {
	return &BidangRepo{
		db: db,
	}
}

func (g *BidangRepo) GetAll(ctx context.Context, limit int64, offset int64) ([]*models.Bidang, error) {
	rows, err := g.db.Query(`SELECT c_IdBidang,c_NamaBidang,c_Status,c_IdKewilayahan,c_Upline 
							FROM t_GOBidang WHERE c_Status = 'Aktif' LIMIT ?,?`, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Bidang, 0)
	for rows.Next() {
		data := new(models.Bidang)
		err := rows.Scan(
			&data.IdBidang,
			&data.NamaBidang,
			&data.Status,
			&data.IdKewilayahan,
			&data.Upline,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (g *BidangRepo) GetOne(ctx context.Context, id string) (*models.Bidang, error) {
	data := &models.Bidang{}
	bidangId, _ := strconv.ParseInt(id, 10, 64)
	row := g.db.QueryRow(`SELECT c_IdBidang,c_NamaBidang,c_Status,c_IdKewilayahan,c_Upline 
				 FROM t_GOBidang WHERE c_Status = 'Aktif' AND (c_IdBidang = ? OR c_NamaBidang = ?)`, bidangId, id)
	err := row.Scan(
		&data.IdBidang,
		&data.NamaBidang,
		&data.Status,
		&data.IdKewilayahan,
		&data.Upline,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}
