package models

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
	Total int64  `json:"total"`
}

func (p *Pagination) Init(c *gin.Context) {
	limit := 12
	page := 1
	sort := `id ASC`
	query := c.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}
	}
	p.Limit = limit
	p.Page = page
	p.Sort = sort

}
