package firebase

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/storage"
	"google.golang.org/api/option"
)

var App *firebase.App
var StorageClient *storage.Client

func InitFirebase() error {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./firestorage-45220-firebase-adminsdk-eiqhg-1f718a9e75.json") // JSON dosyasının yolu

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return err
	}
	App = app

	storageClient, err := App.Storage(ctx)
	if err != nil {
		return err
	}
	StorageClient = storageClient
	return nil
}

func UploadImage(file io.Reader, fileName string) (string, error) {
	if StorageClient == nil {
		return "", fmt.Errorf("firebase storage client is not initialized")
	}

	ctx := context.Background()
	bucket, err := StorageClient.Bucket("firestorage-45220.appspot.com")
	if err != nil {
		return "", fmt.Errorf("failed to get bucket: %v", err)
	}

	uniqueFileName := fileName
	for {
		obj := bucket.Object(filepath.Join("uploads", uniqueFileName))
		_, err := obj.Attrs(ctx)
		if err == nil {
			uniqueFileName = fmt.Sprintf("%s_%d", fileName, time.Now().UnixNano())
		} else if err != nil && !isNotFoundError(err) {
			return "", fmt.Errorf("failed to check if file exists: %v", err)
		} else {
			break
		}
	}

	filePath := filepath.Join("uploads", uniqueFileName)
	wc := bucket.Object(filePath).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("failed to copy file to writer: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	encodedFilePath := url.QueryEscape(filePath)
	url := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media", "firestorage-45220.appspot.com", encodedFilePath)
	return url, nil
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "object doesn't exist")
}
