package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  )
  


  func Db() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
	return db
  }