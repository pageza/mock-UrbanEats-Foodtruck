package api

import (
	"encoding/json"
	"net/http"

	"github.com/pageza/mock-UrbanEats-Foodtruck/db"
)

// GetMenuHandler handles the requests for the food truck menu
func GetMenuHandler(w http.ResponseWriter, r *http.Request) {
	menuItems, err := db.GetMenuItems() // You'll need to implement GetMenuItems in your db package
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menuItems)
}

// PostFeedbackHandler handles the submission of customer feedback
func PostFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	// Here, you would parse the request body to extract feedback data
	// and then store it using your database logic
	// For example:
	// var feedback db.Feedback
	// err := json.NewDecoder(r.Body).Decode(&feedback)
	// ... handle errors and store feedback ...

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Feedback submitted successfully"))
}

// HealthCheckHandler provides a basic health check endpoint
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is running"))
}

// Other handlers can be added here as needed
