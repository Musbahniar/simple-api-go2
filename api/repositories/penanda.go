package repositories

import (
	"context"
	"database/sql"
	"simple-api-go2/api/models"
)

type PenandaRepo struct {
	db *sql.DB
}

func NewPenandaRepo(db *sql.DB) models.PenandaRepo {
	return &PenandaRepo{
		db: db,
	}
}

func (g *PenandaRepo) GetAll(ctx context.Context, limit int64, offset int64) ([]*models.Penanda, error) {
	rows, err := g.db.Query(`SELECT c_IdPenanda,c_Penanda FROM t_Penanda LIMIT ?,?`, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Penanda, 0)
	for rows.Next() {
		data := new(models.Penanda)
		err := rows.Scan(
			&data.IdPenanda,
			&data.NamaPenanda,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (g *PenandaRepo) GetOne(ctx context.Context, id string) (*models.Penanda, error) {
	data := &models.Penanda{}
	row := g.db.QueryRow(`SELECT c_IdPenanda,c_Penanda FROM t_Penanda WHERE c_IdPenanda = ? `, id)
	err := row.Scan(
		&data.IdPenanda,
		&data.NamaPenanda,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}
