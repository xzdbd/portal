package api

import (
	"github.com/gin-gonic/gin"
)

// GinAPIRouter returns a gin router
func GinAPIRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/fileItem", getFileItems)
		v1.GET("/fileItem/:id", getFileItem)
		v1.POST("/fileItem", addFileItem)
		v1.PUT("/fileItem/:id", updateFileItem)
		v1.DELETE("fileItem/:id", deleteFileItem)

		v1.GET("/sharedItem", getSharedItems)
		v1.GET("/sharedItem/:id", getSharedItem)
		v1.POST("/sharedItem", addSharedItem)
		v1.PUT("/sharedItem/:id", updateSharedItem)
		v1.DELETE("sharedItem/:id", deleteSharedItem)
	}

	return router
}
