package repositories

import (
	"context"
	"database/sql"
	"simple-api-go2/api/models"
	"strconv"
)

type DistrikRepo struct {
	db *sql.DB
}

func NewDistrikRepo(db *sql.DB) models.DistrikRepo {
	return &DistrikRepo{
		db: db,
	}
}

func (g *DistrikRepo) GetAll(ctx context.Context, limit int64, offset int64) ([]*models.Distrik, error) {
	rows, err := g.db.Query(`SELECT d.c_IdDistrict,d.c_District,d.c_PhoneCode,p.c_IdProvinsi, p.c_Provinsi
								FROM t_Idn_District d JOIN t_Idn_Provinsi p ON p.c_IdProvinsi = d.c_IdProvinsi
								WHERE d.c_Status = 'Aktif' LIMIT ?,?`, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Distrik, 0)
	for rows.Next() {
		data := new(models.Distrik)
		data.ProvinsiRes = new(models.ProvinsiData)
		err := rows.Scan(
			&data.IdDistrik,
			&data.NamaDistrik,
			&data.PhoneCode,
			&data.ProvinsiRes.IdProvinsi,
			&data.ProvinsiRes.NamaProvinsi,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (g *DistrikRepo) GetOne(ctx context.Context, id string) (*models.Distrik, error) {
	data := &models.Distrik{}
	data.ProvinsiRes = &models.ProvinsiData{}
	distrikId, _ := strconv.ParseInt(id, 10, 64)
	row := g.db.QueryRow(`SELECT d.c_IdDistrict,d.c_District,d.c_PhoneCode,p.c_IdProvinsi, p.c_Provinsi
					FROM t_Idn_District d JOIN t_Idn_Provinsi p ON p.c_IdProvinsi = d.c_IdProvinsi
					WHERE d.c_Status = 'Aktif' AND c_IdDistrict = ?`, distrikId)
	err := row.Scan(
		&data.IdDistrik,
		&data.NamaDistrik,
		&data.PhoneCode,
		&data.ProvinsiRes.IdProvinsi,
		&data.ProvinsiRes.NamaProvinsi,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return data, nil
		}
		return nil, err
	}
	return data, nil
}

func (g *DistrikRepo) GetByIdProvinsi(ctx context.Context, id string) ([]*models.Distrik, error) {
	provinsiId, _ := strconv.ParseInt(id, 10, 64)
	rows, err := g.db.Query(`SELECT d.c_IdDistrict,d.c_District,d.c_PhoneCode,p.c_IdProvinsi, p.c_Provinsi
								FROM t_Idn_District d JOIN t_Idn_Provinsi p ON p.c_IdProvinsi = d.c_IdProvinsi
								WHERE d.c_Status = 'Aktif' AND p.c_IdProvinsi = ?`, provinsiId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Distrik, 0)
	for rows.Next() {
		data := new(models.Distrik)
		data.ProvinsiRes = new(models.ProvinsiData)
		err := rows.Scan(
			&data.IdDistrik,
			&data.NamaDistrik,
			&data.PhoneCode,
			&data.ProvinsiRes.IdProvinsi,
			&data.ProvinsiRes.NamaProvinsi,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}
