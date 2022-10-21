package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	"ilmudata/project/database"
	"ilmudata/project/models"
)

// type ProductForm struct {
// 	Email string `form:"email" validate:"required"`
// 	Address string `form:"address" validate:"required"`
// }

type ProductController struct {
	// declare variables
	store *session.Store
	Db    *gorm.DB
}

type CartController struct {
	// declare variables
	Db *gorm.DB
}

func InitProductController(s *session.Store) *ProductController {
	db := database.InitDb()
	// gorm
	db.AutoMigrate(&models.Product{})

	return &ProductController{Db: db}
}

func InitCartController() *CartController {
	db := database.InitDb()
	// gorm
	db.AutoMigrate(&models.Cart{})

	return &CartController{Db: db}
}

// routing
// GET /products
func (controller *ProductController) IndexProduct(c *fiber.Ctx) error {
	// load all products
	var products []models.Product
	err := models.ReadProducts(controller.Db, &products)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("homeuser", fiber.Map{
		"Title":    "Daftar Produk",
		"Products": products,
	})
}

func (controller *ProductController) Home(c *fiber.Ctx) error {
	// load all products
	var products []models.Product
	err := models.ReadProductsUser(controller.Db, &products)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("homeuser", fiber.Map{
		"Title":    "Daftar Produk",
		"Products": products,
	})
}

func (controller *CartController) Cart(c *fiber.Ctx) error {
	// load all products
	var carts []models.Cart
	err := models.ReadCart(controller.Db, &carts)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("cart", fiber.Map{
		"Title": "Daftar Cart",
		"Cart":  carts,
	})
}

// GET /products/create
func (controller *ProductController) AddProduct(c *fiber.Ctx) error {
	return c.Render("addproduct", fiber.Map{
		"Title": "Tambah Produk",
	})
}

// POST /products/create
func (controller *ProductController) AddPostedProduct(c *fiber.Ctx) error {
	//myform := new(models.Product)
	var myform models.Product

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/products")
	}
	// save product
	err := models.CreateProduct(controller.Db, &myform)
	if err != nil {
		return c.Redirect("/products")
	}
	// if succeed
	return c.Redirect("/products")
}

// GET /products/productdetail?id=xxx
func (controller *ProductController) GetDetailProduct(c *fiber.Ctx) error {
	id := c.Query("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("productdetail", fiber.Map{
		"Title":   "Detail Produk",
		"Product": product,
	})
}

// GET /products/detail/xxx
func (controller *ProductController) GetDetailProduct2(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("productdetail", fiber.Map{
		"Title":   "Detail Produk",
		"Product": product,
	})
}

/// GET products/editproduct/xx
func (controller *ProductController) EditlProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("editproduct", fiber.Map{
		"Title":   "Edit Produk",
		"Product": product,
	})
}

/// POST products/editproduct/xx
func (controller *ProductController) EditlPostedProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	var myform models.Product

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/products")
	}
	product.Name = myform.Name
	product.Quantity = myform.Quantity
	product.Price = myform.Price
	product.Image = myform.Image
	// save product
	models.UpdateProduct(controller.Db, &product)

	return c.Redirect("/products")
}

/// GET /products/deleteproduct/xx
func (controller *ProductController) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	models.DeleteProductById(controller.Db, &product, idn)
	return c.Redirect("/products")
}

// POST /products/create
func (controller *CartController) AddItem(c *fiber.Ctx) error {
	//myform := new(models.Product)
	var myform models.Cart

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/cart")
	}
	// save product
	err := models.CreateCart(controller.Db, &myform)
	if err != nil {
		return c.Redirect("/cart")
	}
	// if succeed
	return c.Redirect("/cart")
}

// func UploadFile(c *fiber.Ctx) error {
// 	ctx := context.Background()
// 	bucketName := os.Getenv("MINIO_BUCKET")
// 	file, err := c.FormFile("fileUpload")

// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	// Get Buffer from file
// 	buffer, err := file.Open()

// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}
// 	defer buffer.Close()

// 	// Create minio connection.
// 	minioClient, err := minioUpload.MinioConnection()
// 	if err != nil {
// 		// Return status 500 and minio connection error.
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	objectName := file.Filename
// 	fileBuffer := buffer
// 	contentType := file.Header["Content-Type"][0]
// 	fileSize := file.Size

// 	// Upload the zip file with PutObject
// 	info, err := minioClient.PutObject(ctx, bucketName, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})

// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

// 	return c.JSON(fiber.Map{
// 		"error": false,
// 		"msg":   nil,
// 		"info":  info,
// 	})
// }
