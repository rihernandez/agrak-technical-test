package main
import (
	"agrak-technical-test/Config"
	"agrak-technical-test/Models"
	"agrak-technical-test/Routes"
	"fmt"
	"github.com/jinzhu/gorm"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	)
	
	var err error

	func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	
	if err != nil {
	 		fmt.Println("Status:", err)
	}
	
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Product{})
	
	r := Routes.SetupRouter()
	
	r.Run()
	}



