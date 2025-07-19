package router

import (
	"stock-and-order-management/infrastructure/database"
	"stock-and-order-management/internal/handler"
	"stock-and-order-management/internal/repository"
	"stock-and-order-management/internal/usecase"

	"github.com/gin-gonic/gin"
)

func InitRouter(version *gin.RouterGroup) {
	ProductRouter(version)
	OrderRouter(version)
}

func ProductRouter(version *gin.RouterGroup) {

	repo := repository.NewProductRepository(&database.DB)
	usecase := usecase.NewProductUsecase(repo)
	handler := handler.NewProductHandler(usecase)

	productRouter := version.Group("/products")

	productRouter.POST("", handler.CreateProduct)
	productRouter.GET("", handler.List)
	productRouter.GET("/:id", handler.GetByID)

}

func OrderRouter(version *gin.RouterGroup) {

	repo := repository.NewOrderRepository(&database.DB)
	productRepo := repository.NewProductRepository(&database.DB)
	usecase := usecase.NewOrderUsecase(repo, productRepo)
	handler := handler.NewOrderHandler(usecase)

	orderRouter := version.Group("/orders")

	orderRouter.POST("", handler.CreateOrder)

}
