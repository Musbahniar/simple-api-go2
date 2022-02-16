package repositories

import (
	"context"
	"database/sql"
	"simple-api-go2/api/models"
	"strconv"
)

type SekolahRepo struct {
	db *sql.DB
}

func NewSekolahRepo(db *sql.DB) models.SekolahRepo {
	return &SekolahRepo{
		db: db,
	}
}

func (g *SekolahRepo) GetAll(ctx context.Context, limit int64, offset int64) ([]*models.Sekolah, error) {
	rows, err := g.db.Query(`SELECT 
								sekolah.c_IdSekolah,sekolah.c_NamaSekolah,sekolah.c_JenjangPendidikan,
								propinsi.c_IdProvinsi,propinsi.c_Provinsi, distrik.c_IdDistrict, distrik.c_District
								FROM t_Sekolah sekolah JOIN t_Idn_District distrik ON distrik.c_IdDistrict = sekolah.c_IdDistrict
								JOIN t_Idn_Provinsi propinsi ON propinsi.c_IdProvinsi = distrik.c_IdProvinsi
								WHERE sekolah.c_Status = 'Aktif' LIMIT ?,?`, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Sekolah, 0)
	for rows.Next() {
		data := new(models.Sekolah)
		data.ProvinsiRes = new(models.GeoData)
		err := rows.Scan(
			&data.IdSekolah,
			&data.NamaSekolah,
			&data.Jenjang,
			&data.ProvinsiRes.IdProvinsi,
			&data.ProvinsiRes.NamaProvinsi,
			&data.ProvinsiRes.IdDistrict,
			&data.ProvinsiRes.NamaDistrict,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (g *SekolahRepo) GetOne(ctx context.Context, id string) (*models.Sekolah, error) {
	data := &models.Sekolah{}
	data.ProvinsiRes = &models.GeoData{}
	sekolahId, _ := strconv.ParseInt(id, 10, 64)
	row := g.db.QueryRow(`SELECT 
							sekolah.c_IdSekolah,sekolah.c_NamaSekolah,sekolah.c_JenjangPendidikan,
							propinsi.c_IdProvinsi,propinsi.c_Provinsi, distrik.c_IdDistrict, distrik.c_District
							FROM t_Sekolah sekolah JOIN t_Idn_District distrik ON distrik.c_IdDistrict = sekolah.c_IdDistrict
							JOIN t_Idn_Provinsi propinsi ON propinsi.c_IdProvinsi = distrik.c_IdProvinsi
							WHERE sekolah.c_Status = 'Aktif' AND c_IdSekolah = ?`, sekolahId)

	err := row.Scan(
		&data.IdSekolah,
		&data.NamaSekolah,
		&data.Jenjang,
		&data.ProvinsiRes.IdProvinsi,
		&data.ProvinsiRes.NamaProvinsi,
		&data.ProvinsiRes.IdDistrict,
		&data.ProvinsiRes.NamaDistrict,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return data, nil
		}
		return nil, err
	}
	return data, nil
}

func (g *SekolahRepo) GetSearchName(ctx context.Context, namaskl string) ([]*models.Sekolah, error) {
	rows, err := g.db.Query(`SELECT 
					sekolah.c_IdSekolah,sekolah.c_NamaSekolah,sekolah.c_JenjangPendidikan,
					propinsi.c_IdProvinsi,propinsi.c_Provinsi, distrik.c_IdDistrict, distrik.c_District
					FROM t_Sekolah sekolah JOIN t_Idn_District distrik ON distrik.c_IdDistrict = sekolah.c_IdDistrict
					JOIN t_Idn_Provinsi propinsi ON propinsi.c_IdProvinsi = distrik.c_IdProvinsi
					WHERE sekolah.c_Status = 'Aktif' AND sekolah.c_NamaSekolah LIKE ?`, namaskl+"%")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Sekolah, 0)
	for rows.Next() {
		data := new(models.Sekolah)
		data.ProvinsiRes = new(models.GeoData)
		err := rows.Scan(
			&data.IdSekolah,
			&data.NamaSekolah,
			&data.Jenjang,
			&data.ProvinsiRes.IdProvinsi,
			&data.ProvinsiRes.NamaProvinsi,
			&data.ProvinsiRes.IdDistrict,
			&data.ProvinsiRes.NamaDistrict,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}
