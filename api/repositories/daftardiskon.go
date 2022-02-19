package repositories

import (
	"context"
	"database/sql"
	"simple-api-go2/api/models"
	"strconv"
)

type DaftarDiskonRepo struct {
	db *sql.DB
}

func NewDaftarDiskonRepo(db *sql.DB) models.DaftarDiskonRepo {
	return &DaftarDiskonRepo{
		db: db,
	}
}

func (g *DaftarDiskonRepo) GetByTahunAjaran(ctx context.Context, tp string) ([]*models.DaftarDiskon, error) {
	rows, err := g.db.Query(`SELECT 
						dd.c_IdDiskon,dd.c_Nilai,dd.c_TanggalStart,dd.c_TanggalAkhir,dd.c_TahunAjaran,
						jd.c_IdJenisDiskon,jd.c_InisialDiskon,jd.c_NamaDiskon,jd.c_IdBidangApprove,jd.c_IsRelatif,
						jd.c_PartDiskon,jd.c_Prioritas,jd.c_Kelompok,jd.c_IsNeedDocument,jd.c_Ikatan,jd.c_Status
						FROM t_DaftarDiskon dd JOIN t_JenisDiskon jd ON jd.c_IdJenisDiskon = dd.c_IdJenisDiskon
						WHERE dd.c_TahunAjaran = ?`, tp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.DaftarDiskon, 0)
	for rows.Next() {
		data := new(models.DaftarDiskon)
		data.JenisDiskonRes = new(models.JenisDiskonData)
		err := rows.Scan(
			&data.IdDaftarDiskon,
			&data.NilaiDiskon,
			&data.TanggalStart,
			&data.TanggalAkhir,
			&data.TahunAjaran,
			&data.JenisDiskonRes.IdJenisDiskon,
			&data.JenisDiskonRes.InisialDiskon,
			&data.JenisDiskonRes.NamaJenisDiskon,
			&data.JenisDiskonRes.IdBidang,
			&data.JenisDiskonRes.IsRelatif,
			&data.JenisDiskonRes.PartDiskon,
			&data.JenisDiskonRes.Prioritas,
			&data.JenisDiskonRes.Kelompk,
			&data.JenisDiskonRes.Dokumen,
			&data.JenisDiskonRes.Ikatan,
			&data.JenisDiskonRes.Status,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (g *DaftarDiskonRepo) GetOne(ctx context.Context, id string) (*models.DaftarDiskon, error) {
	data := &models.DaftarDiskon{}
	data.JenisDiskonRes = &models.JenisDiskonData{}
	daftarDiskonId, _ := strconv.ParseInt(id, 10, 64)
	row := g.db.QueryRow(`SELECT 
				dd.c_IdDiskon,dd.c_Nilai,dd.c_TanggalStart,dd.c_TanggalAkhir,dd.c_TahunAjaran,
				jd.c_IdJenisDiskon,jd.c_InisialDiskon,jd.c_NamaDiskon,jd.c_IdBidangApprove,jd.c_IsRelatif,
				jd.c_PartDiskon,jd.c_Prioritas,jd.c_Kelompok,jd.c_IsNeedDocument,jd.c_Ikatan,jd.c_Status
				FROM t_DaftarDiskon dd JOIN t_JenisDiskon jd ON jd.c_IdJenisDiskon = dd.c_IdJenisDiskon
				WHERE dd.c_IdDiskon = ?`, daftarDiskonId)
	err := row.Scan(
		&data.IdDaftarDiskon,
		&data.NilaiDiskon,
		&data.TanggalStart,
		&data.TanggalAkhir,
		&data.TahunAjaran,
		&data.JenisDiskonRes.IdJenisDiskon,
		&data.JenisDiskonRes.InisialDiskon,
		&data.JenisDiskonRes.NamaJenisDiskon,
		&data.JenisDiskonRes.IdBidang,
		&data.JenisDiskonRes.IsRelatif,
		&data.JenisDiskonRes.PartDiskon,
		&data.JenisDiskonRes.Prioritas,
		&data.JenisDiskonRes.Kelompk,
		&data.JenisDiskonRes.Dokumen,
		&data.JenisDiskonRes.Ikatan,
		&data.JenisDiskonRes.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}
