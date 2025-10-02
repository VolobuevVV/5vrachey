package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type Doctor struct {
	ID                      string   `json:"id"`
	LastName                string   `json:"last_name"`
	FirstName               string   `json:"first_name"`
	Patronymic              string   `json:"patronymic"`
	Positions               []string `json:"positions"`
	Experience              int      `json:"experience"`
	Departments             []string `json:"departments"`
	IsActive                bool     `json:"is_active"`
	AvailableForAppointment bool     `json:"available_for_appointment"`
}

type DoctorHandler struct {
	DB *sql.DB
}

func (h *DoctorHandler) GetDoctors(c *gin.Context) {
	rows, err := h.DB.Query("SELECT id, last_name, first_name, patronymic, positions, experience, departments, is_active, available_for_appointment FROM doctors WHERE is_active = true ORDER BY last_name, first_name")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var doctors []Doctor
	for rows.Next() {
		var doctor Doctor
		err := rows.Scan(
			&doctor.ID,
			&doctor.LastName,
			&doctor.FirstName,
			&doctor.Patronymic,
			pq.Array(&doctor.Positions),
			&doctor.Experience,
			pq.Array(&doctor.Departments),
			&doctor.IsActive,
			&doctor.AvailableForAppointment,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning doctors"})
			return
		}
		doctors = append(doctors, doctor)
	}

	c.JSON(http.StatusOK, doctors)
}

func (h *DoctorHandler) GetDoctorsByDepartment(c *gin.Context) {
	departmentID := c.Param("id")

	rows, err := h.DB.Query("SELECT id, last_name, first_name, patronymic, positions, experience, departments, is_active, available_for_appointment FROM doctors WHERE is_active = true AND $1 = ANY(departments) ORDER BY last_name, first_name", departmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var doctors []Doctor
	for rows.Next() {
		var doctor Doctor
		err := rows.Scan(
			&doctor.ID,
			&doctor.LastName,
			&doctor.FirstName,
			&doctor.Patronymic,
			pq.Array(&doctor.Positions),
			&doctor.Experience,
			pq.Array(&doctor.Departments),
			&doctor.IsActive,
			&doctor.AvailableForAppointment,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning doctors"})
			return
		}
		doctors = append(doctors, doctor)
	}

	c.JSON(http.StatusOK, doctors)
}
