package controller

import (
	"final-project/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint 	`json:"price"`
	Stock       uint   `json:"stock"`
	CategoryID  uint   `json:"category_id"`
}
func filterMinPrice(data []models.Product, minPrice string) ([]models.Product){
	var newResult []models.Product
	intMP, _ := strconv.Atoi(minPrice)
	for i, val := range data {
		harga := val.Price
		if harga >= uint(intMP) {
			newResult = append(newResult, data[i])
		}
		
	}
	return newResult
}
func filterMaxPrice(data []models.Product, maxPrice string) ([]models.Product){
	var newResult []models.Product
	intMP, _ := strconv.Atoi(maxPrice)
	for i, val := range data {
		harga := val.Price
		if harga <= uint(intMP) {
			newResult = append(newResult, data[i])
		}
		
	}
	return newResult
}
// GetAllProduct godoc
// @Summary Get all product.
// @Description Get a list of Product.
// @Tags Product
// @Produce json
// @Param minPrice query string false "Minimum Price (optional)"
// @Param maxPrice query string false "Maximum Price (optional)"
// @Success 200 {object} []models.Product
// @Router /product [get]
func GetAllProduct(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	db.Find(&products)
    minPrice := c.Query("minPrice")
    maxPrice := c.Query("maxPrice")
    if minPrice != "" {
        products = filterMinPrice(products, minPrice)
    }
    if maxPrice != "" {
        products = filterMaxPrice(products, maxPrice)
    }

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// CreateProduct godoc
// @Summary Create New Product. (admin only)
// @Description Creating a new Product.
// @Tags Product
// @Param Body body productInput true "the body to create a new product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Product
// @Router /product [post]
func CreateProduct(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    // Validate input
    var input productInput
    var category models.Category
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.Where("id = ?", input.CategoryID).First(&category).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "categoryID not found!"})
        return
    }

    // Create Product
    product := models.Product{Name: input.Name, Description: input.Description, Price: input.Price, Stock: input.Stock, CategoryID: input.CategoryID}
    db.Create(&product)

    c.JSON(http.StatusOK, gin.H{"data": product})
}

// GetProductById godoc
// @Summary Get Product. (must login)
// @Description Get a Product by id.
// @Tags Product
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} models.Product
// @Router /product/{id} [get]
func GetProductById(c *gin.Context) { // Get model if exist
    var product models.Product

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": product})
}

// UpdateProduct godoc
// @Summary Update Product. (admin only)
// @Description Update product by id.
// @Tags Product
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "product id"
// @Param Body body productInput true "the body to update an product"
// @Success 200 {object} models.Product
// @Router /product/{id} [patch]
func UpdateProduct(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var product models.Product
    var category models.Category
    if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input productInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.Where("id = ?", input.CategoryID).First(&category).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "CategoryID not found!"})
        return
    }

    var updatedInput models.Product
    updatedInput.Name = input.Name
    updatedInput.Description = input.Description
    updatedInput.Price = input.Price
    updatedInput.Stock = input.Stock
    updatedInput.CategoryID = input.CategoryID
    updatedInput.UpdatedAt = time.Now()

    db.Model(&product).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": product})
}

// DeleteProduct godoc
// @Summary Delete one product. (admin only)
// @Description Delete a product by id.
// @Tags Product
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} map[string]boolean
// @Router /product/{id} [delete]
func DeleteProduct(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var product models.Product
    if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&product)

    c.JSON(http.StatusOK, gin.H{"data": true})
}