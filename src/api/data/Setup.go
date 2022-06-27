package models

import (
	"github.com/jinzhu/gorm"
)

var dbconn *gorm.DB

// SetDB -> Get the models and create the corresponding table if it does not exist.
// Also, it creates the foreign key /**
func SetDB(db *gorm.DB) {
	dbconn = db
	var table = GetTable()
	var fields = GetField()
	dbconn.AutoMigrate(&table, &fields)
	db.Model(&Fields{}).AddForeignKey("table_id", "tables(table_id)", "CASCADE", "CASCADE")
	//seed(db)
}

// seed -> Init the database with some data /**
func seed(db *gorm.DB) {
	tables := []Table{
		{
			TableID: 1,
			Name:    "table1",
			Fields:  nil,
		},
		{
			TableID: 2,
			Name:    "table2",
			Fields:  nil,
		},
		{
			TableID: 3,
			Name:    "table3",
			Fields:  nil,
		},
	}

	for _, tables := range tables {
		db.Create(&tables)
	}

	fields := []Fields{
		{
			FieldID:    1,
			Name:       "field1",
			Type:       "",
			Length:     0,
			PrimaryKey: false,
			Nullable:   false,
			Precision:  0,
			TableID:    1,
		},
		{
			FieldID:    2,
			Name:       "field2",
			Type:       "",
			Length:     0,
			PrimaryKey: false,
			Nullable:   false,
			Precision:  0,
			TableID:    1,
		},
		{
			FieldID:    3,
			Name:       "field3",
			Type:       "",
			Length:     0,
			PrimaryKey: false,
			Nullable:   false,
			Precision:  0,
			TableID:    2,
		},
		{
			FieldID:    4,
			Name:       "field4",
			Type:       "",
			Length:     0,
			PrimaryKey: false,
			Nullable:   false,
			Precision:  0,
			TableID:    3,
		},
	}

	for _, fields := range fields {
		db.Create(&fields)
	}
}
