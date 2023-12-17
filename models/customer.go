package models

import (
	"database/sql"
	"dbo/database"
	"dbo/pkg"
	"errors"
	"math"

	"gorm.io/gorm"
)

type Customer struct {
	ID      int    `json:"id" gorm:"id"`
	Name    string `json:"name" gorm:"name"`
	Email   string `json:"email" gorm:"email"`
	Address string `json:"address" gorm:"address"`
}

func GetListCustomers(limit int, page int, search string) ([]Customer, pkg.Pagination, error) {
	var customers []Customer
	var custCount int64

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	where := "name LIKE @search OR email LIKE @search OR address LIKE @search"

	offset := (page - 1) * limit

	err := database.Database.Debug().Model(&customers).Limit(limit).Offset(offset).Where(where, sql.Named("search", "%"+search+"%")).Find(&customers).Count(&custCount)

	totalPage := int(math.Ceil(float64(custCount) / float64(page)))

	if totalPage == 0 {
		totalPage = 1
	}

	var meta pkg.Pagination

	meta.Limit = limit
	meta.Page = page
	meta.TotalRows = custCount
	meta.TotalPages = totalPage

	return customers, meta, err.Error
}

func GetDetailCustomer(id int) (Customer, error) {
	var customer Customer

	err := database.Database.Debug().Model(&customer).Where("id = ?", id).Find(&customer)

	return customer, err.Error
}

func (raw *Customer) Save() (*Customer, error) {
	err := database.Database.Save(&raw).Error
	if err != nil {
		return &Customer{}, err
	}
	return raw, nil
}

func (c *Customer) BeforeDelete(tx *gorm.DB) (err error) {
	if c.ID == 0 {
		return errors.New("ID Customer wajib diisi")
	}
	return
}

func (raw *Customer) Delete() (*Customer, error) {
	err := database.Database.Delete(&raw).Error
	if err != nil {
		return &Customer{}, err
	}
	return raw, nil
}
