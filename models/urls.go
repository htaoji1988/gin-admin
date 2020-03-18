package models

type Urls struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"type:varchar(30);"`
	Url  string `gorm:"type:varchar(255);"`
}

func (Urls) TableName() string {
	return "usermanage_urls"
}
