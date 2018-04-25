package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/xzdbd/portal/internal/storage"
	"gopkg.in/mgo.v2/bson"
)

// APIError return error
type APIError struct {
	Code    string
	Message string
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

// curl -X GET http://localhost:8080/v1/fileItem
func getFileItems(c *gin.Context) {
	collection := mgoSession.DB("portal").C("fileItem")
	fileItems := []FileItem{}
	err := collection.Find(bson.M{"status": true}).All(&fileItems)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, APIError{Code: "E501", Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, fileItems)
	}
}

// curl -X GET http://localhost:8080/v1/fileItem/5adda3305785f38e4b630026
func getFileItem(c *gin.Context) {
	id := c.Param("id")
	if bson.IsObjectIdHex(id) {
		collection := mgoSession.DB("portal").C("fileItem")
		fileItem := FileItem{}
		err := collection.FindId(bson.ObjectIdHex(id)).One(&fileItem)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusInternalServerError, APIError{Code: "E501", Message: err.Error()})
		} else {
			c.JSON(http.StatusOK, fileItem)
		}
	} else {
		c.JSON(http.StatusBadRequest, APIError{Code: "E400", Message: "Invalid ObjectId"})
	}
}

// curl -d '{"name":"name1", "hash":"123", "fsize": 452, "mimeType":"file", "creationTime": "2018-04-23T16:11:00Z", "status": true}' -H "Content-Type: application/json" -X POST http://localhost:8080/v1/fileItem
func addFileItem(c *gin.Context) {
	var fileItem FileItem

	if err := c.ShouldBindJSON(&fileItem); err == nil {
		collection := mgoSession.DB("portal").C("fileItem")
		fileItem.ID = bson.NewObjectId()
		err := collection.Insert(fileItem)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusInternalServerError, APIError{Code: "E501", Message: err.Error()})
		} else {
			c.JSON(http.StatusCreated, fileItem)
		}
	} else {
		c.JSON(http.StatusBadRequest, APIError{Code: "E400", Message: err.Error()})
	}
}

// curl -d '{"name":"name2", "hash":"123", "fsize": 452, "mimeType":"file", "creationTime": "2018-04-23T16:11:00Z", "status": true}' -H "Content-Type: application/json" -X PUT http://localhost:8080/v1/fileItem/5adda3305785f38e4b630026
func updateFileItem(c *gin.Context) {
	id := c.Param("id")
	if bson.IsObjectIdHex(id) {
		fileItem := FileItem{ID: bson.ObjectIdHex(id)}
		if err := c.ShouldBindJSON(&fileItem); err == nil {
			collection := mgoSession.DB("portal").C("fileItem")
			err := collection.UpdateId(bson.ObjectIdHex(id), fileItem)
			if err != nil {
				log.Error(err)
				c.JSON(http.StatusInternalServerError, APIError{Code: "E501", Message: err.Error()})
			} else {
				c.JSON(http.StatusAccepted, fileItem)
			}
		} else {
			c.JSON(http.StatusBadRequest, APIError{Code: "E400", Message: err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, APIError{Code: "E400", Message: "Invalid ObjectId"})
	}
}

//curl -X DELETE http://localhost:8080/v1/fileItem/5adda3305785f38e4b630026
func deleteFileItem(c *gin.Context) {
	id := c.Param("id")
	if bson.IsObjectIdHex(id) {
		collection := mgoSession.DB("portal").C("fileItem")
		err := collection.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"status": false}})
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

// SyncFileItems is to sync file item info between qiniu object storage and mongodb
func SyncFileItems() {
	qiniuReader := storage.QiniuReader{QiniuAuth: qiniuAuth}
	items, err := qiniuReader.StatAll(bucket, "")
	if err != nil {
		log.Error(err)
	}
	// update or insert objects form object storage
	for _, item := range items {
		if fileItem, ok := item.(storage.FileItem); ok {
			collection := mgoSession.DB("portal").C("fileItem")
			err := collection.Find(bson.M{"bucket": fileItem.Bucket, "name": fileItem.Name, "status": true}).One(&fileItem)
			if err != nil {
				fileItem.ID = bson.NewObjectId()
				log.Info("new object insert:", fileItem.Name)
			}
			_, err = collection.Upsert(bson.M{"bucket": fileItem.Bucket, "name": fileItem.Name, "status": true}, fileItem)
			if err != nil {
				log.Error(err)
			}
		} else {
			log.Error("failed to validate file items")
		}
	}
	// ensure all the contents in db is in sync
	var dbFileItems []FileItem
	collection := mgoSession.DB("portal").C("fileItem")
	err = collection.Find(bson.M{"status": true}).All(&dbFileItems)
	if err != nil {
		log.Error(err)
	}
	for _, dbFileItem := range dbFileItems {
		found := false
		for _, item := range items {
			if fileItem, ok := item.(storage.FileItem); ok {
				if dbFileItem.Bucket == fileItem.Bucket && dbFileItem.Name == fileItem.Name && dbFileItem.Hash == fileItem.Hash {
					found = true
					break
				}
			} else {
				log.Error("failed to validate file items")
			}
		}
		if !found {
			collection.UpdateId(dbFileItem.ID, bson.M{"$set": bson.M{"status": false}})
			log.Info("set the file item status to false:", dbFileItem.Name)
		}

	}
	log.Info("Sync file items done.")
}
