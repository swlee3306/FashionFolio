package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DiaryHandler struct{}

func (h *DiaryHandler) WriteDiary(c *gin.Context) {
	// TODO: 스타일 일기를 작성하는 로직 추가
	c.JSON(http.StatusOK, gin.H{"message": "Diary entry added successfully"})
}

func (h *DiaryHandler) ReadDiary(c *gin.Context) {
	// TODO: 스타일 일기를 조회하는 로직 추가
	c.JSON(http.StatusOK, gin.H{"message": "Diary entry retrieved successfully"})
}
