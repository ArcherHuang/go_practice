package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/gojsonq"
)

func main() {

	router := gin.Default()

	router.GET("/role", Get)

	router.GET("/role/:id", GetOne)

	router.POST("/role", Post)

	router.PUT("/role/:id", Put)

	router.DELETE("/role/:id", Delete)

	router.Run(":8088")
}

// 取得全部資料
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, Data)
}

// 取得單一筆資料
func GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	byteArray, err := json.Marshal(Data)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	gq := gojsonq.New().FromString(string(byteArray))
	res := gq.Where("id", "=", id).Get()
	fmt.Println(res)
	c.JSON(http.StatusOK, res)
}

// 新增資料
func Post(c *gin.Context) {
	var r Role
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	Data = append(Data, r)
	c.JSON(http.StatusOK, gin.H{"message": r.Name + " is insert"})
}

type RoleVM struct {
	ID      uint   `json:"id"`      // Key
	Name    string `json:"name"`    // 角色名稱
	Summary string `json:"summary"` // 介紹
}

// 更新資料, 更新角色名稱與介紹
func Put(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var info RoleVM
	err = c.BindJSON(&info)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("info[Name]: %#v\n", info.Name)
	fmt.Printf("info[Summary]: %#v\n", info.Summary)

	for i := range Data {
		if Data[i].ID == uint(id) {
			Data[i].Name = info.Name
			Data[i].Summary = info.Summary
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "ID " + strconv.Itoa(id) + " is updated"})
}

// 刪除資料
func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	for i := range Data {
		if Data[i].ID == uint(id) {
			Data = append(Data[:i], Data[i+1:]...)
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "ID " + strconv.Itoa(id) + " is deleted"})
}
