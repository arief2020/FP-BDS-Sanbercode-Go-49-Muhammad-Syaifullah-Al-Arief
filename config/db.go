package config

import (
	"final-project/models"
	"final-project/utils"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
  username := utils.Getenv("DATABASE_USERNAME", "root")
  password := utils.Getenv("DATABASE_PASSWORD", "cPdgaPiX1XFg80YEaG39")
  host := utils.Getenv("DATABASE_HOST", "containers-us-west-208.railway.app")
  port := utils.Getenv("DATABASE_PORT", "5839")
  database := utils.Getenv("DATABASE_NAME", "railway")

  dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  // username := "root"
  // password := ""
  // host := "tcp(127.0.0.1:3306)"
  // database := "db_ecomerce_go"

  // dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

  // db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
    panic(err.Error())
  }

  // db.AutoMigrate(&models.Category{}, &models.User{}, &models.Product{}, &models.Order{}, &models.OrderItem{}, &models.ReviewProduct{})
  // db.AutoMigrate(&models.Category{}, &models.User{},&models.Product{}, &models.Order{}, &models.OrderItem{})
  db.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{}, &models.Order{}, &models.OrderItem{}, &models.ReviewProduct{})

  return db
}