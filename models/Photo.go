package models

type Photo struct {
	ID    int64  `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Data  []byte `json:"-"`
	URL   string `json:"url"`
}
