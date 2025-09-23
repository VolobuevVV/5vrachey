package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Promotion struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	IsActive bool   `json:"is_active"`
}

type PromotionHandler struct {
	DB *sql.DB
}

func (h *PromotionHandler) GetPromotionText(c *gin.Context) {
	var promotion Promotion
	err := h.DB.QueryRow("SELECT id, text, is_active FROM promotions WHERE is_active = true ORDER BY id LIMIT 1").Scan(&promotion.ID, &promotion.Text, &promotion.IsActive)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"text": "Скидка 20% на первое посещение"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"text": promotion.Text})
}
