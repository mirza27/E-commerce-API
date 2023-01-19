package handler

import (
	"e-commerce_api/database"
	"database/sql"
)



// pembelian
func GetPurchase(db *sql.DB, dataCustomer *database.Customer, dataItem *database.Item, dataDetail *database.Order_history) (error, string, *database.Order_history, bool){
	
	// baca tabel item
	sql := "SELECT name, stock, price FROM item WHERE id = $1" 
	err := db.QueryRow(sql, dataItem.ID).Scan(&dataItem.Name, &dataItem.Stock, &dataItem.Price)
	if err != nil {
		panic(err)
	}

	// baca tabel customer
	sql = "SELECT uname, cash FROM customer WHERE id = $1"
	err = db.QueryRow(sql,dataCustomer.ID).Scan(&dataCustomer.Uname, &dataCustomer.Cash)
	if err != nil {
		panic(err)
	}

	if dataDetail.Number_of_item <= 0 {
		result := "jumlah barang dibeli tidak boleh < 0"
		return err, result, dataDetail, false
	}

	// total harga
	dataDetail.Bill = dataDetail.Number_of_item* dataItem.Price


	// cek saldo customer
	if dataCustomer.Cash < dataDetail.Bill {
		result := "Maaf saldo anda tidak cukup"
		return err, result, dataDetail, false
	}

	// cek stock item
	if dataItem.Stock <= 0 || dataItem.Stock < dataDetail.Number_of_item{
		result := "Maaf, stock habis atau tidak cukup"
		return err, result, dataDetail, false
	}
	

	// pengurangan cash 
	dataCustomer.Cash -= dataDetail.Bill

	// pegurangan stock item
	dataItem.Stock -= dataDetail.Number_of_item

	// update data item
	sql = "UPDATE item SET stock = $1 WHERE id = $2"
	_, err = db.Exec(sql, dataItem.Stock, dataItem.ID)
	if err != nil {
		panic(err)
	}

	// update data customer
	sql = "UPDATE customer SET cash = $1 WHERE id = $2"
	_, err = db.Exec(sql, dataCustomer.Cash, dataCustomer.ID)
	if err != nil {
		panic(err)
	}



	// input riwayat pembelian
	sql = "INSERT INTO order_history(item_id, customer_id, number_of_item, bill) VALUES ($1, $2, $3, $4)"
	_, err = db.Exec(sql,  dataItem.ID, dataCustomer.ID, dataDetail.Number_of_item, dataDetail.Bill)
	if err != nil {
		panic(err)
	}

	result := "Purchase Success"
	return err, result, dataDetail, true

}
