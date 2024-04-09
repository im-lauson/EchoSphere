package models

type Cartoon struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func (cartoon *Cartoon) TableName() string {
	return "cartoon_tencent"
}
