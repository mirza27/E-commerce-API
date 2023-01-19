package handler

import (
	"e-commerce_api/database"
	
	"database/sql"
	
)


func GetAllItem(db *sql.DB) (err error, show []database.Item){
	
	sql := "SELECT id, name, stock, price, description, category_id FROM item"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var item_list = database.Item{}
		 
		err = rows.Scan(&item_list.ID, &item_list.Name, &item_list.Stock,&item_list.Price, &item_list.Description, &item_list.Category_id)
		if err != nil {
			panic(err)
		}

		show = append(show, item_list)
	}

	return 
}

func InsertItem(db *sql.DB, dataItem *database.Item) (error, *database.Item, string, bool){
	// cek price tidak boleh <= 0
	if dataItem.Price <= 0{
		message := "harga tidak boleh kurang dari 0"
		return nil, dataItem, message, false
	}

	sql := "INSERT INTO item(name, description, stock, price, category_id) VALUES ($1, $2, $3, $4, $5)"

	_,errs := db.Exec(sql, dataItem.Name, dataItem.Description, dataItem.Stock, dataItem.Price, dataItem.Category_id)


	message := "sukses memasukkan data"
	return errs, dataItem, message, true
}


func DeleteItem(db *sql.DB, dataItem *database.Item)(error){

	sql := "DELETE FROM Item WHERE id = $1"

	_,errs := db.Exec(sql, dataItem.ID)

	return errs
}


func UpdateItem(db *sql.DB, dataItem *database.Item)(error, *database.Item,string, bool){
	// cek price tidak boleh <= 0
	if dataItem.Price <= 0{
		message := "harga tidak boleh kurang dari 0"
		return nil, dataItem, message, false
	}
	
	sql := "UPDATE item SET name = $1, description = $2, price = $3, stock = $4, category_id = $5 WHERE id = $6"
	
	_,errs := db.Exec(sql, dataItem.Name, dataItem.Description, dataItem.Price, dataItem.Stock, dataItem.Category_id, dataItem.ID)
	
	message := "Sukses mengupdate data"
	return errs, dataItem, message, true
	
	}
