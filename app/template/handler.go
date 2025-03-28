package template

import (
	"go-template/app"

	"github.com/gin-gonic/gin"
)

type service interface {
	Process(req Request) (*Response, error)
}

type Request struct {
	ID int `json:"id"`
}

type Handler struct {
	service service
}

func NewHandler(service service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Handler(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		app.BadRequest(c)
		return
	}

	resp, err := h.service.Process(req)
	if err != nil {
		app.Error(c, err.Error())
		return
	}

	app.Success(c, resp)
}
