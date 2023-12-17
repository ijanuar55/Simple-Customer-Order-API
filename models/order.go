package models

import (
	"database/sql"
	"dbo/database"
	"dbo/pkg"
	"math"
)

type Order struct {
	ID                int            `json:"id" gorm:"id"`
	TransactionNumber string         `json:"transaction_number" gorm:"transaction_number"`
	TotalPrice        float64        `json:"total_price" gorm:"total_price"`
	Details           []Order_Detail `json:"details" gorm:"-"`
	// CreatedAt         string          `json:"created_at" gorm:"created_at"`
	CreatedBy int `json:"created_by" gorm:"created_by"`
}

type Order_Detail struct {
	ID       int     `json:"id" gorm:"id"`
	OrderID  int     `json:"orders_id" gorm:"orders_id"`
	ItemName string  `json:"item_name" gorm:"item_name"`
	Quantity int     `json:"quantity" gorm:"quantity"`
	Price    float64 `json:"price" gorm:"price"`
}

func GetListOrders(limit int, page int, search string) ([]Order, pkg.Pagination, error) {
	var orders []Order
	var orderCount int64

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	where := "transaction_number LIKE @search OR created_by LIKE @search"

	offset := (page - 1) * limit

	err := database.Database.Debug().Model(&orders).Limit(limit).Offset(offset).Where(where, sql.Named("search", "%"+search+"%")).Find(&orders).Count(&orderCount)

	totalPage := int(math.Ceil(float64(orderCount) / float64(page)))

	if totalPage == 0 {
		totalPage = 1
	}

	var meta pkg.Pagination

	meta.Limit = limit
	meta.Page = page
	meta.TotalRows = orderCount
	meta.TotalPages = totalPage

	for i, _ := range orders {
		orders[i].Details = nil
	}

	return orders, meta, err.Error
}

func GetDetailOrder(id int) (Order, error) {
	var order Order
	var details []Order_Detail

	err := database.Database.Debug().Model(&order).Where("id = ?", id).Find(&order)

	err = database.Database.Debug().Select("*").Table("order_details").Where("orders_id = ?", id).Find(&details)

	order.Details = details

	return order, err.Error
}

func (raw *Order) Save() (*Order, error) {
	err := database.Database.Save(&raw).Error
	if err != nil {
		return &Order{}, err
	}
	return raw, nil
}

func DetailSave(detail []Order_Detail) {
	query := `insert into order_details (orders_id, item_name, quantity, price) values `

	values := []interface{}{}

	for _, row := range detail {
		query += "(?, ?, ?, ?),"
		values = append(values, row.OrderID, row.ItemName, row.Quantity, row.Price)
	}
	//trim the last ,
	query = query[0 : len(query)-1]

	//format all vals at once
	database.Database.Exec(query, values...)

	return
}
