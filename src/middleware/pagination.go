package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"

	cm "github.com/zbd20/go-utils/models"
)

func Page() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {
			return
		}
		var pageSize int64 = 10
		var curPage int64 = 1
		if ps, err := strconv.ParseInt(c.Query("limit"), 10, 64); err == nil && ps > 0 {
			pageSize = ps
		}
		if p, err := strconv.ParseInt(c.Query("page"), 10, 64); err == nil && p > 0 {
			curPage = p
		}

		offset := (curPage - 1) * pageSize
		if offset < 0 {
			offset = 0
		}

		page := cm.Page{
			PageSize: pageSize,
			Offset:   offset,
			Page:     curPage,
			Query:    c.Query("query"),
		}

		switch c.Query("sort") {
		case "asc":
			page.Sort = "asc"
		case "desc":
			page.Sort = "desc"
		}

		switch c.Query("order_by") {
		case "name":
			page.OrderBy = "name"
		default:
			page.OrderBy = "update_time"
		}

		c.Set("page", page)

		c.Next()
	}
}
