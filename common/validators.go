package common

import (
	"ttany-chat-service/utils/paginator"

	"github.com/gin-gonic/gin"
)

type PaginationValidator struct {
	AfterCursor string `form:"cursor"`
	Limit       int    `form:"limit"`
}

type PagingQuery struct {
	AfterCursor *string
	Limit       *int
}

func (self *PaginationValidator) BindQueryString(c *gin.Context, key1, key2 string) (p paginator.Paginator, err error) {
	if err = c.ShouldBindQuery(&self); err != nil {
		return p, err
	}

	pq := PagingQuery{AfterCursor: &self.AfterCursor, Limit: &self.Limit}

	p = paginator.New()
	p.SetKeys(key1, key2) // [defualt: "ID"] (order of keys matters)

	if pq.AfterCursor != nil {
		p.SetAfterCursor(*pq.AfterCursor) // [default: ""]
	}

	if pq.Limit != nil {
		p.SetLimit(*pq.Limit) // [default: 10]
	}
	return p, nil
}

func NewPaginationValidator() PaginationValidator {
	paginationValidator := PaginationValidator{}
	return paginationValidator
}
