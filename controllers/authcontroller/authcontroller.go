package authcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/shdiqq/task-5-pbi-btpns-Shadiq/config"
	"github.com/shdiqq/task-5-pbi-btpns-Shadiq/helpers"
	"github.com/shdiqq/task-5-pbi-btpns-Shadiq/models/entity"
	"github.com/shdiqq/task-5-pbi-btpns-Shadiq/models/validation"
)

type LoginSuccess struct {
	token string
}

func Login(response http.ResponseWriter, request *http.Request) {
	var login validation.Login

	// Take input from json
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&login); err != nil {
		helpers.Response(response, http.StatusBadRequest, err.Error(), nil)
		return
	}
	defer request.Body.Close()

	validate := validator.New()
	if err := validate.Struct(login); err != nil {
		helpers.Response(response, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var user entity.User
	// Username Check to database
	if err := config.DB.First(&user, "username = ?", login.Username).Error; err != nil {
		helpers.Response(response, http.StatusNotFound, "Wrong username", nil)
		return
	}

	// Password Check Validation
	if err := helpers.VerifyPassword(user.Password, login.Password); err != nil {
		helpers.Response(response, http.StatusNotFound, "Wrong password", nil)
		return
	}

	// Create token
	token, err := helpers.CreateToken(&user)
	if err != nil {
		helpers.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Add token to cookie
	http.SetCookie(response, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	data := map[string]interface{}{"token": token}

	helpers.Response(response, http.StatusOK, "Success Login", data)
}

func Register(response http.ResponseWriter, request *http.Request) {
	var register validation.Register

	// Take input from json
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&register); err != nil {
		helpers.Response(response, http.StatusBadRequest, err.Error(), nil)
		return
	}
	defer request.Body.Close()

	validate := validator.New()
	if err := validate.Struct(register); err != nil {
		helpers.Response(response, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Check password and passwordConfirm
	if register.Password != register.PasswordConfirm {
		helpers.Response(response, http.StatusBadRequest, "Password Not Match", nil)
		return
	}

	// Hash password using bcrypt
	passwordHash, err := helpers.HashPassword(register.Password)
	if err != nil {
		helpers.Response(response, 500, err.Error(), nil)
		return
	}

	user := entity.User{
		Username: register.Username,
		Email:    register.Email,
		Password: passwordHash,
	}

	// Insert to database
	if err := config.DB.Create(&user).Error; err != nil {
		helpers.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.Response(response, http.StatusCreated, "Success Register", user)
}

func Logout(response http.ResponseWriter, request *http.Request) {
	// Delete token in cookie
	http.SetCookie(response, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	helpers.Response(response, http.StatusOK, "Success Logout", nil)
}
