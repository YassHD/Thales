package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func openBase() {
	//Ouvre la connexion à la base de données
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	router := gin.Default()
	// Démarre le serveur
	openBase()
	// Ouvre le Database

	router.Static("/static", "./static") // charge le css et le script js
	router.LoadHTMLGlob("tmpl/*.html")   // charge les pages HTML

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "inscription.html", gin.H{})
	})

}
