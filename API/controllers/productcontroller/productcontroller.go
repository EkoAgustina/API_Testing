package productcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/jeypc/go-jwt-mux/models"

	"github.com/jeypc/go-jwt-mux/helper"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var userInput models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{
			"responseCode":    "400",
			"responseMessage": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{
			"responseCode":    "500",
			"responseMessage": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{
		"responseCode":    "200",
		"responseMessage": "success"}

	helper.ResponseJSON(w, http.StatusOK, response)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {

	var product []models.Product
	result := models.DB.Find(&product)
	json.Marshal(result)

	response := map[string]interface{}{
		"data":            product,
		"responseCode":    "200",
		"responseMessage": "success"}

	helper.ResponseJSON(w, http.StatusOK, response)
}
