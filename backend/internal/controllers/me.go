package controllers

import (
	"github.com/sdkim96/mcp-marketplace/internal/db"
	"github.com/sdkim96/mcp-marketplace/internal/models"
)

func GetMeCntrl(userName string) (*models.UserDTO, error) {

	var (
		me    *db.UserTable
		meDTO *models.UserDTO
	)
	h := db.GetDBHandler()
	me = &db.UserTable{}
	meDTO = &models.UserDTO{}

	tx := h.First(me, "user_name = ?", userName)

	if tx.Error != nil {
		return meDTO, tx.Error
	}
	meDTO.ID = me.ID
	meDTO.UserName = me.UserName

	return meDTO, nil
}
