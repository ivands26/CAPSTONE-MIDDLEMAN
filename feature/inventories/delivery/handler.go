package delivery

import (
	"log"
	"middleman-capstone/domain"
	"middleman-capstone/feature/common"
	"net/http"

	"github.com/labstack/echo/v4"
)

type inventoryHandler struct {
	inventoryUseCase domain.InventoryUseCase
}

func New(iuc domain.InventoryUseCase) domain.InventoryHandler {
	return &inventoryHandler{
		inventoryUseCase: iuc,
	}
}

func (ih *inventoryHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newInventory InventoryFormat
		id, role := common.ExtractData(c)
		bind := c.Bind(&newInventory)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "there is an error in internal server",
			})
		}

		if role != "user" {
			log.Println("you dont have access")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "you dont have access",
			})
		}

		status := ih.inventoryUseCase.CreateInventory(newInventory.ToIP(), id)

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "wrong input",
			})
		}
		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "data not found",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "there is an error in internal server",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "success create product",
		})
	}
}