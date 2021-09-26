package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Pong struct {

}

func NewPong() Pong {
	return Pong{}
}

func (p Pong) Pong(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
