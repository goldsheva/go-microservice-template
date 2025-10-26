package models

type ExampleModel struct {
	ID    int64  `gorm:"primaryKey" json:"id"`
	Title string `gorm:"type:varchar(255);not null" json:"title"`
}
