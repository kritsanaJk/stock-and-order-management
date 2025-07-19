package handler

import (
	"net/http"
	"stock-and-order-management/internal/domain"
	"stock-and-order-management/internal/dto"
	"stock-and-order-management/shared/error"
	"stock-and-order-management/shared/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	usecase domain.ProductUsecase
}

func NewProductHandler(usecase domain.ProductUsecase) *productHandler {
	return &productHandler{
		usecase: usecase,
	}
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	body := dto.CreateProductRequest{}

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
		if err.Error() == error.ERR_CONFLICT {
			c.AbortWithStatusJSON(http.StatusConflict, error.ErrorResponse{
				Error: error.ErrorDetailResponse{
					Code:    error.ERR_CONFLICT,
					Message: "Product is exist",
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
		Message: "Create product is successfully",
	})
}

func (h *productHandler) List(c *gin.Context) {
	res := []dto.ProductResponse{}
	if err := h.usecase.List(&res); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, error.ErrorResponse{
			Error: error.ErrorDetailResponse{
				Code:    error.ERR_INTERNAL_SERVER_ERROR,
				Message: http.StatusText(http.StatusInternalServerError),
			},
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, response.SuccessfullyResponse{
		Message: "Get list product is successfully",
		Items:   res,
	})

}

func (h *productHandler) GetByID(c *gin.Context) {
	productIDStr := c.Param("id")
	res := dto.ProductResponse{}

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, error.ErrorResponse{
			Error: error.ErrorDetailResponse{
				Code:    error.ERR_BAD_REQUEST,
				Message: "id is invalid",
			},
		})
		return
	}

	if err := h.usecase.Product((productID), &res); err != nil {
		if err.Error() == error.ERR_NOT_FOUND {
			c.AbortWithStatusJSON(http.StatusNotFound, error.ErrorResponse{
				Error: error.ErrorDetailResponse{
					Code:    error.ERR_NOT_FOUND,
					Message: "Product is not found",
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
	c.AbortWithStatusJSON(http.StatusOK, response.SuccessfullyResponse{
		Message: "Get  product is successfully",
		Item:    res,
	})

}
