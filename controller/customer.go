package controller

import (
	"e-commerce_api/database"
	"e-commerce_api/handler"
	"e-commerce_api/middleware"

	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


// datfar customer
func SignUp(c *gin.Context){
	var customer database.Customer
	var output *database.Customer

	err := c.ShouldBindJSON(&customer)
	if err != nil {
		panic(err)
	}

	// password tidak boleh kosng
	if customer.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password diperlukan",
		})
		return
	}


	err , output = handler.SignUp(database.DbConnection, &customer)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Anda telah terdaftar",
		"username" : output.Uname,
		"pwd" : output.Password,
	})
}


// melihata data customer
func LoginCustomer(c *gin.Context){
	var dataLogin database.Customer
	var dataDB database.Customer
	var order []database.Order_history

	err:= c.ShouldBindJSON(&dataLogin)
	if err != nil{
		panic(err)
	}

	err, istrue, dataDB, order:= handler.LoginCustomer(database.DbConnection, &dataLogin)
	if err != nil {
		panic(err)
	}
	if istrue{
		// add auth redis
		middleware.Login(c, strconv.FormatInt(dataLogin.ID, 10), dataLogin.Uname)
		middleware.SessionUser = dataLogin.ID
		

		c.JSON(http.StatusOK, gin.H{
			"message": "berhasil masuk",
			"id" : dataDB.ID,
			"user" : dataDB.Uname,
			"cash" : dataDB.Cash,
			"order_history" : order,
			})
	

	} else{
	c.JSON(http.StatusForbidden, gin.H{
		"message": "Username / password tidak ditemukan",
	})}
}

// mengubah data customer
func UpdateCustomer(c *gin.Context){
	// authenticate session
	isTrue := middleware.Authentication(strconv.FormatInt(middleware.SessionUser, 10))
	if isTrue == false{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var dataCustomer database.Customer
	id,_ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&dataCustomer)
	if err != nil {
		panic(err)
	}
	dataCustomer.ID = int64(id)

	// password tidak boleh kosng
	if dataCustomer.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password diperlukan",
		})
		return
	}
	

	err = handler.UpdateCustomer(database.DbConnection, &dataCustomer)
	if err != nil {
		panic(err)
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message" : "Success Memperbaharui data",
		"user_name" : dataCustomer.Uname,
		"password" : dataCustomer.Password,
		"cash" : dataCustomer.Cash,
	})
}


// menghapus customer 
func DeleteCustomer(c *gin.Context){
	// authenticate session
	isTrue := middleware.Authentication(strconv.FormatInt(middleware.SessionUser, 10))
	if isTrue == false{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unauthorized",
		})
		return
	}


	var dataCustomer database.Customer

	id,_ := strconv.Atoi(c.Param("id"))

	dataCustomer.ID = int64(id)

	err := handler.DeleteCustomer(database.DbConnection, &dataCustomer)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Success Delete Customer",
	})
}


func Logout(c *gin.Context){

}

