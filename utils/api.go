package utils

import (
	"time"
)

// ResTime 返回UTC时间，用ISO8601格式化。YYYY-MM-DDTHH:MM:SSZ
func ResTime(t time.Time) string {
	return t.UTC().Format("2006-01-02T15:04:05Z")
}

// ReqDate 将客户端提交的日期（比如搜索时的日期范围选择），换算成UTC日期
func ReqDate(s string) time.Time {
	t, _ := time.ParseInLocation("2006-01-02", s, time.Local)
	return t
}

// ResError 返回错误，JSON格式，例{ "errcode": 000001, "errmsg": "参数错误" }
func ResError(c *gin.Context, status int, errcode int, errmsg string) {
	c.JSON(status, gin.H{
		"errcode": errcode,
		"errmsg":  errmsg,
	})
}

// ResAPIError 返回后端api间调用错误，status始终为200
func ResAPIError(c *gin.Context, errcode int, errmsg string) {
	c.JSON(200, gin.H{
		"errcode": errcode,
		"errmsg":  errmsg,
	})
}
