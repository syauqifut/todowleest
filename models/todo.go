package models

type Todo struct {
	Id    int64  `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"type:varchar(255)"`
	Desc  string `json:"desc" gorm:"type:varchar(255)"`
}
