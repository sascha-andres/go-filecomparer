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

// Save stores file information in
func (file File) Save() error {
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
	if "" == file.RelativePath {
		return fmt.Errorf("No path for file given")
	}
	existing, err := Get(file.RelativePath)
	if err != nil {
		return err
	}
	DB.Delete(existing)
	return nil
}
