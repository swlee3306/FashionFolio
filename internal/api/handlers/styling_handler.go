package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StylingHandler struct{}

func (h *StylingHandler) VirtualSlice(c *gin.Context) {
	// TODO: 의상을 슬라이스하여 가상 스타일링을 하는 로직 추가
	c.JSON(http.StatusOK, gin.H{"message": "Virtual styling successful"})
}
