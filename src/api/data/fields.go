package models

type Fields struct {
	FieldID    int    `gorm:"primary_key"`
	Name       string `gorm:"type:varchar(255);" json:"name"`
	Type       string `gorm:"type:varchar(255);" json:"type"`
	Length     int    `gorm:"type:integer;" json:"length"`
	PrimaryKey bool   `gorm:"boolean;" json:"primaryKey"`
	Nullable   bool   `gorm:"boolean;" json:"nullable"`
	Precision  int    `gorm:"type:integer;" json:"precision"`
	TableID    int
}

func GetField() Fields {
	var field Fields
	return field
}

func GetFields() []Fields {
	var fields []Fields
	return fields
}
