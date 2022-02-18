package repositories

import (
	"context"
	"database/sql"
	"simple-api-go2/api/models"
	"strconv"
	"strings"
)

type JenisDiskonRepo struct {
	db *sql.DB
}

func NewJenisDiskonRepo(db *sql.DB) models.JenisDiskonRepo {
	return &JenisDiskonRepo{
		db: db,
	}
}

func (g *JenisDiskonRepo) Create(ctx context.Context, r *models.CreateJenisDiskon) (*models.CreateJenisDiskonResponse, error) {
	inisialDiskon, err := g.cekInisialJenisDiskon(r.InisialDiskon)
	if err != nil {
		return nil, err
	}
	//insert record in database
	res, err := g.db.Exec(`INSERT INTO t_JenisDiskon
						(c_InisialDiskon,c_IdBidangApprove,c_NamaDiskon,c_IsRelatif,c_PartDiskon,c_Prioritas,c_Kelompok,c_IsNeedDocument,c_Ikatan,c_Status)
						VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, strings.TrimSpace(inisialDiskon), r.IdBidang, r.NamaJenisDiskon, r.IsRelatif, r.PartDiskon, r.Prioritas, r.Kelompk, r.Dokumen, r.Ikatan, r.Status)
	if err != nil {
		return nil, err
	}
	insertedId, _ := res.LastInsertId()
	payload := &models.CreateJenisDiskonResponse{
		IdJenisDiskon:   insertedId,
		InisialDiskon:   r.InisialDiskon,
		NamaJenisDiskon: r.NamaJenisDiskon,
	}
	return payload, nil
}

func (g *JenisDiskonRepo) Update(ctx context.Context, r *models.JenisDiskon) (bool, error) {
	rows, err := g.db.Exec(`UPDATE t_JenisDiskon SET c_InisialDiskon = ?, c_IdBidangApprove = ?, c_NamaDiskon = ?, c_IsRelatif = ?, c_PartDiskon = ? ,
							c_Prioritas = ?, c_Kelompok = ?, c_IsNeedDocument = ?, c_Ikatan = ?, c_Status = ? WHERE c_IdJenisDiskon = ?`,
		r.InisialDiskon, r.IdBidang, r.NamaJenisDiskon, r.IsRelatif, r.PartDiskon, r.Prioritas, r.Kelompk, r.Dokumen, r.Ikatan, r.Status, r.IdJenisDiskon)
	if err != nil {
		return false, err
	}
	row, _ := rows.RowsAffected()
	if row > 0 {
		return true, nil
	}
	return false, models.ErrNotUpdate
}

func (g *JenisDiskonRepo) Delete(ctx context.Context, id string) (bool, error) {
	jenisDiskonId, _ := strconv.ParseInt(id, 10, 64)
	res, err := g.db.Exec(`DELETE FROM t_JenisDiskon WHERE c_IdJenisDiskon = ?`, jenisDiskonId)
	rows, err := res.RowsAffected()
	if err != nil {
		return false, nil
	}

	if rows > 0 {
		return true, nil
	}

	return false, nil
}

func (g *JenisDiskonRepo) GetAll(ctx context.Context, limit int64, offset int64) ([]*models.JenisDiskon, error) {
	rows, err := g.db.Query(`SELECT c_IdJenisDiskon,c_InisialDiskon,c_IdBidangApprove,c_NamaDiskon,c_IsRelatif,c_PartDiskon,c_Prioritas,c_Kelompok,c_IsNeedDocument,c_Ikatan,c_Status
							FROM t_JenisDiskon LIMIT ?,?`, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.JenisDiskon, 0)
	for rows.Next() {
		data := new(models.JenisDiskon)
		err := rows.Scan(
			&data.IdJenisDiskon,
			&data.InisialDiskon,
			&data.IdBidang,
			&data.NamaJenisDiskon,
			&data.IsRelatif,
			&data.PartDiskon,
			&data.Prioritas,
			&data.Kelompk,
			&data.Dokumen,
			&data.Ikatan,
			&data.Status,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (g *JenisDiskonRepo) GetOne(ctx context.Context, id string) (*models.JenisDiskon, error) {
	data := &models.JenisDiskon{}
	jenisDiskonId, _ := strconv.ParseInt(id, 10, 64)
	row := g.db.QueryRow(`SELECT c_IdJenisDiskon,c_InisialDiskon,c_IdBidangApprove,c_NamaDiskon,c_IsRelatif,c_PartDiskon,c_Prioritas,c_Kelompok,c_IsNeedDocument,c_Ikatan,c_Status
					FROM t_JenisDiskon WHERE c_IdJenisDiskon = ?`, jenisDiskonId)
	err := row.Scan(
		&data.IdJenisDiskon,
		&data.InisialDiskon,
		&data.IdBidang,
		&data.NamaJenisDiskon,
		&data.IsRelatif,
		&data.PartDiskon,
		&data.Prioritas,
		&data.Kelompk,
		&data.Dokumen,
		&data.Ikatan,
		&data.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return data, nil
		}
		return nil, err
	}
	return data, nil
}

func (g *JenisDiskonRepo) cekInisialJenisDiskon(inisialDiskon string) (string, error) {
	var id string
	inisialDiskon = strings.TrimSpace(inisialDiskon)
	// inisialDiskon = strings.ToLower(inisialDiskon)
	inisialDiskon = strings.Replace(inisialDiskon, " ", "-", -1)
	row := g.db.QueryRow(`SELECT c_InisialDiskon FROM t_JenisDiskon WHERE c_InisialDiskon = ?`, inisialDiskon)
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return inisialDiskon, nil
		} else {
			return "", err
		}
	}
	return "", models.ErrAlreadyPresent
}
