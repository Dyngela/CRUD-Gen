package models

type Table struct {
	TableID int      `gorm:"primary_key"`
	Name    string   `gorm:"type:varchar(255);" json:"name"`
	Fields  []Fields `json:"fields" gorm:"ForeignKey:TableID"`
}

// GetTable -> Method to get a Table object variable/**
func GetTable() Table {
	var table Table
	return table
}

// GetTables -> Method to get an array of Table object variable/**
func GetTables() []Table {
	var table []Table
	return table
}

// FindAllTables -> Find all table and his related fields/**
func FindAllTables() ([]Table, error) {
	var tables = GetTables()
	var fields = GetFields()
	err := dbconn.Find(&tables).Error

	for index, t := range tables {
		err = dbconn.Where("table_id = ?", t.TableID).Find(&fields).Error
		tables[index].Fields = fields
	}

	return tables, err
}

// FindTableById -> Find a table and his related field according to a tableID/**
func FindTableById(id int) (Table, error) {
	var table = GetTable()
	var fields = GetFields()
	err := dbconn.Where("table_id = ?", id).Find(&table).Find(&fields).Error
	table.Fields = fields
	return table, err
}

// CreateTable -> Create table and his eventual fields/**
func CreateTable(table Table) (Table, error) {
	err := dbconn.Create(&table).Error
	return table, err
}

// UpdateTable -> Update a table and his fields/**
func UpdateTable(table Table) (Table, error) {
	err := dbconn.Save(&table).Error
	return table, err
}

// DeleteTableById -> Delete a table and his fields related/**
func DeleteTableById(id int) error {
	err := dbconn.Delete(&Table{}, id).Error
	return err
}
