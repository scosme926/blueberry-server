package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/scosme926/blueberry-server/utils"
	"github.com/scosme926/blueberry-server/utils/models"
	cache "github.com/patrickmn/go-cache"
)

func (c *Controller) postRegister(w http.ResponseWriter, r *http.Request) {
	var requestData models.RegisterRequest
	data := r.Body
	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if requestData.Email == "" || requestData.Password == "" {
		http.Error(w, "Fields are not properly formatted", http.StatusBadRequest)
		return
	}

	hashedPassword := utils.GenerateHashedPassword(w, r, requestData.Password)
	c.cache.Set("email", requestData.Email, cache.NoExpiration)
	c.cache.Set("hashed_password", string(hashedPassword), cache.NoExpiration)

	var responseData = &models.RegisterResponse{
		Message: "You have successfully registered. Please login to continue!",
	}
	err = json.NewEncoder(w).Encode(&responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Controller) postLogin(w http.ResponseWriter, r *http.Request) {
	var (
		hashedPassword string
		email          string
	)
	data := r.Body

	var requestData models.LoginRequest

	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if requestData.Email == "" || requestData.Password == "" {
		http.Error(w, "Fields are not properly formatted", http.StatusBadRequest)
		return
	}

	if cachedPasskey, found := c.cache.Get("hashed_password"); found {
		hashedPassword = cachedPasskey.(string)
	}
	cachedEmail, found := c.cache.Get("email")
	if found {
		email = cachedEmail.(string)
	}

	if requestData.Email != email {
		http.Error(w, "This user does not match our records", http.StatusBadRequest)
		return
	}

	if utils.CompareHashedPassword(w, r, []byte(hashedPassword), []byte(requestData.Password)) == false {
		http.Error(w, "The password you entered is incorrect", http.StatusBadRequest)
		return
	}

	var responseData = &models.LoginResponse{
		Message: "Success! You're Logging in",
	}
	err = json.NewEncoder(w).Encode(&responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}