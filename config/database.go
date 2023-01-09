package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {

	PORT, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Println("Error during conversion")
	}

	fmt.Printf("Connected to database host: %s and port: %s\n", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     PORT,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
	}
	return &dbConfig
}
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
