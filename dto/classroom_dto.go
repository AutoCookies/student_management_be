package dto

type CreateClassroomRequest struct {
	Name        string `json:"name"`
	Total       int    `json:"total"`
	TimeTableID *int   `json:"timetableId"`
}

type ClassroomResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Total       int    `json:"total"`
	TimeTableID *int   `json:"timetableId"`
	CreatedAt   string `json:"createdAt"`
}
