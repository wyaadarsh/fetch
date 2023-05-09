package core

import (
	"database/sql"
	"log"
	"os"
)

var DB_NAME string = "fetch.db"

func get_db_path() string {
	if os.Getenv("HOME") != "" {
		return os.Getenv("HOME") + "/.fetch-config/" + DB_NAME
	} else if os.Getenv("USERPROFILE") != "" {
		return os.Getenv("USERPROFILE") + "/.fetch-config/" + DB_NAME
	} else {
		return "/tmp"
	}
}

func Get_DB() *sql.DB {
	history_db, err := sql.Open("sqlite3", get_db_path()+"/"+DB_NAME)
	if err != nil {
		panic(err)
	}
	return history_db
}

type History struct {
	Id       int
	Url      string
	Success  string
	FilePath string
	FileName string
	FileSize string
	Date     string
}

func CreateSqliteDB() {
	os.MkdirAll(get_db_path(), 0755)
	if _, err := os.Stat(get_db_path() + "/" + DB_NAME); os.IsNotExist(err) {
		os.Create(get_db_path() + "/" + DB_NAME)
	}
	history_db, err := sql.Open("sqlite3", get_db_path()+"/"+DB_NAME)
	if err != nil {
		panic(err)
	}
	defer history_db.Close()
	Initialize_table(history_db)
}

func Initialize_table(db *sql.DB) {
	createHistoryTable := `CREATE TABLE IF NOT EXISTS history (
		id Integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		url TEXT,
		success TEXT,
		file_path TEXT,
		file_name TEXT,
		file_size TEXT,
		date TEXT
	);`
	_, err := db.Exec(createHistoryTable)
	if err != nil {
		log.Println("Error creating history table:", err.Error())
	}
}

func InsertHistory(db *sql.DB, hist History) {
	insertHistory := `INSERT INTO history (url, success, file_path, file_name, file_size, date) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(insertHistory, hist.Url, hist.Success, hist.FilePath, hist.FileName, hist.FileSize, hist.Date)
	if err != nil {
		log.Println("Error inserting history:", err.Error())
	}
}

func GetHistory(db *sql.DB, limit int) []History {
	var history []History
	rows, err := db.Query("SELECT * FROM history ORDER BY date DESC limit ?", limit)
	if err != nil {
		log.Println("Error getting history:", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var hist History
		err = rows.Scan(&hist.Id, &hist.Url, &hist.Success, &hist.FilePath, &hist.FileName, &hist.FileSize, &hist.Date)
		if err != nil {
			log.Println("Error getting history:", err.Error())
		}
		history = append(history, hist)
	}
	return history
}

func DeleteAllHistory(db *sql.DB) error {
	deleteAllHistory := `DELETE FROM history`
	_, err := db.Exec(deleteAllHistory)
	if err != nil {
		log.Println("Error deleting history:", err.Error())
	}
	return err
}
