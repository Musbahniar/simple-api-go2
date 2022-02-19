package repositories

import (
	"context"
	"database/sql"
	"simple-api-go2/api/models"
	"strconv"
)

type ProvinsiRepo struct {
	db *sql.DB
}

func NewProvinsiRepo(db *sql.DB) models.ProvinsiRepo {
	return &ProvinsiRepo{
		db: db,
	}
}

func (g *ProvinsiRepo) GetAll(ctx context.Context, limit int64, offset int64) ([]*models.Provinsi, error) {
	rows, err := g.db.Query(`SELECT c_IdProvinsi,c_Provinsi FROM t_Idn_Provinsi WHERE c_Status = 'Aktif' LIMIT ?,?`, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Provinsi, 0)
	for rows.Next() {
		data := new(models.Provinsi)
		err := rows.Scan(
			&data.IdProvinsi,
			&data.NamaProvinsi,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (g *ProvinsiRepo) GetOne(ctx context.Context, id string) (*models.Provinsi, error) {
	data := &models.Provinsi{}
	provinsiId, _ := strconv.ParseInt(id, 10, 64)
	row := g.db.QueryRow(`SELECT c_IdProvinsi,c_Provinsi FROM t_Idn_Provinsi WHERE c_Status = 'Aktif' AND c_IdProvinsi = ?`, provinsiId)
	err := row.Scan(
		&data.IdProvinsi,
		&data.NamaProvinsi,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}
