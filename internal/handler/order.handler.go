package handler

import (
	"net/http"
	"stock-and-order-management/internal/domain"
	"stock-and-order-management/internal/dto"
	"stock-and-order-management/shared/error"
	"stock-and-order-management/shared/response"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	usecase domain.OrderUsecase
}

func NewOrderHandler(usecase domain.OrderUsecase) *orderHandler {
	return &orderHandler{
		usecase: usecase,
	}
}

func (h *orderHandler) CreateOrder(c *gin.Context) {
	body := dto.CreateOrderRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, error.ErrorResponse{
			Error: error.ErrorDetailResponse{
				Code:    error.ERR_BAD_REQUEST,
				Message: "invalid form",
			},
		})
		return
	}

	if err := h.usecase.Create(body); err != nil {
		if err.Error() == "QUANTITY_NOT_ENOUGH" {
			c.AbortWithStatusJSON(http.StatusBadRequest, error.ErrorResponse{
				Error: error.ErrorDetailResponse{
					Code:    error.ERR_BAD_REQUEST,
					Message: "quantity not enough",
				},
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, error.ErrorResponse{
			Error: error.ErrorDetailResponse{
				Code:    error.ERR_INTERNAL_SERVER_ERROR,
				Message: http.StatusText(http.StatusInternalServerError),
			},
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusCreated, response.SuccessfullyResponse{
		Message: "Create order is successfully",
	})
}
