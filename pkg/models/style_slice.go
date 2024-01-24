package models

// StyleSlice 모델은 의상을 가상 슬라이스하여 다양한 스타일을 나타냅니다.
type StyleSlice struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	// 추가적인 필드들을 필요에 따라 정의
}
