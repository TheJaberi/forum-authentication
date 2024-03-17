package forum

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define a struct to match the incoming JSON structure
type GoogleAuthPayload struct {
	Sub        string `json:"sub"`         // User ID
	Name       string `json:"name"`        // Full Name
	GivenName  string `json:"given_name"`  // Given Name
	FamilyName string `json:"family_name"` // Family Name
	Picture    string `json:"picture"`     // Image URL
	Email      string `json:"email"`       // Email
}

func HandlerReceiveCodeGoogle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/receive-code/google" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	var payload GoogleAuthPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Received Google Auth for user: %s\n", payload.Email)

	// Process the authentication payload as needed

	// Respond to the request indicating success
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
