package usercontroller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/shdiqq/task-5-pbi-btpns-Shadiq/config"
	"github.com/shdiqq/task-5-pbi-btpns-Shadiq/helpers"
	"github.com/shdiqq/task-5-pbi-btpns-Shadiq/models/entity"
	"github.com/shdiqq/task-5-pbi-btpns-Shadiq/models/validation"
	"gorm.io/gorm"
)

func ListUser(response http.ResponseWriter, request *http.Request) {
	var user []entity.User

	// Select into database
	if err := config.DB.Find(&user).Error; err != nil {
		helpers.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.Response(response, http.StatusOK, "Success Get List User", user)
}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	// Get params
	vars := mux.Vars(request)
	id, err := strconv.ParseInt(vars["userId"], 10, 64)
	if err != nil {
		helpers.Response(response, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Check user in database
	var user entity.User
	if err := config.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(response, http.StatusNotFound, "User Not Found", nil)
			return
		}

		helpers.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Take input from json
	var userRequest validation.UserRequestPut
	if err := json.NewDecoder(request.Body).Decode(&userRequest); err != nil {
		helpers.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	defer request.Body.Close()

	validate := validator.New()
	if err := validate.Struct(userRequest); err != nil {
		helpers.Response(response, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if userRequest.Username != "" {
		user.Username = userRequest.Username
	}
	if userRequest.Email != "" {
		user.Email = userRequest.Email
	}
	if userRequest.Password != "" {
		passwordHash, err := helpers.HashPassword(userRequest.Password)
		if err != nil {
			helpers.Response(response, 500, err.Error(), nil)
			return
		}
		user.Password = passwordHash
	}

	// Update user in database
	if err := config.DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		helpers.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.Response(response, http.StatusCreated, "Success Update User", user)
}

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	// Get params
	vars := mux.Vars(request)
	id, err := strconv.ParseInt(vars["userId"], 10, 64)
	if err != nil {
		helpers.Response(response, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Check user in database
	var user entity.User
	if err := config.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(response, http.StatusNotFound, "User Not Found", nil)
			return
		}

		helpers.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Delete Photos Associated with The User
	var photo entity.Photo
	config.DB.Where("user_id = ?", id).Delete(&photo)

	// Delete user in database
	res := config.DB.Delete(&user, id)
	if res.Error != nil {
		helpers.Response(response, http.StatusInternalServerError, res.Error.Error(), nil)
		return
	}
	if res.RowsAffected == 0 {
		helpers.Response(response, http.StatusNotFound, "User Not Found", nil)
		return
	}

	helpers.Response(response, http.StatusOK, "Success Delete User", user)
}
