package models

type PaginationRequest struct {
	Page    int64 `form:"page"`
	PerPage int64 `form:"per_page"`
}

type UserDataRequest struct {
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required" `
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
