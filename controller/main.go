package controller

import (
	"gin/config"
	"gin/entity"
	"gin/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderRequest struct {
	BuyerEmail   string `json:"buyer_email"`
	BuyerAddress string `json:"buyer_address"`
	ProductId    int    `json:"product_id"`
}

// @Summary Get Product
// @Schemes Product
// @Description Get list of all available Products
// @Tags Product
// @Produce json
// @Success 200 {array} entity.Product
// @Router /products [get]
func HandlerGetProduct(ctx *gin.Context) {
	var products []entity.Product
	err := config.DB.Find(&products).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// @Summary Get Order
// @Schemes Order
// @Description Get list of all available Orders
// @Tags Order
// @Produce json
// @Param data body OrderRequest true "Data Order"
// @Success 200 {array} entity.Order
// @Router /orders [get]
func HandlerGetOrder(ctx *gin.Context) {
	var orders []entity.Order
	err := config.DB.Preload("Product").Find(&orders).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

// @Summary Post Order
// @Schemes Order
// @Description Post list of all available Orders
// @Tags Order
// @Produce json
// @Param data body OrderRequest true "Data Order"
// @Success 200 {array} entity.Order
// @Router /orders [post]
func HandlerPostOrder(ctx *gin.Context) {
	// retrieve data
	var orderBody OrderRequest
	err := ctx.ShouldBindJSON(&orderBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var product entity.Product
	result := config.DB.Where("ID = ?", orderBody.ProductId).First(&product)

	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	newOrder := entity.Order{
		BuyerEmail:   orderBody.BuyerEmail,
		BuyerAddress: orderBody.BuyerAddress,
		ProductId:    int(product.ID),
		OrderDate:    time.Now(),
	}

	err = config.DB.Create(&newOrder).Error

	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// trigger mailer
	services.SendMail(newOrder.BuyerEmail, newOrder.BuyerAddress, product.Name)

	ctx.JSON(http.StatusOK, "Successfully Created Data")
}
