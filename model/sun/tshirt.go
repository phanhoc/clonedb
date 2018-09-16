package sun

type TShirt struct {
	ID      uint   `gorm:"primary_key"`
	Url     string `gorm:"not null;type:varchar(2000)"`
	Images  string `gorm:"not null;type:varchar(2000)"`
	Title   string `gorm:"type:varchar(2000)"`
	Desc    string `gorm:"type:varchar(2000)"`
	Time    string `gorm:"type:varchar(2000)"`
	Content string `gorm:"type:varchar(5000)"`
	Money   string `gorm:"type:varchar(1000)"`
	View    uint
}
