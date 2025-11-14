package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
