package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type handlerError struct {
	Message string    `json:"message"`
	Status  int       `json:"status"`
	Now     time.Time `json:"now"`
}

func AergiaResponseOk(c *gin.Context, b any) {
	if b == nil {
		c.AbortWithStatus(http.StatusOK)
	} else {
		c.JSON(http.StatusOK, b)
		c.Abort()
	}
	return
}

func AergiaResponseForbidden(c *gin.Context, err error) {
	status := http.StatusForbidden
	c.Status(status)
	c.JSON(status, handlerError{
		Message: err.Error(),
		Status:  status,
		Now:     time.Now(),
	})
	c.Abort()
	return
}

func AergiaResponseUnauthorized(c *gin.Context, err error) {
	status := http.StatusUnauthorized
	c.Status(status)
	c.JSON(status, handlerError{
		Message: err.Error(),
		Status:  status,
		Now:     time.Now(),
	})
	c.Abort()
	return
}

func AergiaResponseStatusBadRequest(c *gin.Context, err error) {
	status := http.StatusBadRequest
	c.Status(status)
	c.JSON(status, handlerError{
		Message: err.Error(),
		Status:  status,
		Now:     time.Now(),
	})
	c.Abort()
	return
}
