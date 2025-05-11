package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sdkim96/mcp-marketplace/internal/middleware"
	"github.com/sdkim96/mcp-marketplace/internal/models"
)

func HealthCheck(c *gin.Context) {

	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func Login(c *gin.Context) {

	loginRequest := models.NewLoginRequest()
	err := c.ShouldBindJSON(loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse[any]{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	secret, ok := c.Get("projectSecret")

	if !ok {
		log.Println("projectSecret not found in context")
		c.JSON(http.StatusInternalServerError, models.APIResponse[any]{
			Success: false,
			Data:    nil,
			Error:   "Serious Server Error",
		})
	}

	token := middleware.GetJWTtoken(
		secret.(string),
		loginRequest.UserName,
		60*60*24,
	)
	c.JSON(http.StatusOK, models.APIResponse[any]{
		Success: true,
		Data:    models.LoginRespBody{Token: token},
		Error:   nil,
	})

}
