package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/EddieAlvarez01/golang-api-chat/database"
	"github.com/EddieAlvarez01/golang-api-chat/handlers"
	"github.com/gorilla/mux"
)

var router *mux.Router //RUTAS

//Inicializar las rutas del servidor
func init() {
	initDB()
	router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlers.Test)
	router.HandleFunc("/user/create-guest", handlers.AddGuest).Methods("POST")
}

func initDB() {
	var err error
	database.DBConn, err = gorm.Open("mysql", "root:@(localhost:3306)/db_chat_golang?charset=utf8&parseTime=True")
	if err != nil {
		panic("Error al conectar a la base de datos")
	}
	fmt.Println("Database is connected")
}

func main() {
	defer database.DBConn.Close() //CIERRA LA CONEXION CUANDO MAIN SE TERMINE DE EJECUTAR
	fmt.Println("Server on port 4500")
	log.Fatal(http.ListenAndServe(":4500", router))
}
