package configs

import (
	"fmt"
	models "gomvc02/Models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

const Dsn = "sqlserver://usergo:Zz123456789%401@natty-grow.com:1433?database=DBusergo"

var DB *gorm.DB

func GetconnectionString() string {
	// dbHost := "localhost"
	// dbPort := "1433"
	// dbUser := "sa"
	// dbPassword := "Your_password123"
	// dbName := "YourDB"

	//return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	return Dsn
}

func ConnectDatabase() {
	// Implement database connection logic here
	database, err := gorm.Open(sqlserver.Open(GetconnectionString()), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database!")
	} else {
		fmt.Println("Database connected")
	}
	database.AutoMigrate(&models.User{})

	DB = database
}

func CloseDatabase() {
	db, err := DB.DB()
	if err != nil {
		fmt.Println("Failed to close database connection!")
	} else {
		db.Close()
		fmt.Println("Database connection closed")
	}
}

func Testconnect() bool {
	_, err := gorm.Open(sqlserver.Open(GetconnectionString()), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database!")
		return false
	}
	return true
}
