package main

import (
  "github.com/gin-gonic/gin"
)

type TodoItem struct {
    ID        string  `json:"id"`
    Item      string  `json:"item"`
    Completed bool    `json:"completed"`
}

var TodoItems = []TodoItem{}

func main(){
  router := gin.Default()

  // TODOアイテムの取得
  router.GET("/todos", func(c *gin.Context) {
    c.IndentedJSON(200, TodoItems)
  })

  // 新しいTODOアイテムの作成
    router.POST("/todos", func(c *gin.Context) {
      var newItem TodoItem
      if err := c.BindJSON(&newItem); err != nil {
        return
      }
      TodoItems = append(TodoItems, newItem)
      c.IndentedJSON(200, newItem)
    })

    router.Run(":8080")
}
