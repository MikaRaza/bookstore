package handlers

import (
	"database/sql"
	"example/bookstore/database"
	"example/bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBooks récupère tous les livres
func GetBooks(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, title, author, price FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

// PostBooks ajoute un nouveau livre
func PostBooks(c *gin.Context) {
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Utilisez Exec pour insérer le livre et récupérer l'ID généré
	result, err := database.DB.Exec("INSERT INTO books (title, author, price) VALUES (?, ?, ?)",
		book.Title, book.Author, book.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Récupérer l'ID généré
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convertir l'ID en uint et mettre à jour le modèle
	book.ID = uint(id)

	c.JSON(http.StatusCreated, book)
}

// GetBookByID récupère un livre par son ID
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	row := database.DB.QueryRow("SELECT id, title, author, price FROM books WHERE id = ?", id)

	var book models.Book
	if err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Price); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, book)
}
