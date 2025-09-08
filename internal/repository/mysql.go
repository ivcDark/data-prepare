package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQL(dsn string) (*MySQLRepository, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MySQLRepository{db: db}, nil
}

type LeagueSeason struct {
	ID   string
	Name string
}

func (r *MySQLRepository) GetAllLeagueSeasons() ([]LeagueSeason, error) {
	rows, err := r.db.Query("SELECT id, name FROM league_seasons")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var leagues []LeagueSeason
	for rows.Next() {
		var ls LeagueSeason
		if err := rows.Scan(&ls.ID, &ls.Name); err != nil {
			return nil, err
		}
		leagues = append(leagues, ls)
	}

	return leagues, nil
}