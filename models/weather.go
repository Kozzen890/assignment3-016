package models

type WeatherData struct {
	ID    uint `gorm:"primaryKey" json:"-"`
	Water int  `json:"water"`
	Wind  int  `json:"wind"`
}
