package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id       int     `form:"id" json:"id" validate:"required"`
	Name     string  `form:"name" json:"name" validate:"required"`
	Quantity int     `form:"quantity" json:"quantity" validate:"required"`
	Price    float32 `form:"price" json:"price" validate:"required"`
	Image    []byte  `form:"image" validate:"required"`
}

type Cart struct {
	gorm.Model
	Name     string  `form:"name" json:"name" validate:"required"`
	Quantity int     `form:"quantity" json:"quantity" validate:"required"`
	Price    float32 `form:"price" json:"price" validate:"required"`
}

// CRUD
func CreateProduct(db *gorm.DB, newProduct *Product) (err error) {
	err = db.Create(newProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateCart(db *gorm.DB, newCart *Cart) (err error) {
	err = db.Create(newCart).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadProducts(db *gorm.DB, products *[]Product) (err error) {
	err = db.Find(products).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadProductById(db *gorm.DB, product *Product, id int) (err error) {
	err = db.Where("id=?", id).First(product).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateProduct(db *gorm.DB, product *Product) (err error) {
	db.Save(product)
	return nil
}
func DeleteProductById(db *gorm.DB, product *Product, id int) (err error) {
	db.Where("id=?", id).Delete(product)
	return nil
}

func ReadCart(db *gorm.DB, carts *[]Cart) (err error) {
	err = db.Find(carts).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadProductsUser(db *gorm.DB, products *[]Product) (err error) {
	err = db.Find(products).Error
	if err != nil {
		return err
	}
	return nil
}
