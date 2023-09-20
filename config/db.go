package config

import (
	"belajarGin/models"
	"belajarGin/utils"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
    username := utils.Getenv("DATABASE_USERNAME", "${{MySQL.MYSQLUSER}}")
    password := utils.Getenv("DATABASE_PASSWORD", "${{MySQL.MYSQLPASSWORD}}")
    host := utils.Getenv("DATABASE_HOST", "${{MySQL.MYSQLHOST}}")
    port := utils.Getenv("DATABASE_PORT", "${{MySQL.MYSQLPORT}}")
    database := utils.Getenv("DATABASE_NAME", "${{MySQL.MYSQLDATABASE}}")

    dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic(err.Error())
    }

    db.AutoMigrate(&models.User{}, &models.Movie{}, &models.AgeRatingCategory{})

    return db
}