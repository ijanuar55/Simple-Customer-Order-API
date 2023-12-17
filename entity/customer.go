package entity

type GetCustomerReq struct {
	ID     int    `form:"id" uri:"id"`
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	Search string `form:"search"`
}
