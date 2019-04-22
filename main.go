package  main

import (
       "github.com/gin-gonic/gin"
       "github.com/jinzhu/gorm"
       _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func main() {
router := gin.Default()
v1 := router.Group("/api/v1/todos")
 {
  v1.POST("/", CreateTodo)
  v1.GET("/", getAllTodo)
  v1.GET("/:id", GetSingleTodo)
  v1.PUT("/:id", UpdateTodo)
  v1.DELETE("/:id", DeleteTodo)
 }
 router.Run()
}

func init() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "sql12288950:9dlIf4bXpU@(sql12.freesqldatabase.com:3306)/sql12288950")
	if err != nil {
	 panic("failed to connect database")
	}
   //Migrate the schema
	db.AutoMigrate(&TodoModel{})
   }