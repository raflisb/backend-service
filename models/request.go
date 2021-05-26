package models

type PaginationRequest struct {
	Page    int64 `form:"page"`
	PerPage int64 `form:"per_page"`
}
