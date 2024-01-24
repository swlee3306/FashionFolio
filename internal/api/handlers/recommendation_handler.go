package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecommendationHandler struct{}

func (h *RecommendationHandler) RecommendOutfit(c *gin.Context) {
	// TODO: 날씨 및 계절에 따라 의상을 추천하는 로직 추가
	c.JSON(http.StatusOK, gin.H{"message": "Outfit recommendation successful"})
}
