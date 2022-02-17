package repositories

import (
	"context"
	"database/sql"
	"simple-api-go2/api/models"
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
		NamaJenisDiskon: r.NamaJenisDiskon,
	}
	return payload, nil
}

func (g *JenisDiskonRepo) cekInisialJenisDiskon(inisialDiskon string) (string, error) {
	var id int32
	inisialDiskon = strings.TrimSpace(inisialDiskon)
	inisialDiskon = strings.ToLower(inisialDiskon)
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
