package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kkato/gin-fleamarket/controllers"
	"github.com/kkato/gin-fleamarket/infra"
	"github.com/kkato/gin-fleamarket/middlewares"

	// "github.com/kkato/gin-fleamarket/models"
	"github.com/kkato/gin-fleamarket/repositories"
	"github.com/kkato/gin-fleamarket/services"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	// items := []models.Item{
	// 	{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
	// 	{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
	// 	{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	// }

	// itemRepository := repositories.NewItemMemoryRepository(items)
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	router := gin.Default()
	router.Use(cors.Default())
	itemRouter := router.Group("/items")
	itemRouterWithAuth := router.Group("/items", middlewares.AuthMiddleware(authService))
	authRouter := router.Group("/auth")

	itemRouter.GET("", itemController.FindAll)
	itemRouterWithAuth.GET("/:id", itemController.FindById)
	itemRouterWithAuth.POST("", itemController.Create)
	itemRouterWithAuth.PUT("/:id", itemController.Update)
	itemRouterWithAuth.DELETE("/:id", itemController.Delete)

	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)
	router.Run("localhost:8080")
}
