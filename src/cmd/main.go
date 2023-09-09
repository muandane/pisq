package main

import (
	"flag"
	"fmt"
	"os"

	"pisq/pkg/azure"
	"pisq/pkg/db"

	"github.com/charmbracelet/log"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flag] [value]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}
	username := flag.String("u", "", "Database username")
	dbName := flag.String("d", "", "Database name")
	dbHost := flag.String("h", "", "Database Host")
	backupPath := flag.String("p", "", "Backup path")
	azureAccountName := flag.String("n", "", "Azure Account Name")
	azureAccountKey := flag.String("k", "", "Azure Account Key")
	azureContainerName := flag.String("c", "", "Azure Blob Storage Container Name")
	azureUpload := flag.Bool("a", false, "Upload to Azure Blob Storage")
	flag.Parse()

	if *azureUpload {
		db.Backup(*username, *dbName, *dbHost, *backupPath)
		azure.Upload(*azureContainerName, *backupPath, *azureAccountName, *azureAccountKey)
	} else {
		db.Backup(*username, *dbName, *dbHost, *backupPath)
		log.Warn("Not uploading to Azure!", "azureUpload", false)
	}
}
