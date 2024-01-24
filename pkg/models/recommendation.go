package models

// Recommendation 모델은 날씨 및 계절에 따른 의상 추천 정보를 나타냅니다.
type Recommendation struct {
	ID      uint   `json:"id"`
	Weather string `json:"weather"`
	Season  string `json:"season"`
	// 추가적인 필드들을 필요에 따라 정의
}
