package dto

type CreateAccountRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Birth       string `json:"birth"`
	Role        string `json:"role"`
	TimeTableID *int   `json:"timetableId,omitempty"`
}

type AccountResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Birth       string `json:"birth"`
	Role        string `json:"role"`
	TimeTableID *int   `json:"timetableId,omitempty"`
	CreatedAt   string `json:"createdAt"`
}
