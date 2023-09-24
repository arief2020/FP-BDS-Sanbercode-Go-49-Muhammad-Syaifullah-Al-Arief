package controller

import (
	"final-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderItemInput struct {
	Quantity 		uint 		`json:"quantity"`
	SubTotal 		uint 		`json:"sub_total"`
	OrderID 		uint 		`json:"order_id"`
	ProductID 		uint 		`json:"product_id"`
}

// GetAllOrderItem godoc
// @Summary Get all Order Product.
// @Description Get a list of Order Product.
// @Tags OrderItem
// @Produce json
// @Success 200 {object} []models.OrderItem
// @Router /order-item [get]
func GetAllOrderItem(c *gin.Context){
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var orderProduct []models.OrderItem
	db.Find(&orderProduct)

	c.JSON(http.StatusOK, gin.H{"data": orderProduct})
}

// CreateOrderItem godoc
// @Summary Create New Order Product.
// @Description Creating a new Order Product.
// @Tags OrderItem
// @Param Body body OrderItemInput true "the body to create a new Order"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.OrderItem
// @Router /order-item [post]
func CreateOrderItem(c *gin.Context){
	// validate input
	var input OrderItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create Orders
	orderProduct := models.OrderItem{Quantity: input.Quantity, SubTotal: input.SubTotal, OrderID: input.OrderID, ProductID: input.ProductID}

	db:= c.MustGet("db").(*gorm.DB)
	db.Create(&orderProduct)

	c.JSON(http.StatusOK, gin.H{"data": orderProduct})
}

// GetOrderItemById godoc
// @Summary Get Order Product.
// @Description Get an Order Product by id.
// @Tags OrderItem
// @Produce json
// @Param id path string true "OrderItem id"
// @Success 200 {object} models.OrderItem
// @Router /order-item/{id} [get]
func GetOrderItemById(c *gin.Context) { // Get model if exist
    var orderProduct models.OrderItem

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&orderProduct).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": orderProduct})
}


// UpdateOrderItem godoc
// @Summary Update Order Product.
// @Description Update Order Product.
// @Tags OrderItem
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "OrderItem id"
// @Param Body body OrderItemInput true "the body to update orderItem"
// @Success 200 {object} models.OrderItem
// @Router /order-item/{id} [patch]
func UpdateOrderItem(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var orderProduct models.OrderItem
    if err := db.Where("id = ?", c.Param("id")).First(&orderProduct).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input OrderItemInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.OrderItem
    updatedInput.Quantity = input.Quantity
    updatedInput.SubTotal = input.SubTotal
    updatedInput.ProductID = input.ProductID
    updatedInput.OrderID = input.OrderID
    updatedInput.UpdatedAt = time.Now()

    db.Model(&orderProduct).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": orderProduct})
}

// DeleteOrderItem godoc
// @Summary Delete one Order Product.
// @Description Delete a OrderItem by id.
// @Tags OrderItem
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "OrderItem id"
// @Success 200 {object} map[string]boolean
// @Router /order-item/{id} [delete]
func DeleteOrderItem(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var orderProduct models.OrderItem
    if err := db.Where("id = ?", c.Param("id")).First(&orderProduct).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&orderProduct)

    c.JSON(http.StatusOK, gin.H{"data": true})
}

