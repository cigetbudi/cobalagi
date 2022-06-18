package models

import "cobalagi/db"

type Users struct {
	Email       string `json:"email" form:"email" gorm:"primaryKey"`
	Nama        string `json:"nama" form:"nama"`
	NoHandphone string `json:"no_handphone" form:"no_handphone"`
	Alamat      string `json:"alamat" form:"alamat"`
	Ktp         string `json:"ktp" form:"ktp"`
}

func (user *Users) CreateUser() error {
	if err := db.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) UpdateUser(email string) error {
	if err := db.DB.Model(&Users{}).Where("email = ?", email).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) DeleteUser() error {
	if err := db.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func GetOneByEmail(email string) (Users, error) {
	var user Users
	result := db.DB.Where("email= ?", email).First(&user)
	return user, result.Error
}
func GetAll(keywords string) ([]Users, error) {
	var users []Users
	result := db.DB.Where("email LIKE ? or nama LIKE ?", "%"+keywords+"%", "%"+keywords+"%").Find(&users)
	return users, result.Error
}
