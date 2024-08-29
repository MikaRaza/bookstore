package main

import (
	"example/bookstore/database"
	"example/bookstore/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Charger les variables d'environnement
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connecter à la base de données
	database.ConnectDatabase()

	// Initialiser le routeur Gin
	r := gin.Default()

	// Enregistrer les routes des livres
	routes.RegisterBookRoutes(r)

	// Démarrer le serveur sur le port 8080
	r.Run(":8080")
}
