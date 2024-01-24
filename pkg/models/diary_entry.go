package models

// DiaryEntry 모델은 스타일 일기에 대한 정보를 나타냅니다.
type DiaryEntry struct {
	ID      uint   `json:"id"`
	Date    string `json:"date"`
	Content string `json:"content"`
	// 추가적인 필드들을 필요에 따라 정의
}
