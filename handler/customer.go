package handler

import (
	"database/sql"
	"e-commerce_api/database"
	"e-commerce_api/middleware"
	e "e-commerce_api/encrypt"
	"fmt"
	
)

var (
	db *sql.DB = database.DbConnection 
)

func SignUp(db *sql.DB, dataLogin *database.Customer) (error, *database.Customer) {
	// generate hash password
	hashed,_ := e.MakePassword(dataLogin.Password)

	// memasukkan ke dalam database
	sql := "INSERT INTO customer(uname, password, cash) VALUES ($1, $2, $3)"
	_, err2 := db.Exec(sql, dataLogin.Uname, hashed, dataLogin.Cash)
	
	return err2, dataLogin
}


// router login
func LoginCustomer(db *sql.DB, dataLogin *database.Customer) (error, bool, database.Customer, []database.Order_history) {	
	stmt := "SELECT id, uname, cash, password FROM customer WHERE uname = $1"
	rows, err := db.Query(stmt, dataLogin.Uname)
	if err != nil {
		panic(err)
	}

	var dataDB database.Customer

	for rows.Next() {
		 
		err = rows.Scan(&dataDB.ID, &dataDB.Uname, &dataDB.Cash ,&dataDB.Password)
		if err != nil {
			panic(err)
		}
		if dataDB.Uname == dataLogin.Uname && e.CheckPassword([]byte(dataDB.Password), dataLogin.Password) {
			// menampilkan riwayat pembelian
			smtt := "SELECT id, item_id, customer_id, number_of_item, bill, purchased_at FROM order_history WHERE customer_id = $1"
			row, err2 := db.Query(smtt, dataDB.ID)
			fmt.Println(dataDB.ID)
			if err != nil {
				panic(err2)
			}
	
			var orderhistory database.Order_history 
			var listorderhistory []database.Order_history

			for row.Next() {
				err2 = row.Scan(&orderhistory.ID, &orderhistory.ItemID, &orderhistory.Number_of_item, &orderhistory.Bill, &orderhistory.CustomerID, &orderhistory.Time)

				if err != nil {
					panic(err2)
				}
				listorderhistory = append(listorderhistory, orderhistory)
			}
			row.Close()

			return err, true, dataDB, listorderhistory
		}
	}
	rows.Close()
	

	return err, false, dataDB, nil
}


// mengubah data pelanggan
func UpdateCustomer(db *sql.DB, dataCustomer *database.Customer)(error){
	// generate hash password
	hashed,_:= e.MakePassword(dataCustomer.Password) 
	
	sql := "UPDATE customer SET uname = $1, password = $2, cash = $3 WHERE id = $4"
	_,errs := db.Exec(sql, dataCustomer.Uname, hashed, dataCustomer.Cash, dataCustomer.ID)

	fmt.Println(dataCustomer.ID, middleware.SessionUser)
	fmt.Println(dataCustomer.Uname, dataCustomer.Password, dataCustomer.Cash)
	return errs
}

func DeleteCustomer(db *sql.DB, dataCustomer *database.Customer)(error){

	sql := "DELETE FROM customer WHERE id = $1"

	_,errs := db.Exec(sql, dataCustomer.ID)

	return errs
}