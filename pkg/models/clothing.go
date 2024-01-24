package models

// Clothing 모델은 의류 아이템에 대한 정보를 나타냅니다.
type Clothing struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Color    string `json:"color"`
	// 추가적인 필드들을 필요에 따라 정의
}
