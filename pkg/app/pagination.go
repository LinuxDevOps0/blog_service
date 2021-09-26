package app

import (
	"github.com/gin-gonic/gin"
	"practice.com/blog_service/global"
	"practice.com/blog_service/pkg/convert"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page < 1 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

func GetOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result =  (page - 1) * pageSize
	}
	return result
}
