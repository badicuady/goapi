package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goapi.com/api/models/book"
)

func GetAll(c *gin.Context) {
	books := book.GetBooks()
	c.IndentedJSON(http.StatusOK, books)
}
