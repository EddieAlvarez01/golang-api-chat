package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/EddieAlvarez01/golang-api-chat/database"
	"github.com/EddieAlvarez01/golang-api-chat/models"
	"github.com/EddieAlvarez01/golang-api-chat/responses"
)

//AddGuest REGISTRA UN INVITADO AL CHAT
func AddGuest(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserta un valor valido")
	}
	db := database.DBConn
	var role models.Role
	var user models.User
	var userTest models.User
	json.Unmarshal(req, &user)
	w.Header().Set("Content-Type", "application/json")
	db.Preload("Role").Where(&models.User{Username: user.Username}).First(&userTest)
	if userTest.ID != 0 {
		switch userTest.Role.Name {
		case "USER":
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(responses.ErrorResponse{Code: "403", Message: "Un usuario ya ha registrado el nick"})
			return
		case "INVITADO":
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(userTest)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(responses.ErrorResponse{Code: "500", Message: "Error en el servidor"})
			return
		}
	}
	db.Where(&models.Role{Name: "INVITADO"}).First(&role)
	user.RoleID = role.ID
	db.Create(&user)
	if db.NewRecord(user) {
		//NO SE CREO EL USUARIO
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.ErrorResponse{Code: "500", Message: "Error en el servidor"})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	return
}
