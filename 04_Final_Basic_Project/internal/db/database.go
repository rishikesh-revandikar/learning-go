package db

import (
	"database/sql"
	"log/slog"
	_ "modernc.org/sqlite"
)


type Student struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Grade string `json:"grade"`
}

var DB *sql.DB

func InitDB(){
	var err error

	DB, err = sql.Open("sqlite","data.db")

	if err != nil{
		slog.Error("Database connection failed","error",err)
	}

	query := `CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		grade TEXT
	);`

	_, err = DB.Exec(query)
	if err != nil {
		slog.Error("Table creation failed", "error", err)
	}

	slog.Info("Database initialized successfully")

}