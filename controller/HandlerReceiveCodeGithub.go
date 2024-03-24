package forum

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	m "forum/model"
)

// Assuming CLIENT_ID and CLIENT_SECRET are constants defined elsewhere
const (
	CLIENT_ID     = "7e5c9074f38b520b3e48"
	CLIENT_SECRET = "959bed0370ec86564dea2c3f8f8ab84420b147be"
)

func HandlerReceiveCodeGithub(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/receive-code/github" {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	if req.Method != "POST" {
		ErrorHandler(w, req, http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		Code string `json:"code"`
	}
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if requestBody.Code == "" {
		http.Error(w, "missing code parameter", http.StatusBadRequest)
		return
	}

	// Prepare data for the POST request to GitHub
	data := url.Values{}
	data.Set("client_id", CLIENT_ID)
	data.Set("client_secret", CLIENT_SECRET)
	data.Set("code", requestBody.Code)

	// Create the request to GitHub
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Add("Accept", "application/json") // GitHub will return the response in JSON format
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// fmt.Println(req)
	// Perform the POST request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	// fmt.Println(resp)

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Extract the access token from the response
	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Println(tokenResponse)
	// if tokenResponse.AccessToken != "" {
	// 	fmt.Printf("Access Token: %s\n", tokenResponse.AccessToken)
	// }
	// Here, you could use the access token to perform API requests on behalf of the user
	username,_,_ := FetchUserInformation(tokenResponse.AccessToken)
	CheckScopes(tokenResponse.AccessToken, []string{"user:email"})
	email, _,_ := FetchPrivateEmails(tokenResponse.AccessToken)
	fmt.Println(email)
	fmt.Println(username)
	// Send a response back to the client indicating success
	cookie, err := m.UserLoginGithubAuth(username, email, m.GithubPass, 2)
	if err != nil {
		m.LoginError2 = true
	} else {
		m.LoginError2 = false
	}
	http.SetCookie(w, cookie)
	jsonResponse := map[string]string{"status": "received", "access_token": tokenResponse.AccessToken}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonResponse)
	http.Redirect(w, req, "/", http.StatusOK)
}

// FetchUserInformation fetches user information from GitHub using the access token
func FetchUserInformation(accessToken string) (string, map[string]interface{}, error) {
	// Create the request to GitHub User API
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return "", nil, err
	}
	// Include the access token in the Authorization header
	req.Header.Add("Authorization", fmt.Sprintf("token %s", accessToken))

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}

	// Unmarshal the JSON response into a map
	var userInfo map[string]interface{}
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return "", nil, err
	}
	fmt.Print("Received Github Auth for user: ")
	fmt.Println(userInfo)
	fmt.Print(userInfo["login"])
	fmt.Print(" with email: ")
	// fmt.Println(userInfo)
	// fmt.Println(userInfo["login"])
	return fmt.Sprint(userInfo["login"]), userInfo, nil
}

// CheckScopes checks if the required scope is present
func CheckScopes(scope string, requiredScopes []string) bool {
	scopes := strings.Split(scope, ",")
	for _, requiredScope := range requiredScopes {
		for _, scope := range scopes {
			if scope == requiredScope {
				return true
			}
		}
	}
	return false
}

// FetchPrivateEmails fetches private emails from GitHub using the access token
func FetchPrivateEmails(accessToken string) (string, []map[string]interface{}, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	if err != nil {
		return "", nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("token %s", accessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}

	var emails []map[string]interface{}
	if err := json.Unmarshal(body, &emails); err != nil {
		return "", nil, err
	}
	fmt.Println(emails[0]["email"])

	return ("test@gmail.com"), emails, nil
}
