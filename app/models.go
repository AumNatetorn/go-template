package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(code, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (r *Response) JSON(c *gin.Context, httpStatus int) {
	c.JSON(httpStatus, r)
}

func Success(c *gin.Context, data interface{}) {
	NewResponse("S0000", "success", data).JSON(c, http.StatusOK)
}

func Error(c *gin.Context, message string) {
	NewResponse("E9999", message, nil).JSON(c, http.StatusInternalServerError)
	c.Abort()
}

func BadRequest(c *gin.Context) {
	NewResponse("E9000", "bad request", nil).JSON(c, http.StatusBadRequest)
	c.Abort()
}

func Unauthorized(c *gin.Context) {
	NewResponse("E9001", "unauthorize", nil).JSON(c, http.StatusUnauthorized)
	c.Abort()
}

func Exceed(c *gin.Context) {
	NewResponse("E4290", "too many requests", nil).JSON(c, http.StatusTooManyRequests)
	c.Abort()
}

func Cancel(c *gin.Context, data interface{}) {
	NewResponse("S202", "cancel process success", data).JSON(c, http.StatusOK)
}
