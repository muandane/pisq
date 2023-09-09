package db

import (
	"database/sql"
	"os/exec"

	"github.com/charmbracelet/log"

	_ "github.com/lib/pq"
)

func Backup(username, dbName, dbHost, backupPath string) {
	connStr := "user=" + username + "dbname=" + dbName + "host=" + dbHost + "sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error("Invalid parameters", err)
	}
	defer db.Close()

	cmd := exec.Command("pg_dump", "-U", username, "-d", dbName, "-h", dbHost, "-f", backupPath)
	err = cmd.Run()
	if err != nil {
		log.Error("Error while backing up", err)
	} else {
		log.Info("Upload completed successfully!")
	}
}
