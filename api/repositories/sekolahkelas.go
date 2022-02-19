package repositories

import (
	"context"
	"database/sql"
	"simple-api-go2/api/models"
	"strconv"
)

type SekolahKelasRepo struct {
	db *sql.DB
}

func NewSekolahKelasRepo(db *sql.DB) models.SekolahKelasRepo {
	return &SekolahKelasRepo{
		db: db,
	}
}

func (g *SekolahKelasRepo) GetAll(ctx context.Context, limit int64, offset int64) ([]*models.SekolahKelas, error) {
	rows, err := g.db.Query(`SELECT c_IdSekolahKelas, CONCAT(c_TingkatKelas,' ',c_KelompokSekolah,' ',c_Jurusan) AS namasekolahkelas,
							c_TingkatKelas, c_Jurusan, c_KelompokSekolah
							FROM t_SekolahKelas WHERE c_Status = 'Aktif' LIMIT ?,?`, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.SekolahKelas, 0)
	for rows.Next() {
		data := new(models.SekolahKelas)
		err := rows.Scan(
			&data.IdSekolahKelas,
			&data.NamaSekolahKelas,
			&data.TingkatKelas,
			&data.Jurusan,
			&data.KelompokKelas,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (g *SekolahKelasRepo) GetOne(ctx context.Context, id string) (*models.SekolahKelas, error) {
	data := &models.SekolahKelas{}
	sekolahkelasiId, _ := strconv.ParseInt(id, 10, 64)
	row := g.db.QueryRow(`SELECT c_IdSekolahKelas, CONCAT(c_TingkatKelas,' ',c_KelompokSekolah,' ',c_Jurusan) AS namasekolahkelas,
					c_TingkatKelas, c_Jurusan, c_KelompokSekolah
					FROM t_SekolahKelas WHERE c_Status = 'Aktif' AND c_IdSekolahKelas = ?`, sekolahkelasiId)

	err := row.Scan(
		&data.IdSekolahKelas,
		&data.NamaSekolahKelas,
		&data.TingkatKelas,
		&data.Jurusan,
		&data.KelompokKelas,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}

func (g *SekolahKelasRepo) GetByKelompok(ctx context.Context, kelompok string) ([]*models.SekolahKelas, error) {
	// data := &models.SekolahKelas{}
	// sekolahkelasiId, _ := strconv.ParseInt(id, 10, 64)
	rows, err := g.db.Query(`SELECT c_IdSekolahKelas, CONCAT(c_TingkatKelas,' ',c_KelompokSekolah,' ',c_Jurusan) AS namasekolahkelas,
					c_TingkatKelas, c_Jurusan, c_KelompokSekolah
					FROM t_SekolahKelas WHERE c_Status = 'Aktif' AND c_KelompokSekolah = ?`, kelompok)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.SekolahKelas, 0)
	for rows.Next() {
		data := new(models.SekolahKelas)
		err := rows.Scan(
			&data.IdSekolahKelas,
			&data.NamaSekolahKelas,
			&data.TingkatKelas,
			&data.Jurusan,
			&data.KelompokKelas,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}
