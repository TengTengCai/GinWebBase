package api

import (
	"github.com/gin-gonic/gin"
	"GinWebBase/utils"
	"net/http"
	"GinWebBase/model"
	"fmt"
	"strconv"
	"time"
)

func Ping(c *gin.Context) {
	utils.Log.Info("ping is come, pong will back")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// 获取用户信息
func GetUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	user, err := model.GetOneUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"user": user,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"user": user,
	})
}

// 添加用户
func AddUser(c *gin.Context)  {
	name := c.PostForm("name")
	age, err := strconv.Atoi(c.PostForm("age"))
	birthday := c.PostForm("birthday")
	user := new(model.Users)
	user.Name = name
	user.Age = age
	bTime, err := time.ParseInLocation("2006-01-02 15:04:05", birthday, time.Local)
	user.Birthday = &bTime

	model.AddOneUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"user": user,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"user": user,
	})
}

// 更新用户信息
func UpdateUser(c *gin.Context)  {
	id := c.Param("id")
	name := c.PostForm("name")
	age, err := strconv.Atoi(c.PostForm("age"))
	birthday := c.PostForm("birthday")
	user, err := model.GetOneUserById(id)
	bTime, err := time.ParseInLocation("2006-01-02 15:04:05", birthday, time.Local)
	err = user.UpdateUser(name, age, bTime)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
			"user": user,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"user": user,
	})
}

// 删除用户
func DeleteUser(c *gin.Context)  {
	id := c.Param("id")
	user, err := model.GetOneUserById(id)
	err = user.DeleteUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}