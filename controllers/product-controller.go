package controllers

import (
	"burakyucel/test/entities"
	"burakyucel/test/services"
	"burakyucel/test/validators"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductController interface {
	FindAll() []entities.Product
	Store(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service services.ProductService
}

var validate *validator.Validate

func New(service services.ProductService) ProductController {
	validate = validator.New()
	validate.RegisterValidation("should-not-contain-methylparaben", validators.ValidateIfContainsMethylParaben)

	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entities.Product {
	return c.service.FindAll()
}

func (c *controller) Store(ctx *gin.Context) error {
	var product entities.Product
	err := ctx.ShouldBindJSON(&product)

	if err != nil {
		return err
	}

	err = validate.Struct(product)
	if err != nil {
		return err
	}

	c.service.Store(product)

	return nil
}

func (c *controller) Update(ctx *gin.Context) error {
	var product entities.Product
	err := ctx.ShouldBindJSON(&product)

	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	product.ID = id

	err = validate.Struct(product)
	if err != nil {
		return err
	}

	c.service.Update(product)

	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var product entities.Product

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	product.ID = id
	c.service.Delete(product)

	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	products := c.service.FindAll()

	data := gin.H{
		"title":    "Product Page",
		"products": products,
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}
