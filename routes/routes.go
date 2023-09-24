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
    
    // productMiddlewareRoute := r.Group("/product")
    // productMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    // productMiddlewareRoute.POST("/:id", controller.GetProductById)
    r.GET("/product/:id",middlewares.JwtAuthMiddleware(), controller.GetProductById)

    productAuthMiddlewareRoute := r.Group("/product")
    productAuthMiddlewareRoute.Use(middlewares.JwtAdminMiddleware())
    productAuthMiddlewareRoute.POST("/", controller.CreateProduct)
    productAuthMiddlewareRoute.PATCH("/:id", controller.UpdateProduct)
    productAuthMiddlewareRoute.DELETE("/:id", controller.DeleteProduct)

    r.GET("/categories", controller.GetAllCategory)
    r.GET("/categories/:id", controller.GetCategoryByID)
    r.GET("/categories/:id/product", controller.GetProductByCategoryId)
    
    categoryMiddlewareRoute := r.Group("/categories")
    categoryMiddlewareRoute.Use(middlewares.JwtAdminMiddleware())
    categoryMiddlewareRoute.POST("/", controller.CreateCategory)
    categoryMiddlewareRoute.PATCH("/:id", controller.UpdateCategory)
    categoryMiddlewareRoute.DELETE("/:id", controller.DeleteCategory)
    
    
    orderMiddlewareRoute := r.Group("/order")
    orderMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    orderMiddlewareRoute.GET("/order/:id", controller.GetOrderByID)
    orderMiddlewareRoute.POST("/", controller.CreateOrder)
    
    r.GET("/user/:id/order",middlewares.JwtAuthMiddleware(), controller.GetOrderByUserId)

    orderAdminMiddlewareRoute := r.Group("/order")
    orderAdminMiddlewareRoute.Use(middlewares.JwtAdminMiddleware())
    orderAdminMiddlewareRoute.PATCH("/:id", controller.UpdateOrder)
    orderAdminMiddlewareRoute.DELETE("/:id", controller.DeleteOrder)
    orderAdminMiddlewareRoute.GET("/", controller.GetAllOrder)
    
    // orderitem berdasarkan oder id
    
    orderItemMiddlewareRoute := r.Group("/order-item")
    orderItemMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    orderItemMiddlewareRoute.POST("/", controller.CreateOrderItem)
    orderItemMiddlewareRoute.PATCH("/:id", controller.UpdateOrderItem)
    orderItemMiddlewareRoute.DELETE("/:id", controller.DeleteOrderItem)
    orderItemMiddlewareRoute.GET("/order-item", controller.GetAllOrderItem)
    orderItemMiddlewareRoute.GET("/order-item/:id", controller.GetOrderItemById)
    orderItemMiddlewareRoute.GET("/order/:id/order-item", controller.GetOrderItemByOrderId)
    
    
    // review berdasarkan product id
    r.GET("/product/:id/review", controller.GetReviewByProductId)
    
    reviewMiddlewareRoute := r.Group("/review")
    reviewMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    reviewMiddlewareRoute.GET("/review/:id", controller.GetReviewById)
    reviewMiddlewareRoute.POST("/", controller.CreateReview)
    reviewMiddlewareRoute.PATCH("/:id", controller.UpdateReview)
    reviewMiddlewareRoute.DELETE("/:id", controller.DeleteReview)
    reviewMiddlewareRoute.GET("/review", controller.GetAllReview)


    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}