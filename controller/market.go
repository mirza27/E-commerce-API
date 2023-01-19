package controller

import (
	"github.com/gin-gonic/gin"
	"e-commerce_api/database"
	"e-commerce_api/handler"
	//"e-commerce_api/middleware"
	"net/http"
	"strconv"
)

// melihat semua item
func GetAllItem(c *gin.Context){
	var show gin.H

	err, list := handler.GetAllItem(database.DbConnection)
	if err != nil {
		show = gin.H{
			"mesaage" : err,
		}
	} else {
		show = gin.H{
			"result" : list,
		}
	}

	c.JSON(http.StatusOK, show)

} 

// membuat item
func InsertItem(c *gin.Context){
	var (
		dataItem database.Item
		output *database.Item
		message string
		isTrue bool
	)

	err := c.ShouldBindJSON(&dataItem)
	if err != nil {
		panic(err)
	}

	err, output, message, isTrue  = handler.InsertItem(database.DbConnection, &dataItem)
	if err != nil {
		panic(err)
	}
	
	if isTrue{
		c.JSON(http.StatusOK, gin.H{
			"message": message,
			"stock" : output.Stock,
			"nama": output.Name,
			"des" : output.Description,
			"price" : output.Price,
			"category" : output.Category_id,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
	}

	
}


// menghapus item
func DeleteItem(c *gin.Context){
	var dataItem database.Item

	id,_ := strconv.Atoi(c.Param("item_id"))

	dataItem.ID = int64(id)

	err := handler.DeleteItem(database.DbConnection, &dataItem)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Success Delete Item",
	})
}


// update item
func UpdateItem(c *gin.Context){
	var (
		dataItem database.Item
		output *database.Item
		message string
		isTrue bool
	)

	id,_ := strconv.Atoi(c.Param("item_id"))

	err := c.ShouldBindJSON(&dataItem)
	if err != nil {
		panic(err)
	}

	dataItem.ID = int64(id)

	err, output, message,isTrue= handler.UpdateItem(database.DbConnection, &dataItem)
	if err != nil {
		panic(err)
	}

	if isTrue {
		c.JSON(http.StatusOK, gin.H{
			"result" : message,
			"Name" : output.Name,
			"Price" : output.Price,
			"Stock" : output.Stock,
			"Description" : output.Description,
			"category_id" : output.Category_id,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"result" : message,
		})
	}


}
