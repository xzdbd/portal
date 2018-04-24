package api

import (
	"net/http"
	"time"

	"net/url"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

type SharedItem struct {
	ID            bson.ObjectId `json:"id" bson:"_id"`
	URL           *url.URL      `json:"url" bson:"url"`
	ViewCount     int32         `json:"viewCount" bson:"viewCount"`
	DownloadCount int32         `json:"downloadCount" bson:"downloadCount"`
	CreationTime  time.Time     `json:"creationTime" bson:"creationTime"`
	ExpireIn      time.Time     `json:"expireIn" bson:"expireIn"`
	Stauts        bool          `json:"status" bson:"status"`
	FileItem      *FileItem     `json:"fileItem" bson:"fileItem"`
}

// curl -X GET http://localhost:8080/v1/sharedItem
func getSharedItems(c *gin.Context) {
	collection := mgoSession.DB("portal").C("sharedItem")
	sharedItems := []FileItem{}
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

// curl -d '{"ViewCount": 3, "DownloadCount":2, "expireIn": "2018-04-24T12:00:00Z", "creationTime": "2018-04-23T16:11:00Z", "status": true}' -H "Content-Type: application/json" -X POST http://localhost:8080/v1/sharedItem
func addSharedItem(c *gin.Context) {
	var sharedItem SharedItem
	if err := c.ShouldBindJSON(&sharedItem); err == nil {
		collection := mgoSession.DB("portal").C("sharedItem")
		sharedItem.ID = bson.NewObjectId()
		err := collection.Insert(sharedItem)
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
