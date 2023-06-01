package main

import (
	"log"
	"os"

	"github.com/PraveenRaizo/GoEcom/controllers"
	controllers "github.com/PraveenRaizo/GoEcom/database"
	"github.com/PraveenRaizo/GoEcom/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	// above line: ProductData gets from Products collection and UserData gets data from Users collection

	router := gin.New() // creating a router
	router.Use(gin.Logger())

	routes.UserRoutes(router) // passing the router to UserRoutes()
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddtoCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port)) // logIfError(run the server on default port(8000))

}
