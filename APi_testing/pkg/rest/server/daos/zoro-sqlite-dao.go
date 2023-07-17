package daos

import (
	"database/sql"
	"errors"
	"github.com/bheemeshkammak/API/api_testing/pkg/rest/server/daos/clients/sqls"
	"github.com/bheemeshkammak/API/api_testing/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type ZoroDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateZoros(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS zoros(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Dog TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewZoroDao() (*ZoroDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateZoros(sqlClient)
	if err != nil {
		return nil, err
	}
	return &ZoroDao{
		sqlClient,
	}, nil
}

func (zoroDao *ZoroDao) CreateZoro(m *models.Zoro) (*models.Zoro, error) {
	insertQuery := "INSERT INTO zoros(Dog)values(?)"
	res, err := zoroDao.sqlClient.DB.Exec(insertQuery, m.Dog)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("zoro created")
	return m, nil
}

func (zoroDao *ZoroDao) UpdateZoro(id int64, m *models.Zoro) (*models.Zoro, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	zoro, err := zoroDao.GetZoro(id)
	if err != nil {
		return nil, err
	}
	if zoro == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE zoros SET Dog = ? WHERE Id = ?"
	res, err := zoroDao.sqlClient.DB.Exec(updateQuery, m.Dog, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("zoro updated")
	return m, nil
}

func (zoroDao *ZoroDao) DeleteZoro(id int64) error {
	deleteQuery := "DELETE FROM zoros WHERE Id = ?"
	res, err := zoroDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("zoro deleted")
	return nil
}

func (zoroDao *ZoroDao) ListZoros() ([]*models.Zoro, error) {
	selectQuery := "SELECT * FROM zoros"
	rows, err := zoroDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var zoros []*models.Zoro
	for rows.Next() {
		m := models.Zoro{}
		if err = rows.Scan(&m.Id, &m.Dog); err != nil {
			return nil, err
		}
		zoros = append(zoros, &m)
	}
	if zoros == nil {
		zoros = []*models.Zoro{}
	}

	log.Debugf("zoro listed")
	return zoros, nil
}

func (zoroDao *ZoroDao) GetZoro(id int64) (*models.Zoro, error) {
	selectQuery := "SELECT * FROM zoros WHERE Id = ?"
	row := zoroDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Zoro{}
	if err := row.Scan(&m.Id, &m.Dog); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("zoro retrieved")
	return &m, nil
}
