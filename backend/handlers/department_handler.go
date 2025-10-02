package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Department struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Color    string `json:"color"`
	IsActive bool   `json:"is_active"`
}

type DepartmentHandler struct {
	DB *sql.DB
}

func (h *DepartmentHandler) GetDepartments(c *gin.Context) {
	rows, err := h.DB.Query("SELECT id, name, color, is_active FROM departments WHERE is_active = true ORDER BY id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var departments []Department
	for rows.Next() {
		var dept Department
		err := rows.Scan(&dept.ID, &dept.Name, &dept.Color, &dept.IsActive)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning departments"})
			return
		}
		departments = append(departments, dept)
	}

	c.JSON(http.StatusOK, departments)
}
