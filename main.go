package main

import (
	"log"
	"os"
	"github.com/Michaell99/Go-Ecommerce/controllers"
	"github.com/Michaell99/Go-Ecommerce/database"
	"github.com/Michaell99/Go-Ecommerce/middleware"
	"github.com/Michaell99/Go-Ecommerce/routes"
	"github.com/gin-gonic/gin"
)
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.Client, "Products"), database.UserData(database.Client, "Users")
	
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}

