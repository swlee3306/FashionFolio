// main.go

package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin 엔진 초기화
	router := gin.Default()

	// 엔드포인트 예시
	// router.GET("/health", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// })

	// 서버 포트 설정 (기본값: 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 서버 시작
	log.Printf("Server is running on :%s...\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Server startup failed:", err)
	}
}
