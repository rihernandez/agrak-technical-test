package Routes

import (
	"agrak-technical-test/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api/v1")
	{
		grp1.GET("products", Controllers.GetProducts)
		grp1.POST("products", Controllers.CreateProduct)
		grp1.GET("products/:id", Controllers.GetProductByID)
		grp1.PUT("products/:id", Controllers.UpdateProduct)
		grp1.DELETE("products/:id", Controllers.DeleteProduct)
	}
	return r
}
