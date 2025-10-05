package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"student_management_be/dto"
	"student_management_be/models"
)

type ClassroomRepository interface {
	Add(classroom *models.Classroom) error
	Get(id int) (*models.Classroom, error)
	Update(classroom *models.Classroom) error
	Delete(id int) error
	List(limit, offset int) ([]*models.Classroom, error)
}

type ClassroomController struct {
	repo ClassroomRepository
}

func NewClassroomController(repo ClassroomRepository) *ClassroomController {
	return &ClassroomController{repo: repo}
}

func (c *ClassroomController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/classrooms")
	if path == "" || path == "/" {
		switch r.Method {
		case http.MethodGet:
			c.handleList(w, r)
		case http.MethodPost:
			c.handleCreate(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	} else {
		// /classrooms/{id}
		idStr := strings.TrimPrefix(path, "/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid classroom id", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			c.handleGet(w, r, id)
		case http.MethodPut:
			c.handleUpdate(w, r, id)
		case http.MethodDelete:
			c.handleDelete(w, r, id)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// List classrooms
func (c *ClassroomController) handleList(w http.ResponseWriter, r *http.Request) {
	limit := 10
	offset := 0

	if l := r.URL.Query().Get("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil && val > 0 {
			limit = val
		}
	}
	if o := r.URL.Query().Get("offset"); o != "" {
		if val, err := strconv.Atoi(o); err == nil && val >= 0 {
			offset = val
		}
	}

	classrooms, err := c.repo.List(limit, offset)
	if err != nil {
		http.Error(w, "failed to get classrooms: "+err.Error(), http.StatusInternalServerError)
		return
	}

	res := make([]dto.ClassroomResponse, 0, len(classrooms))
	for _, cl := range classrooms {
		res = append(res, dto.ClassroomResponse{
			ID:          cl.ID,
			Name:        cl.Name,
			Total:       cl.Total,
			TimeTableID: cl.TimeTableID,
			CreatedAt:   cl.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Create classroom
func (c *ClassroomController) handleCreate(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateClassroomRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	classroom := models.NewClassroom(0, req.Name, req.Total, req.TimeTableID)
	if err := c.repo.Add(classroom); err != nil {
		http.Error(w, "failed to create classroom: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto.ClassroomResponse{
		ID:          classroom.ID,
		Name:        classroom.Name,
		Total:       classroom.Total,
		TimeTableID: classroom.TimeTableID,
		CreatedAt:   classroom.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

// Get classroom by ID
func (c *ClassroomController) handleGet(w http.ResponseWriter, r *http.Request, id int) {
	classroom, err := c.repo.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(dto.ClassroomResponse{
		ID:          classroom.ID,
		Name:        classroom.Name,
		Total:       classroom.Total,
		TimeTableID: classroom.TimeTableID,
		CreatedAt:   classroom.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

// Update classroom
func (c *ClassroomController) handleUpdate(w http.ResponseWriter, r *http.Request, id int) {
	var req dto.CreateClassroomRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	classroom, err := c.repo.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	classroom.SetName(req.Name)
	classroom.SetTotal(req.Total)
	if req.TimeTableID != nil && *req.TimeTableID != 0 {
		classroom.SetTimeTableID(req.TimeTableID)
	} else {
		classroom.SetTimeTableID(nil)
	}

	if err := c.repo.Update(classroom); err != nil {
		http.Error(w, "failed to update classroom: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(dto.ClassroomResponse{
		ID:          classroom.ID,
		Name:        classroom.Name,
		Total:       classroom.Total,
		TimeTableID: classroom.TimeTableID,
		CreatedAt:   classroom.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

// Delete classroom
func (c *ClassroomController) handleDelete(w http.ResponseWriter, r *http.Request, id int) {
	if err := c.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
