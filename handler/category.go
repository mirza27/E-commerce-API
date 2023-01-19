package handler

import (
	"e-commerce_api/database"
	"database/sql"
	
)


func InsertCategories(db *sql.DB, dataCategory *database.Category)(error){
	sql := "INSERT INTO category(name) VALUES ($1)"


	_,errs := db.Exec(sql, dataCategory.Name)

	return errs
}


func GetItemByCategory(db *sql.DB, dataCategory *database.Category)(err error, results []database.Item){
	stmt := "SELECT id FROM category WHERE name = $1"
	row, err := db.Query(stmt, dataCategory.Name)
	if err != nil {
		panic(err)
	}

	defer row.Close()
	for row.Next() {
		err = row.Scan(&dataCategory.ID)
		if err != nil {
			panic(err)
		}
	}
	
	// menampilkan yang sesuai kategori
	sql := "SELECT id, name, description, price, stock, category_id FROM item WHERE category_id = $1"

	rows, err := db.Query(sql, dataCategory.ID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var dataItem = database.Item{}

	for rows.Next() {
		err = rows.Scan(&dataItem.ID, &dataItem.Name, &dataItem.Description, &dataItem.Price, &dataItem.Stock, &dataItem.Category_id)
		if err != nil {
			panic(err)
		}

		results = append(results, dataItem)
	}

	return
}


func UpdateCategory(db *sql.DB, dataCategory *database.Category)(error, *database.Category,string, bool){
	// cek price tidak boleh <= 0
	if dataCategory.Name == ""{
		message := "Kategory tidak boleh kossong"
		return nil, dataCategory, message, false
	}
	
	sql := "UPDATE category SET name = $1 WHERE id = $2"
	
	_,errs := db.Exec(sql, dataCategory.Name, dataCategory.ID)
	
	message := "Sukses mengupdate data"
	return errs, dataCategory, message, true
	
	}
