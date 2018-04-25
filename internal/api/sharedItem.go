package api

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/xzdbd/portal/internal/storage"
	"gopkg.in/mgo.v2/bson"
)

var (
	qiniuAuth storage.QiniuAuth
	ak        = os.Getenv("QINIU_ACCESS_KEY")
	sk        = os.Getenv("QINIU_SECRET_KEY")
	domain    = os.Getenv("QINIU_DOMAIN")
	bucket    = os.Getenv("QINIU_BUCKET")
)

type SharedItem struct {
	ID            bson.ObjectId `json:"id" bson:"_id"`
	URL           string        `json:"url" bson:"url"`
	ViewCount     int32         `json:"viewCount" bson:"viewCount"`
	DownloadCount int32         `json:"downloadCount" bson:"downloadCount"`
	CreationTime  time.Time     `json:"creationTime" bson:"creationTime"`
	ExpireIn      time.Time     `json:"expireIn" bson:"expireIn"`
	Stauts        bool          `json:"status" bson:"status"`
	FileItem      *FileItem     `json:"fileItem" bson:"fileItem"`
}

func init() {
	qiniuAuth = storage.QiniuAuth{AccessKey: ak, SecretKey: sk}
}

// curl -X GET http://localhost:8080/v1/sharedItem
func getSharedItems(c *gin.Context) {
	collection := mgoSession.DB("portal").C("sharedItem")
	sharedItems := []SharedItem{}
	err := collection.Find(bson.M{"status": true}).All(&sharedItems)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, APIError{Code: "E501", Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, sharedItems)
	}
}

// curl -X GET http://localhost:8080/v1/sharedItem/5adda3305785f38e4b630026
func getSharedItem(c *gin.Context) {
	id := c.Param("id")
	if bson.IsObjectIdHex(id) {
		collection := mgoSession.DB("portal").C("sharedItem")
		sharedItem := SharedItem{}
		err := collection.FindId(bson.ObjectIdHex(id)).One(&sharedItem)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusInternalServerError, APIError{Code: "E501", Message: err.Error()})
		} else {
			c.JSON(http.StatusOK, sharedItem)
		}
	} else {
		c.JSON(http.StatusBadRequest, APIError{Code: "E400", Message: "Invalid ObjectId"})
	}
}

// curl -d '{"id": "5addb00f3e27b0b20c6f1c2b", "name":"main.go", "bucket":"xzdbd1", "status": true}' -H "Content-Type: application/json" -X POST http://localhost:8080/v1/sharedItem
func addSharedItem(c *gin.Context) {
	var fileItem FileItem
	var sharedItem SharedItem
	if err := c.ShouldBindJSON(&fileItem); err == nil {
		// generate share url
		qiniuDownloader := storage.QiniuDownloader{QiniuAuth: qiniuAuth}
		downloadURL, err := qiniuDownloader.Download(domain, fileItem.Bucket, fileItem.Name, time.Now().Add(time.Second*3600).Unix())
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIError{Code: "E501", Message: err.Error()})
		}
		sharedItem = SharedItem{
			ID:           bson.NewObjectId(),
			URL:          downloadURL,
			CreationTime: time.Now(),
			ExpireIn:     time.Now().Add(time.Second * 3600),
			Stauts:       true,
			FileItem:     &fileItem,
		}
		collection := mgoSession.DB("portal").C("sharedItem")
		err = collection.Insert(sharedItem)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusInternalServerError, APIError{Code: "E501", Message: err.Error()})
		} else {
			c.JSON(http.StatusCreated, sharedItem)
		}
	} else {
		c.JSON(http.StatusBadRequest, APIError{Code: "E400", Message: err.Error()})
	}
}

// curl -d '{"ViewCount": 10, "DownloadCount":2, "expireIn": "2018-04-24T12:00:00Z", "creationTime": "2018-04-23T16:11:00Z", "status": true}' -H "Content-Type: application/json" -X PUT http://localhost:8080/v1/sharedItem/5adda3305785f38e4b630026
func updateSharedItem(c *gin.Context) {
	id := c.Param("id")
	if bson.IsObjectIdHex(id) {
		sharedItem := SharedItem{ID: bson.ObjectIdHex(id)}
		if err := c.ShouldBindJSON(&sharedItem); err == nil {
			collection := mgoSession.DB("portal").C("sharedItem")
			err := collection.UpdateId(bson.ObjectIdHex(id), sharedItem)
			if err != nil {
				log.Error(err)
				c.JSON(http.StatusInternalServerError, APIError{Code: "E501", Message: err.Error()})
			} else {
				c.JSON(http.StatusAccepted, sharedItem)
			}
		} else {
			c.JSON(http.StatusBadRequest, APIError{Code: "E400", Message: err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, APIError{Code: "E400", Message: "Invalid ObjectId"})
	}
}

//curl -X DELETE http://localhost:8080/v1/sharedItem/5adda3305785f38e4b630026
func deleteSharedItem(c *gin.Context) {
	id := c.Param("id")
	if bson.IsObjectIdHex(id) {
		collection := mgoSession.DB("portal").C("fileItem")
		err := collection.UpdateId(bson.ObjectIdHex(id), bson.M{"status": false})
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusInternalServerError, APIError{Code: "E501", Message: err.Error()})
		} else {
			c.JSON(http.StatusAccepted, id)
		}
	} else {
		c.JSON(http.StatusBadRequest, APIError{Code: "E400", Message: "Invalid ObjectId"})
	}
}
