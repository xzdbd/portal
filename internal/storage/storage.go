package storage

type Reader interface {
	Stat(bucket, key string) (fileItem interface{}, err error)
	StatAll(bucket, prefix string) (fileItems []interface{}, err error)
}

type Uploader interface {
	Upload(file, bucket, key string) (r interface{}, err error)
}

type Downloader interface {
	Download(domain, bucket, key string, deadline int64) (url string, err error)
}

type UploaderDownloader interface {
	Uploader
	Downloader
}
