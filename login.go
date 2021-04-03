package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	router := gin.Default()

	// 匹配的url格式:  http://localhost:2975/getMes/mes
	router.GET("/getMes/:mes", func(c *gin.Context) {
		id := c.DefaultQuery("id", "Guest")
		password := c.Query("password") // 是 c.Request.URL.Query().Get("lastname") 的简写
		//id, password, ok := Check(id, password)
		ok := Check(id, password)
		if !ok {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "验证失败",
			})
			return
		} else{
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "验证成功,正在跳转",

			})
		}
		c.String(http.StatusOK, "Hello %s", id)
	})
	router.Run(":2975")
}


func Check(id string,password string)(bool){
	if id=="LiuC"&&password=="123456"{
		return true
	}else{
		return false
	}
}

