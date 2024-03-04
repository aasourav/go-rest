package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"
// )

type todo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var todos = []todo{
	{
		Id:          "1",
		Title:       "First Tod",
		Description: "First todo descritpion",
	},
	{
		Id:          "2",
		Title:       "Second Tod",
		Description: "Second todo descritpion",
	},
	{
		Id:          "3",
		Title:       "Third Tod",
		Description: "Third todo descritpion",
	},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func getTodo(c *gin.Context) {
	var todoDoc todo
	isFound := false
	id := c.Param("id")
	for _, todo := range todos {
		if todo.Id == id {
			todoDoc = todo
			isFound = true
			break
		}
	}

	if isFound {
		c.IndentedJSON(http.StatusOK, todoDoc)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Todo not found",
		})
	}

}

func createTodo(c *gin.Context) {
	var newTodo todo
	// c.IndentedJSON(http.StatusOK, c.Request.Body)
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	c.IndentedJSON(http.StatusOK, newTodo)
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	isFound := false

	for i, todoDoc := range todos {
		if todoDoc.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			isFound = true
			return
		}
	}

	if !isFound {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Todo id not found",
		})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "Todo successfully deleted",
		})
	}

}

func main() {
	app := gin.Default()
	app.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "api is working",
		})
	})
	app.GET("/", getTodos)
	app.GET("/todo/:id", getTodo)
	app.POST("/create-todo", createTodo)
	app.DELETE("/todo/:id", deleteTodo)

	app.Run()
}
