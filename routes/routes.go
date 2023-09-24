package routes

import (
	"final-project/controller"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    // set db to gin context
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })

    r.POST("/register", controller.Register)
    r.POST("/login", controller.Login)

    r.GET("/product", controller.GetAllProduct)
    r.GET("/product/:id", controller.GetProductById)

    productMiddlewareRoute := r.Group("/product")
    productMiddlewareRoute.Use(middlewares.JwtAdminMiddleware())
    productMiddlewareRoute.POST("/", controller.CreateProduct)
    productMiddlewareRoute.PATCH("/:id", controller.UpdateProduct)
    productMiddlewareRoute.DELETE("/:id", controller.DeleteProduct)

    r.GET("/categories", controller.GetAllCategory)
    r.GET("/categories/:id", controller.GetCategoryByID)
    r.GET("/categories/:id/product", controller.GetProductByCategoryId)
    
    ratingMiddlewareRoute := r.Group("/categories")
    ratingMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    ratingMiddlewareRoute.POST("/", controller.CreateCategory)
    ratingMiddlewareRoute.PATCH("/:id", controller.UpdateCategory)
    ratingMiddlewareRoute.DELETE("/:id", controller.DeleteCategory)
    
    r.GET("/order", controller.GetAllOrder)
    r.GET("/order/:id", controller.GetOrderByID)

    orderMiddlewareRoute := r.Group("/order")
    orderMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    orderMiddlewareRoute.POST("/", controller.CreateOrder)
    orderMiddlewareRoute.PATCH("/:id", controller.UpdateOrder)
    orderMiddlewareRoute.DELETE("/:id", controller.DeleteOrder)

    r.GET("/order-item", controller.GetAllOrderItem)
    r.GET("/order-item/:id", controller.GetOrderItemById)

    orderItemMiddlewareRoute := r.Group("/order-item")
    orderItemMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    orderItemMiddlewareRoute.POST("/", controller.CreateOrderItem)
    orderItemMiddlewareRoute.PATCH("/:id", controller.UpdateOrderItem)
    orderItemMiddlewareRoute.DELETE("/:id", controller.DeleteOrderItem)
    
    r.GET("/review", controller.GetAllReview)
    r.GET("/review/:id", controller.GetReviewById)

    reviewMiddlewareRoute := r.Group("/review")
    reviewMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    reviewMiddlewareRoute.POST("/", controller.CreateReview)
    reviewMiddlewareRoute.PATCH("/:id", controller.UpdateReview)
    reviewMiddlewareRoute.DELETE("/:id", controller.DeleteReview)


    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}