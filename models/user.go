package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Id       uint   `gorm:"primaryKey; not null"`
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
	Email    string `form:"email" json:"email" validate:"required"`
}

func CreateAkun(db *gorm.DB, newAcc *Account) (err error) {
	err = db.Create(newAcc).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadAkun(db *gorm.DB, account *[]Account) (err error) {
	err = db.Find(account).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadAkunByUsername(db *gorm.DB, account *Account, username string) (err error) {
	err = db.Where("username=?", username).First(account).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateAkun(db *gorm.DB, account *Account) (err error) {
	db.Save(account)

	return nil
}
func DeleteAkunById(db *gorm.DB, account *Account, id int) (err error) {
	db.Where("id=?", id).Delete(account)
	return nil
}
