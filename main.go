package main 

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"e-commerce_api/database"
	"e-commerce_api/controller"
	"e-commerce_api/middleware"
	"os"
	"fmt"
)



func main(){
	// menyambung ke database
	db, _ := database.ConnectDatabase()
	
	// migrasi tabel 
	database.DbMigrate(db)
	defer db.Close()


	
	// basic auth
	basicAuth := gin.BasicAuth(gin.Accounts{
		"admin" : "12345678",
	})
	

	// fungsi connect redis
	middleware.ConnectRedis()
	fmt.Println("berhasil sambung ke redis")

	
	router := gin.Default()
	
	
	authorized := router.Group("/", basicAuth)



	router.POST("/customer/signup", controller.SignUp) // daftar
	router.GET("/customer/login", controller.LoginCustomer) // login lihat info pribadi
	router.PUT("/customer/:id", controller.UpdateCustomer) // ubah info data customer
	router.DELETE("/customer/:id", controller.DeleteCustomer) // delete customer
	
	authorized.POST("/market", controller.InsertItem) // menambah barang
	router.GET("/market",controller.GetAllItem) // menampilkan semua barang
	authorized.PUT("/market/:item_id",controller.UpdateItem) // mengubah barang di market
	authorized.DELETE("/market/:item_id", controller.DeleteItem) // menghapus item dari market

	router.GET("/market/:item_id/customer/:id", controller.GetPurchase) // membeli barang

	router.GET("/market/category", controller.GetItemByCategory) // menampilkan item berdasakan katogeori
	authorized.POST("/market/category", controller.InsertCategories) // membuat kategry
	authorized.PUT("/market/category/:id", controller.UpdateCategory) // mengupdate category


	router.Run(":"+os.Getenv("PORT"))

}
