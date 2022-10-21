package main

import (
	// "fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"

	"ilmudata/project/controllers"
)

func main() {
	// session
	store := session.New()

	// load template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// static
	app.Static("/public", "./public")

	// controllers
	prodController := controllers.InitProductController(store)
	authController := controllers.InitAuthController(store)
	// authAdmController := controllers.InitAuthAdmController(store)
	// regController := controllers.InitRegController(store)
	// cartController := controllers.InitCartController()

	app.Get("/profile", func(c *fiber.Ctx) error {
		sess, _ := store.Get(c)
		val := sess.Get("username")
		if val != nil {
			return c.Next()
		}

		return c.Redirect("/login")

	}, authController.Profile)

	prod := app.Group("/products")
	prod.Get("/", prodController.IndexProduct)
	prod.Get("/create", prodController.AddProduct)
	prod.Post("/create", prodController.AddPostedProduct)
	prod.Get("/productdetail", prodController.GetDetailProduct)
	prod.Get("/detail/:id", prodController.GetDetailProduct2)
	prod.Get("/editproduct/:id", prodController.EditlProduct)
	prod.Post("/editproduct/:id", prodController.EditlPostedProduct)
	prod.Get("/deleteproduct/:id", prodController.DeleteProduct)

	// user
	app.Get("/register", authController.Register)
	app.Post("/register", authController.RegisAkun)
	app.Get("/login", authController.Login)
	app.Post("/login", authController.LoginAkun)
	app.Get("/logout", authController.Logout)

	app.Get("/regis_forms/", authController.ListAcc)

	// app.Get("homeuser", cartController.Home)
	// app.Get("/cart", cartController.Cart) // href="/products/addproduct/{{.Id}}"
	// app.Post("/", cartController.AddItem)

	app.Listen(":3000")
}

//app.Get("/profile",authController.Profile)

// app.Use("/profile", func(c *fiber.Ctx) error {
// 	sess,_ := store.Get(c)
// 	val := sess.Get("username")
// 	if val != nil {
// 		return c.Next()
// 	}

// 	return c.Redirect("/login")

// })
