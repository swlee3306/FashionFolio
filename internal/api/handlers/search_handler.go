package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchHandler struct{}

func (h *SearchHandler) KeywordSearch(c *gin.Context) {
	// TODO: 특정 키워드로 의상을 검색하는 로직 추가
	c.JSON(http.StatusOK, gin.H{"message": "Keyword search successful"})
}

func (h *SearchHandler) FilterClothing(c *gin.Context) {
	// TODO: 다양한 기준으로 의상을 필터링하는 로직 추가
	c.JSON(http.StatusOK, gin.H{"message": "Clothing filtering successful"})
}
