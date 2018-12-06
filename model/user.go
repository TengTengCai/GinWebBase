package model

import (
	"github.com/jinzhu/gorm"
	"time"
	"errors"
	"GinWebBase/utils"
)

// 数据模型结构体
type Users struct {
	gorm.Model

	Name string `gorm:"size:32;"`
	Age  int	`gorm:"default: 18"`
	Birthday *time.Time
}

// 获取用户
func GetOneUserById(id string) (*Users, error){
	user := &Users{}
	err := DB.Raw("Select * from users where `id` = ?", id).Scan(user).Error
	//DB.Where("id = ?", id).Find(user)
	if err != nil {
		utils.Log.Error(err.Error())
		return user, errors.New(err.Error())
	}
	return user, nil
}

// 添加用户
func AddOneUser(user *Users) error {
	err := DB.Create(user).Error
	//DB.Exec("insert into user(`name`, `age`, `birthday`) values (?, ?, ?)", user.Name, user.Age, user.Birthday)
	if err != nil {
		utils.Log.Error(err.Error())
		return errors.New(err.Error())
	}
	return nil

}

// 更新用户
func (user *Users)UpdateUser(name string, age int, birthday time.Time) error {
	user.Name = name
	user.Age = age
	user.Birthday = &birthday
	return DB.Save(user).Error
}

// 删除用户
func (user * Users)DeleteUser() error {
	return DB.Delete(user).Error
}
