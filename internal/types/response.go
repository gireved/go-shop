package types

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 数据内容
}

// 成功响应
func RespSuccess(ctx *gin.Context, data interface{}) Response {
	return Response{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

// 错误响应
func ErrorResponse(ctx *gin.Context, code int, message string, err error) Response {
	return Response{
		Code:    code, // 自定义错误码
		Message: message + err.Error(),
		Data:    nil,
	}
}
