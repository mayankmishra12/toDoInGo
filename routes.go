package main
import (
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := TodoModel{Title: c.PostForm("title"), Completed: completed,
	Comment :c.PostForm("comment"), CreatedAt :time.Now()}
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "title": todo.Title, "resouceId": todo.ID, "comment": todo.Comment})
   }

   func getAllTodo(c *gin.Context) {
	var todos []TodoModel
	//var todo  []TodoModel
    db.Find(&todos)
    if len(todos) <= 0 {
	 c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
	 return
	}
   c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todos})
    }

    func GetSingleTodo(c *gin.Context) {
	var todo TodoModel
	todoID := c.Param("id")
    db.First(&todo, todoID)
    if todo.ID == 0 {
	 c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
	 return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todo})
    }

    
	func UpdateTodo(c *gin.Context) {
		var todo TodoModel
		todoID := c.Param("id")
	    db.First(&todo, todoID)
	    if todo.ID == 0 {
		 c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		 return
		}
	    db.Model(&todo).Update("title", c.PostForm("title"))
		completed:= c.PostForm("completed")
		db.Model(&todo).Update("completed", completed)
		db.Model(&todo).Update("Comment",c.PostForm("comment"))
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
	    }

	
	
	func DeleteTodo (c *gin.Context) {
		var todo TodoModel
		todoID := c.Param("id")
	    db.First(&todo, todoID)
	    if todo.ID == 0 {
		 c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		 return
		}
	    db.Delete(&todo)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
	    
	}