package storage

import (
	"context"
	"time"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"gopkg.in/mgo.v2/bson"
)

type QiniuAuth struct {
	AccessKey string
	SecretKey string
}

type QiniuReader struct {
	QiniuAuth
}

type QiniuUploader struct {
	QiniuAuth
}

type QiniuDownloader struct {
	QiniuAuth
}

// FileItem
type FileItem struct {
	ID           bson.ObjectId `json:"id" bson:"_id"`
	Bucket       string        `json:"bucket" bson:"bucket"`
	Name         string        `json:"name" bson:"name"`
	Hash         string        `json:"hash" bson:"hash"`
	FSize        int64         `json:"fsize" bson:"fsize"`
	MimeType     string        `json:"mimeType" bson:"mimeType"`
	CreationTime time.Time     `json:"creationTime" bson:"creationTime"`
	Status       bool          `json:"status" bson:"status"`
}

func (q *QiniuReader) Stat(bucket, key string) (fileItem interface{}, err error) {
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	cfg := storage.Config{
		UseHTTPS: false,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	fileInfo, err := bucketManager.Stat(bucket, key)
	if err != nil {
		return
	}

	fileItem = FileItem{
		Bucket:       bucket,
		Name:         key,
		Hash:         fileInfo.Hash,
		FSize:        fileInfo.Fsize,
		MimeType:     fileInfo.MimeType,
		CreationTime: storage.ParsePutTime(fileInfo.PutTime),
		Status:       true,
	}
	return
}

func (q *QiniuReader) StatAll(bucket, prefix string) (fileItems []interface{}, err error) {
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	cfg := storage.Config{
		UseHTTPS: false,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)

	limit := 1000
	delimiter := ""
	marker := ""

	for {
		entries, _, nextMarker, hashNext, err := bucketManager.ListFiles(bucket, prefix, delimiter, marker, limit)
		if err != nil {
			break
		}
		for _, entry := range entries {
			fileItem := FileItem{
				Bucket:       bucket,
				Name:         entry.Key,
				Hash:         entry.Hash,
				FSize:        entry.Fsize,
				MimeType:     entry.MimeType,
				CreationTime: storage.ParsePutTime(entry.PutTime),
				Status:       true,
			}
			fileItems = append(fileItems, fileItem)
		}
		if hashNext {
			marker = nextMarker
		} else {
			break
		}
	}

	return
}

func (q *QiniuDownloader) Download(domain, bucket, key string, deadline int64) (url string, err error) {
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)

	// detect file existence
	cfg := storage.Config{
		UseHTTPS: false,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	_, err = bucketManager.Stat(bucket, key)
	if err != nil {
		return
	}

	url = storage.MakePrivateURL(mac, domain, key, deadline)
	return
}

func (q *QiniuUploader) Upload(file, bucket, key string) (fileItem interface{}, err error) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}

	formUploader := storage.NewFormUploader(&cfg)
	r := storage.PutRet{}

	err = formUploader.PutFile(context.Background(), &r, upToken, key, file, nil)
	if err != nil {
		return
	}
	reader := QiniuReader{QiniuAuth{q.AccessKey, q.SecretKey}}
	fileItem, err = reader.Stat(bucket, key)
	return
}
