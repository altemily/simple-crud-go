package main

import (
	"log"
	"net/http"
	"github.com/altemily/simple-go-mod/config"
	"github.com/altemily/simple-go-mod/models"
	
)

// Função chamada quando minha aplicação é iniciada. É o ponto de entrada da aplicação.
func main(){
	dbConnection := config.SetupDB()
	_, err := dbConnection.Exec(models.CreateTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	defer dbConnection.Close()

	log.Fatal(http.ListenAndServe(":8080", nil))
}