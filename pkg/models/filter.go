package models

// Filter 모델은 검색 및 필터링에 사용되는 기준을 나타냅니다.
type Filter struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
	// 추가적인 필드들을 필요에 따라 정의
}
