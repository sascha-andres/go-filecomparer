package filedb

import (
	"fmt"
	"time"
)

type (
	// File represents the data stored for one file
	File struct {
		RelativePath string `gorm:"primary_key"`
		Hash         string
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
)

var (
	// ErrDatabaseNotPresent will be returned whenever a crud operation is performed in a non-initialized diretory
	ErrDatabaseNotPresent = fmt.Errorf("Database file does not exist. Please run init")
)

// Save stores file information in
func (file File) Save() error {
	sugar.Debug("Save", "file", file)
	if nil == DB {
		return ErrDatabaseNotPresent
	}
	if "" == file.RelativePath {
		return fmt.Errorf("No path for file given")
	}
	existing, err := Get(file.RelativePath)
	if err != nil {
		return err
	}
	if existing != nil {
		DB.Model(&file).Update("Hash", file.Hash)
	} else {
		DB.Model(&file).Create(&file)
	}
	if DB.Error != nil {
		return DB.Error
	}
	return nil
}

// Get retrieves file information from the database
func Get(path string) (*File, error) {
	sugar.Debug("Get", "path", path)
	if nil == DB {
		return nil, ErrDatabaseNotPresent
	}
	if "" == path {
		return nil, fmt.Errorf("No path for file given")
	}
	var file *File
	DB.Where("RelativePath = ?", path).First(file)
	if DB.Error != nil {
		return nil, DB.Error
	}
	return file, nil
}

// Delete removes a file from the database
func (file File) Delete() error {
	sugar.Debug("Delete", "file", file)
	if nil == DB {
		return ErrDatabaseNotPresent
	}
	if "" == file.RelativePath {
		return fmt.Errorf("No path for file given")
	}
	existing, err := Get(file.RelativePath)
	if err != nil {
		return err
	}
	DB.Delete(existing)
	return DB.Error
}
