package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"student_management_be/dto"
	"student_management_be/models"
	"time"
)

type AccountRepository interface {
	Add(account *models.Account) error
	Get(id int) (*models.Account, error)
	Update(account *models.Account) error
	Delete(id int) error
	List(limit, offset int) ([]*models.Account, error)
}

type AccountController struct {
	repo AccountRepository
}

func NewAccountController(repo AccountRepository) *AccountController {
	return &AccountController{repo: repo}
}

func (c *AccountController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/accounts")
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
		// handle /accounts/{id}
		idStr := strings.TrimPrefix(path, "/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid account id", http.StatusBadRequest)
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

// List accounts
func (c *AccountController) handleList(w http.ResponseWriter, r *http.Request) {
	// mặc định
	limit := 10
	offset := 0

	// lấy từ query param
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

	accounts, err := c.repo.List(limit, offset)
	if err != nil {
		http.Error(w, "failed to get accounts: "+err.Error(), http.StatusInternalServerError)
		return
	}

	res := make([]dto.AccountResponse, 0, len(accounts))
	for _, a := range accounts {
		res = append(res, dto.AccountResponse{
			ID:          a.ID,
			Name:        a.Username,
			Email:       a.Email,
			Birth:       a.Birth.Format("2006-01-02"),
			Role:        a.Role,
			TimeTableID: a.TimeTableID,
			CreatedAt:   a.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Create account
func (c *AccountController) handleCreate(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	birth, err := time.Parse("2006-01-02", req.Birth)
	if err != nil {
		http.Error(w, "invalid birth date", http.StatusBadRequest)
		return
	}

	var timetableID *int
	if req.TimeTableID != nil && *req.TimeTableID != 0 {
		timetableID = req.TimeTableID
	}

	acc := models.NewAccount(0, req.Name, req.Email, birth, req.Role, timetableID)
	if err := c.repo.Add(acc); err != nil {
		http.Error(w, "failed to create account: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto.AccountResponse{
		ID:          acc.ID,
		Name:        acc.Username,
		Email:       acc.Email,
		Birth:       acc.Birth.Format("2006-01-02"),
		Role:        acc.Role,
		TimeTableID: acc.TimeTableID,
		CreatedAt:   acc.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

// Get account by ID
func (c *AccountController) handleGet(w http.ResponseWriter, r *http.Request, id int) {
	acc, err := c.repo.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(dto.AccountResponse{
		ID:          acc.ID,
		Name:        acc.Username,
		Email:       acc.Email,
		Birth:       acc.Birth.Format("2006-01-02"),
		Role:        acc.Role,
		TimeTableID: acc.TimeTableID,
		CreatedAt:   acc.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

// Update account by ID
// Update account by ID
func (c *AccountController) handleUpdate(w http.ResponseWriter, r *http.Request, id int) {
	// Decode body JSON
	var req dto.CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Parse birth date
	birth, err := time.Parse("2006-01-02", req.Birth)
	if err != nil {
		http.Error(w, "invalid birth date", http.StatusBadRequest)
		return
	}

	// Lấy account hiện tại từ DB
	acc, err := c.repo.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Cập nhật dữ liệu
	acc.SetName(req.Name)
	acc.SetEmail(req.Email)
	acc.SetRole(req.Role)
	acc.SetBirth(birth)

	if req.TimeTableID != nil && *req.TimeTableID != 0 {
		acc.SetTimeTableID(req.TimeTableID)
	} else {
		acc.SetTimeTableID(nil)
	}

	// Update trong DB
	if err := c.repo.Update(acc); err != nil {
		http.Error(w, "failed to update account: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Trả về response
	json.NewEncoder(w).Encode(dto.AccountResponse{
		ID:          acc.ID,
		Name:        acc.Username,
		Email:       acc.Email,
		Birth:       acc.Birth.Format("2006-01-02"),
		Role:        acc.Role,
		TimeTableID: acc.TimeTableID,
		CreatedAt:   acc.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

// Delete account by ID
func (c *AccountController) handleDelete(w http.ResponseWriter, r *http.Request, id int) {
	if err := c.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
