package handlers

import (
	"encoding/json"
	"my-go-server/api/db"
	"net/http"
)

// DeleteUserResponse defines the JSON structure for responses to delete user requests.
type DeleteUserResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// HealthCheckHandler provides an HTTP endpoint for health checks.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// DeleteUserHandler handles HTTP requests to delete a user by their ID.
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "User ID not provided", http.StatusBadRequest)
		return
	}

	err := db.DeleteUser(userID)
	response := DeleteUserResponse{}

	if err != nil {
		response.Error = err.Error()
		jsonResponse, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResponse)
		return
	}

	response.Message = "User deleted successfully"
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// GetUsersHandler handles HTTP requests to retrieve a list of all users from the database.
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
