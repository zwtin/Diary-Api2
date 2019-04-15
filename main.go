package main

import (
	"Diary-Api2/controllers"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

func main() {
	r := gin.Default()

	// 1. http://localhost:8080/ へアクセスすると「Hello world」と表示する。
	//	r.GET("/", func(c *gin.Context) {
	//		c.String(200, "Hello world")
	//	})

	//2. http://localhost:8080/hoge へアクセスすると、「fuga」と表示する。
	//	r.GET("/hoge", func(c *gin.Context) {
	//		c.String(200, "fuga")
	//	})

	r.GET("/create/:number/:name", func(c *gin.Context) {
		// Pramを処理する
		num := c.Param("number")
		number, err := strconv.Atoi(num)
		name := c.Param("name")
		if err != nil {
			c.JSON(400, err)
			return
		}
		if number <= 0 {
			c.JSON(400, gin.H{"status": "number should be bigger than 0"})
			return
		}
		// データを処理する
		ctrl := controllers.NewUser()
		result := ctrl.Create(number, name)
		if result == nil || reflect.ValueOf(result).IsNil() {
			c.JSON(404, gin.H{})
			return
		}
		c.JSON(200, result)
	})

	r.GET("/read/:number", func(c *gin.Context) {
		// Pramを処理する
		n := c.Param("number")
		number, err := strconv.Atoi(n)
		if err != nil {
			c.JSON(400, err)
			return
		}
		if number <= 0 {
			c.JSON(400, gin.H{"status": "number should be bigger than 0"})
			return
		}
		// データを処理する
		ctrl := controllers.NewUser()
		result := ctrl.Read(number)
		if result == nil || reflect.ValueOf(result).IsNil() {
			c.JSON(404, gin.H{})
			return
		}
		c.JSON(200, result)
	})

	r.GET("/update/:originnumber/:changednumber/:name", func(c *gin.Context) {
		// Pramを処理する
		originnumber := c.Param("originnumber")
		origin, err := strconv.Atoi(originnumber)
		changednumber := c.Param("changednumber")
		changed, err := strconv.Atoi(changednumber)
		name := c.Param("name")
		if err != nil {
			c.JSON(400, err)
			return
		}
		if origin <= 0 {
			c.JSON(400, gin.H{"status": "origin should be bigger than 0"})
			return
		}
		if changed <= 0 {
			c.JSON(400, gin.H{"status": "changed should be bigger than 0"})
			return
		}
		// データを処理する
		ctrl := controllers.NewUser()
		result := ctrl.Update(origin, changed, name)
		if result == nil || reflect.ValueOf(result).IsNil() {
			c.JSON(404, gin.H{})
			return
		}
		c.JSON(200, result)
	})

	r.GET("/delete/:number", func(c *gin.Context) {
		// Pramを処理する
		n := c.Param("number")
		number, err := strconv.Atoi(n)
		if err != nil {
			c.JSON(400, err)
			return
		}
		if number <= 0 {
			c.JSON(400, gin.H{"status": "number should be bigger than 0"})
			return
		}
		// データを処理する
		ctrl := controllers.NewUser()
		result := ctrl.Delete(number)
		if result == nil || reflect.ValueOf(result).IsNil() {
			c.JSON(404, gin.H{})
			return
		}
		c.JSON(200, result)
	})

	r.Run(":8080")
}
