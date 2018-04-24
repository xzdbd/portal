package storage

import (
	"os"
	"testing"
	"time"
)

var (
	qiniuAuth QiniuAuth
	testAK    = os.Getenv("QINIU_ACCESS_KEY")
	testSK    = os.Getenv("QINIU_SECRET_KEY")
)

func init() {
	qiniuAuth = QiniuAuth{AccessKey: testAK, SecretKey: testSK}
}

func TestQiniuDownloader(t *testing.T) {
	qiniuDownloader := QiniuDownloader{qiniuAuth}
	url, err := qiniuDownloader.Download("http://7xktk1.com1.z0.glb.clouddn.com", "xzdbd1", "main.go", time.Now().Add(time.Second*3600).Unix())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
}

func TestQiniuUploader(t *testing.T) {
	qiniuUploader := QiniuUploader{qiniuAuth}
	ret, err := qiniuUploader.Upload("./storage_test.go", "xzdbd1", "storage_test.go")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestQiniuStatAll(t *testing.T) {
	qiniuReader := QiniuReader{qiniuAuth}
	ret, err := qiniuReader.StatAll("xzdbd1", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}
