package controller

import (
	"final-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderInput struct {
	Status 		string 		`json:"status"`
	TotalPrice 	uint 		`json:"total_price"`
	UserID 		uint 		`json:"user_id"`
}

// GetAllOrder godoc
// @Summary Get all Order.(admin only)
// @Description Get a list of Order User.
// @Tags Order
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.Order
// @Router /order [get]
func GetAllOrder(c *gin.Context){
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var order []models.Order
	db.Find(&order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// CreateOrder godoc
// @Summary Create New Order. (must login)
// @Description Creating a new Order User.
// @Tags Order
// @Param Body body OrderInput true "the body to create a new Order"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Order
// @Router /order [post]
func CreateOrder(c *gin.Context){
	// validate input
	var input OrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create Order
	order := models.Order{Status: input.Status, TotalPrice: input.TotalPrice, UserID: input.UserID}

	db:= c.MustGet("db").(*gorm.DB)
	db.Create(&order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}


// GetOrderById godoc
// @Summary Get Order. (must login)
// @Description Get an Order by id.
// @Tags Order
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Order id"
// @Success 200 {object} models.Order
// @Router /order/{id} [get]
func GetOrderByID(c *gin.Context) { // Get model if exist
    var order models.Order

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": order})
}

// GetOrderByUserId godoc
// @Summary Get Order. (must login)
// @Description Get all Order by UserID.
// @Tags Order
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "User id"
// @Success 200 {object} []models.Order
// @Router /user/{id}/order [get]
func GetOrderByUserId(c *gin.Context) { // Get model if exist
    var order []models.Order

    db := c.MustGet("db").(*gorm.DB)

    if err := db.Where("user_id = ?", c.Param("id")).Find(&order).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": order})
}

// UpdateOrder godoc
// @Summary Update Order. (admin only)
// @Description Update Order user.
// @Tags Order
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Order id"
// @Param Body body OrderInput true "the body to update order"
// @Success 200 {object} models.Order
// @Router /order/{id} [patch]
func UpdateOrder(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var order models.Order
    if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input OrderInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Order
    updatedInput.Status = input.Status
    updatedInput.TotalPrice = input.TotalPrice
    updatedInput.UpdatedAt = time.Now()

    db.Model(&order).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": order})
}

// DeleteOrder godoc
// @Summary Delete one Order.(admin only)
// @Description Delete a Order by id.
// @Tags Order
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Order id"
// @Success 200 {object} map[string]boolean
// @Router /order/{id} [delete]
func DeleteOrder(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var order models.Order
    if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&order)

    c.JSON(http.StatusOK, gin.H{"data": true})
}