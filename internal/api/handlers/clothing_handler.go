package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClothingHandler struct{}

func (h *ClothingHandler) AddClothing(c *gin.Context) {
	// TODO: 새로운 의류를 데이터베이스에 추가하는 로직 추가
	c.JSON(http.StatusOK, gin.H{"message": "Clothing added successfully"})
}

func (h *ClothingHandler) EditClothing(c *gin.Context) {
	// TODO: 기존 의류를 데이터베이스에서 편집하는 로직 추가
	c.JSON(http.StatusOK, gin.H{"message": "Clothing edited successfully"})
}

func (h *ClothingHandler) DeleteClothing(c *gin.Context) {
	// TODO: 의류를 데이터베이스에서 삭제하는 로직 추가
	c.JSON(http.StatusOK, gin.H{"message": "Clothing deleted successfully"})
}
