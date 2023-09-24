package controller

import (
	"final-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReviewInput struct {
	Rating     uint `json:"Rating"`
	Comment    string `json:"Comment"`
	ProductID uint `json:"products_id"`
	UserID     uint `json:"user_id"`
}

// GetAllReview godoc
// @Summary Get all Review.
// @Description Get a list of Review.
// @Tags Review
// @Produce json
// @Success 200 {object} []models.ReviewProduct
// @Router /review [get]
func GetAllReview(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var review []models.ReviewProduct
	db.Find(&review)

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// CreateReview godoc
// @Summary Create New Review.
// @Description Creating a new Review.
// @Tags Review
// @Param Body body ReviewInput true "the body to create a new Review"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.ReviewProduct
// @Router /review [post]
func CreateReview(c *gin.Context){
	// validate input
	var input ReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create Orders
	review := models.ReviewProduct{Rating: input.Rating, Comment: input.Comment, ProductID: input.ProductID, UserID: input.UserID}

	db:= c.MustGet("db").(*gorm.DB)
	db.Create(&review)

	c.JSON(http.StatusOK, gin.H{"data": review})
}


// GetReviewById godoc
// @Summary Get Review.
// @Description Get a Review by id.
// @Tags Review
// @Produce json
// @Param id path string true "Review id"
// @Success 200 {object} models.ReviewProduct
// @Router /review/{id} [get]
func GetReviewById(c *gin.Context) { // Get model if exist
    var review models.OrderItem

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": review})
}

// UpdateReview godoc
// @Summary Update Review.
// @Description Update Review.
// @Tags Review
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Review id"
// @Param Body body ReviewInput true "the body to update review"
// @Success 200 {object} models.ReviewProduct
// @Router /review/{id} [patch]
func UpdateReview(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var review models.ReviewProduct
    if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input ReviewInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.ReviewProduct
    updatedInput.Rating = input.Rating
    updatedInput.Comment = input.Comment
    updatedInput.ProductID = input.ProductID
    updatedInput.UserID = input.UserID
    updatedInput.UpdatedAt = time.Now()

    db.Model(&review).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": review})
}


// DeleteReview godoc
// @Summary Delete one Order Product.
// @Description Delete a Review by id.
// @Tags Review
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Review id"
// @Success 200 {object} map[string]boolean
// @Router /review/{id} [delete]
func DeleteReview(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var review models.ReviewProduct
    if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&review)

    c.JSON(http.StatusOK, gin.H{"data": true})
}
