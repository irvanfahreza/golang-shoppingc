package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	// "strconv"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"ilmudata/project/database"
	"ilmudata/project/models"
)

type LoginForm struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

type AuthController struct {
	// declare variables
	Db    *gorm.DB
	store *session.Store
}

type RegController struct {
	// declare variables
	Db    *gorm.DB
	store *session.Store
}

// func InitRegController(s *session.Store) *RegController {
// 	db := database.InitDb()
// 	// gorm sync
// 	db.AutoMigrate(&models.Account{})
// 	return &RegController{store: s}
// }

func InitAuthController(s *session.Store) *AuthController {
	db := database.InitDb()
	db.AutoMigrate(&models.Account{})

	return &AuthController{Db: db, store: s}
}

// register
func (controller *AuthController) Register(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title": "Daftar",
	})
}

// post login
func (controller *AuthController) RegisAkun(c *fiber.Ctx) error {
	var acc models.Account

	if err := c.BodyParser(&acc); err != nil {
		return c.SendStatus(400)
	}

	bytes, _ := bcrypt.GenerateFromPassword([]byte(acc.Password), 10)
	sHash := string(bytes)

	acc.Password = sHash

	// save user
	err := models.CreateAkun(controller.Db, &acc)
	if err != nil {
		return c.SendStatus(500)
	}

	errs := models.ReadAkunByUsername(controller.Db, &acc, acc.Username)
	if errs != nil {
		return c.SendStatus(500)
	}

	// if succeed
	return c.Redirect("/login")
}

// get login
func (controller *AuthController) Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

// Post login
func (controller *AuthController) LoginAkun(c *fiber.Ctx) error {
	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}

	var input LoginForm
	var myform models.Account
	// password := models.RegisForm

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/login")
	}
	// save product
	errs := models.ReadAkunByUsername(controller.Db, &myform, myform.Username)
	if errs != nil {
		return c.Redirect("/login")
	}

	comppassw := bcrypt.CompareHashAndPassword([]byte(myform.Password), []byte(input.Password))
	if comppassw != nil {
		sess.Set("username", myform.Username)
		sess.Save()

		return c.Redirect("/products")
	}

	return c.Redirect("/login")
}

func (controller *AuthController) Logout(c *fiber.Ctx) error {

	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Destroy()
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

func (controller *AuthController) Profile(c *fiber.Ctx) error {
	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	val := sess.Get("username")

	return c.JSON(fiber.Map{
		"username": val,
	})
}

func (controller *AuthController) ListAcc(c *fiber.Ctx) error {
	// load all products
	var accounts []models.Account
	err := models.ReadAkun(controller.Db, &accounts)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("regis_forms", fiber.Map{
		"Title":   "Daftar Akun",
		"Account": accounts,
	})
}
