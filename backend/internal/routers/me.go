package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/sdkim96/mcp-marketplace/internal/controllers"
	"github.com/sdkim96/mcp-marketplace/internal/models"
)

func GetMe(c *gin.Context) {

	userName := c.Keys["userName"].(string)

	me, err := controllers.GetMeCntrl(userName)
	if err != nil {
		c.JSON(500, models.APIResponse[models.UserDTO]{
			Success: false,
			Data:    *me,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(200, models.APIResponse[models.UserDTO]{
		Success: true,
		Data:    *me,
		Error:   nil,
	})
}
