package database

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitPostgre() (*sql.DB, error) {
	databaseURL := "postgresql://netlifydb_owner:npg_I0zOC6kXVHdR@ep-frosty-shadow-anc6e5l2.c-6.us-east-1.db.netlify.com/netlifydb?channel_binding=require&sslmode=require"
	databaseURL = "postgresql://netlifydb_owner:npg_WeUHn48uqyRB@ep-weathered-cherry-anx8qh59.c-6.us-east-1.db.netlify.com/netlifydb?sslmode=require"
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
