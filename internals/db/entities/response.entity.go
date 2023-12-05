package entities

type Response struct {
	ID         int    `gorm:"primaryKey"`
	Url        string `gorm:"unique not null type:TEXT"`
	Result     string `gorm:"type:TEXT"`
	StatusCode int    `gorm:"type:INTEGER"`
}

type Translate struct {
	ID          int    `gorm:"primaryKey"`
	Url         string `gorm:"unique not null type:TEXT"`
	Translation string `gorm:"type:translation"`
}
