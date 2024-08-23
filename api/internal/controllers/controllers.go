package controllers

import (
	"fmt"
	"net/http"
	"wachirawut_agnos_backend/internal/domain"
	"wachirawut_agnos_backend/internal/usecases"
	"wachirawut_agnos_backend/pkg/responses"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	uc domain.Usecase
}

func NewController(uc *usecases.Usecase) *Controller {
	return &Controller{
		uc: uc,
	}
}

func (c *Controller) StrongPasswordStep(ctx *gin.Context) {
	// mapping data from request to StrongPasswordStep model
	var strongPasswordStep StrongPasswordStep

	if err := ctx.ShouldBindJSON(&strongPasswordStep); err != nil {
		fmt.Printf("error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	strongPasswordDto := &domain.StrongPasswordStepDtO{
		Password: strongPasswordStep.InitPassword,
	}

	// call usecase to check strong password step
	numOfSteps, err := c.uc.CheckStrongPasswordStep(strongPasswordDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, numOfSteps)
}
