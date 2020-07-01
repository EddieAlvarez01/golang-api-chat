package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/EddieAlvarez01/golang-api-chat/database"
	"github.com/EddieAlvarez01/golang-api-chat/models"
)

func Test(w http.ResponseWriter, r *http.Request) {
	var roles models.Role
	db := database.DBConn
	db.First(&roles)
	//db.Table("Role").Where("id = ?", "3").Delete(&models.Role{})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}
