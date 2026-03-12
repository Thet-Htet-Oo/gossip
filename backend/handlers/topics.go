package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"gossip-backend/db"

	"github.com/gin-gonic/gin"
)

type Topic struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetTopics(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, title, description FROM topics ORDER BY id DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	defer rows.Close()

	var topics []Topic

	for rows.Next() {
		var t Topic
		rows.Scan(&t.ID, &t.Title, &t.Description)
		topics = append(topics, t)
	}

	c.JSON(http.StatusOK, topics)
}

func CreateTopic(c *gin.Context) {
	var topic Topic

	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	fmt.Println("Received topic:", topic)

	err := db.DB.QueryRow(
		"INSERT INTO topics (title, description) VALUES ($1,$2) RETURNING id",
		topic.Title,
		topic.Description,
	).Scan(&topic.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusOK, topic)
}

func DeleteTopic(c *gin.Context) {
	// Get topic ID from URL
	topicID := c.Param("id")

	_, err := db.DB.Exec("DELETE FROM topics WHERE id = $1", topicID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Topic deleted"})
}

func UpdateTopic(c *gin.Context) {
	topicID := c.Param("id")

	var topic Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Update topic
	result, err := db.DB.Exec(
		"UPDATE topics SET title = $1, description = $2 WHERE id = $3",
		topic.Title, topic.Description, topicID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "topic not found"})
		return
	}

	topic.ID = atoi(topicID)
	c.JSON(http.StatusOK, topic)
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
