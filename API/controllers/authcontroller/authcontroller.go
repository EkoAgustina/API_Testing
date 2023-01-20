package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/jeypc/go-jwt-mux/config"

	"github.com/jeypc/go-jwt-mux/helper"
	"gorm.io/gorm"

	"github.com/jeypc/go-jwt-mux/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	/*
		get user data by username
	*/
	var user models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{
				"responseCode":    "401",
				"responseMessage": "Incorrect username or password"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{
				"responseCode": "500", "message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	/*
		check if the password is valid
	*/
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"responseCode": "401", "responseMessage": "Incorrect username or password"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	/*
		JWT token generation
	*/
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	/*
		Declares the algorithm to use for signing
	*/
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	/*
		Set token on Cookie
	*/
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{
		"responseCode":    "200",
		"responseMessage": "Login successful",
		"token":           token}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {

	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{
			"responseCode":    "400",
			"responseMessage": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	/*
		hash password using bcrypt
	*/
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	/*
		Insert database
	*/
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{
			"responseCode":    "500",
			"responseMessage": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{
		"data": map[string]interface{}{
			"nama_lengkap": userInput.NamaLengkap},
		"responseCode":    "200",
		"responseMessage": "success"}

	helper.ResponseJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	/*
		Delete token that is in
	*/
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{
		"responseCode":    "200",
		"responseMessage": "logout successful"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
