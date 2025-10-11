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
	IsHeadDoctor            bool     `json:"is_head_doctor"`
}

type DoctorHandler struct {
	DB *sql.DB
}

func (h *DoctorHandler) LoadPositionValues() (map[string]string, error) {
	rows, err := h.DB.Query("SELECT id, name FROM positions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	positions := make(map[string]string)
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		positions[id] = name
	}
	return positions, nil
}

func (h *DoctorHandler) GetDoctors(c *gin.Context) {
	positions, err := h.LoadPositionValues()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	rows, err := h.DB.Query("SELECT id, last_name, first_name, patronymic, positions, experience, departments, is_active, available_for_appointment FROM doctors WHERE is_active = true ORDER BY last_name, first_name")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var doctors []Doctor
	for rows.Next() {
		var doctor Doctor
		var positionKeys []string
		err := rows.Scan(
			&doctor.ID,
			&doctor.LastName,
			&doctor.FirstName,
			&doctor.Patronymic,
			pq.Array(&positionKeys),
			&doctor.Experience,
			pq.Array(&doctor.Departments),
			&doctor.IsActive,
			&doctor.AvailableForAppointment,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning doctors"})
			return
		}

		doctor.Positions = make([]string, 0, len(positionKeys))
		for _, key := range positionKeys {
			if name, exists := positions[key]; exists {
				doctor.Positions = append(doctor.Positions, name)
			}
		}

		doctors = append(doctors, doctor)
	}

	c.JSON(http.StatusOK, doctors)
}

func (h *DoctorHandler) GetDoctorsByDepartment(c *gin.Context) {
	departmentID := c.Param("id")

	var headDoctorID string
	h.DB.QueryRow("SELECT head_doctor_id FROM departments WHERE id = $1", departmentID).Scan(&headDoctorID)

	positions, err := h.LoadPositionValues()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	rows, err := h.DB.Query(`
		SELECT id, last_name, first_name, patronymic, positions, experience, departments, is_active, available_for_appointment 
		FROM doctors 
		WHERE is_active = true AND $1 = ANY(departments) 
		ORDER BY 
			CASE WHEN id = $2 THEN 0 ELSE 1 END,
			CASE WHEN available_for_appointment = true THEN 0 ELSE 1 END,
			position,
			last_name, 
			first_name
	`, departmentID, headDoctorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var doctors []Doctor
	for rows.Next() {
		var doctor Doctor
		var positionKeys []string
		err := rows.Scan(
			&doctor.ID,
			&doctor.LastName,
			&doctor.FirstName,
			&doctor.Patronymic,
			pq.Array(&positionKeys),
			&doctor.Experience,
			pq.Array(&doctor.Departments),
			&doctor.IsActive,
			&doctor.AvailableForAppointment,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning doctors"})
			return
		}

		doctor.Positions = make([]string, 0, len(positionKeys))
		for _, key := range positionKeys {
			if name, exists := positions[key]; exists {
				doctor.Positions = append(doctor.Positions, name)
			}
		}

		doctor.IsHeadDoctor = doctor.ID == headDoctorID

		doctors = append(doctors, doctor)
	}

	c.JSON(http.StatusOK, doctors)
}
