package controllers

import (
	"net/http"
	"wachirawut_agnos_backend/internal/usecases"
	"wachirawut_agnos_backend/pkg/responses"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	uc *usecases.Usecase
}

func NewController(uc *usecases.Usecase) *Controller {
	return &Controller{
		uc: uc,
	}
}

func (c *Controller) StrongPasswordStep(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responses.Ok(http.StatusOK, "successfully get item", nil))
}
