package util

import (
	"strconv"

	"main/pkg/setting"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	result := 0

	page, _ := strconv.Atoi((c.Query("page")))

	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result

}
