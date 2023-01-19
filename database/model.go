package database

import (
	"time"
)

type Customer struct {
	ID int64 `json:"id"`
	Uname string `json:"uname"`
	Password string `json:"password"`
	Cash int64 `json:"cash"`
}

type Item struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Stock int64 `json:"stock"`
	Description string `json:"description"`
	Price int64 `json:"price"`
	Category_id int64 `json:"category_id"`
}

type Order_history struct {
	ID int64 `json:"id"`
	ItemID int64 `json:"item"`
	CustomerID int64 `json:"customer"`
	Number_of_item int64 `json:"number_of_item"`
	Bill int64 `json:"bill"`
	Time time.Time `json:"time"`
}

type Category struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}