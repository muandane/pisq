package azure

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/charmbracelet/log"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

func Upload(azureContainerName string, backupPath, azureAccountName string, azureAccountKey string) {
	credential, err := azblob.NewSharedKeyCredential(azureAccountName, azureAccountKey)
	if err != nil {
		log.Fatalf("Wrong Credntials: %v", err)
	}
	pipeline := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	URL, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", azureAccountName, azureContainerName))

	containerURL := azblob.NewContainerURL(*URL, pipeline)

	file, err := os.Open(backupPath)
	if err != nil {
		log.Fatalf("File already exists:%v", err)
	}
	defer file.Close()

	blockBlobURL := containerURL.NewBlockBlobURL(backupPath)
	_, err = azblob.UploadFileToBlockBlob(context.Background(), file, blockBlobURL, azblob.UploadToBlockBlobOptions{})
	if err != nil {
		log.Error("Error while uploading file to container in Azure Storage account", err)
	} else {
		log.Info("Upload completed successfully!")
	}
}
