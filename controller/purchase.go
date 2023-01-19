package controller

import (
	"e-commerce_api/database"
	"e-commerce_api/handler"
	"e-commerce_api/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetPurchase(c *gin.Context){
	// authenticate session
	isAuthorized := middleware.Authentication(strconv.FormatInt(middleware.SessionUser, 10))
	if isAuthorized == false{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	customer_id,_ := strconv.Atoi(c.Param("id"))
	item_id,_ :=  strconv.Atoi(c.Param("item_id"))

	var (
		err error
		result string
		dataCustomer database.Customer
		dataItem database.Item
		dataDetail database.Order_history
		dataDetailout *database.Order_history
		isTrue bool
	)
	
	err = c.ShouldBindJSON(&dataDetail) // baca banyak item dibeli
	if err != nil {
		panic(err)
	}
	
	
	dataCustomer.ID = int64(customer_id)
	dataItem.ID = int64(item_id)

	err ,result ,dataDetailout, isTrue= handler.GetPurchase(database.DbConnection, &dataCustomer, &dataItem, &dataDetail)
	if err != nil {
		panic(err)
	}

	if isTrue{
		c.JSON(http.StatusOK, gin.H{
			"message" : result,
			"user_name" : dataCustomer.Uname,
			"item_buyed" : dataItem.Name,
			"number_of_items" : dataDetailout.Number_of_item,
			"bill": dataDetailout.Bill,
			"balance" : dataCustomer.Cash,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : result,
		})
	}
	
	
	
}