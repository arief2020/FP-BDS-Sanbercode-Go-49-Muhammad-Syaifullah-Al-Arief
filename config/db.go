package config

import (
	"belajarGin/models"
	// "belajarGin/utils"
	// "fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
    // username := utils.Getenv("DATABASE_USERNAME", "root")
    // password := utils.Getenv("DATABASE_PASSWORD", "cPdgaPiX1XFg80YEaG39")
    // host := utils.Getenv("DATABASE_HOST", "containers-us-west-208.railway.app")
    // port := utils.Getenv("DATABASE_PORT", "5839")
    // database := utils.Getenv("DATABASE_NAME", "railway")
    // mysql://root:cPdgaPiX1XFg80YEaG39@containers-us-west-208.railway.app:5839/railway

    // dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
    dsn := "mysql://root:cPdgaPiX1XFg80YEaG39@containers-us-west-208.railway.app:5839/railway"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic(err.Error())
    }

    db.AutoMigrate(&models.User{}, &models.Movie{}, &models.AgeRatingCategory{})

    return db
}