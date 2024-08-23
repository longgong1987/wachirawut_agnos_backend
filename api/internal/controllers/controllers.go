package controllers

import (
	"net/http"
	"wachirawut_agnos_backend/internal/usecases"
	"wachirawut_agnos_backend/pkg/responses"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	StrongPasswordStep(ctx *gin.Context)
}

type controller struct {
	uc usecases.Usecase
}

func NewController(uc usecases.Usecase) Controller {
	return &controller{
		uc: uc,
	}
}

func (c *controller) StrongPasswordStep(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responses.Ok(http.StatusOK, "successfully get items", nil))
}
