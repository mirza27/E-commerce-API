package controller

import(
	"github.com/gin-gonic/gin"
	"e-commerce_api/database"
	"e-commerce_api/handler"
	"net/http"
	"strconv"
)


// insert categories
func InsertCategories(c *gin.Context){
	var dataCategory database.Category

	err := c.ShouldBindJSON(&dataCategory)
	if err != nil {
		panic(err)
	}

	err  = handler.InsertCategories(database.DbConnection, &dataCategory)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sucess inputing Category",
		"category_name": dataCategory.Name,
	})

}


// filter item by category
func GetItemByCategory(c *gin.Context){
	var (
		result gin.H
		dataCategory database.Category
	)


	parameter := c.Query("category")
	dataCategory.Name = parameter

	err, items := handler.GetItemByCategory(database.DbConnection, &dataCategory)
	if err != nil {
		result = gin.H{
			"result" : err,
		}
	} else {
		result = gin.H{
			"result" : items,
		}
	}

	c.JSON(http.StatusOK, result)
}

// update categoy
func UpdateCategory(c *gin.Context){
	var (
		dataCategory database.Category
		output *database.Category
		message string
		isTrue bool
	)

	id,_ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&dataCategory)
	if err != nil {
		panic(err)
	}

	dataCategory.ID = int64(id)

	err, output, message,isTrue= handler.UpdateCategory(database.DbConnection, &dataCategory)
	if err != nil {
		panic(err)
	}

	if isTrue {
		c.JSON(http.StatusOK, gin.H{
			"result" : message,
			"Name" : output.Name,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"result" : message,
		})
	}


}
