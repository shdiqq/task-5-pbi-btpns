package photocontroller

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

func createPhotoResponse(photo entity.Photo) entity.PhotoResponse {
	return entity.PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
	}
}

func ListPhoto(response http.ResponseWriter, request *http.Request) {
	var photo []entity.Photo
	userInfo := request.Context().Value("userinfo").(*helpers.JWTClaims)

	// Select into database
	if err := config.DB.Where("user_id = ?", userInfo.ID).Find(&photo).Error; err != nil {
		helpers.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var photoResponses []entity.PhotoResponse

	for _, item := range photo {
		photoResponse := createPhotoResponse(item)
		photoResponses = append(photoResponses, photoResponse)
	}

	if photoResponses == nil {
		data := []string{}
		helpers.Response(response, http.StatusOK, "Success Get List Photo User", data)
		return
	}

	helpers.Response(response, http.StatusOK, "Success Get List Photo User", photoResponses)
}

func CreatePhoto(response http.ResponseWriter, request *http.Request) {
	var PhotoRequest validation.PhotoRequestPost

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&PhotoRequest); err != nil {
		helpers.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	defer request.Body.Close()

	validate := validator.New()
	if err := validate.Struct(PhotoRequest); err != nil {
		helpers.Response(response, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Check user
	userInfo := request.Context().Value("userinfo").(*helpers.JWTClaims)
	var user entity.User
	if err := config.DB.First(&user, userInfo.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(response, http.StatusNotFound, "User Not Found", nil)
			return
		}

		helpers.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	photo := entity.Photo{
		Title:    PhotoRequest.Title,
		Caption:  PhotoRequest.Caption,
		PhotoUrl: PhotoRequest.PhotoUrl,
		UserID:   userInfo.ID,
	}

	if err := config.DB.Create(&photo).Error; err != nil {
		helpers.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	photoResponses := entity.PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
	}

	helpers.Response(response, http.StatusCreated, "Success Create Photo", photoResponses)
}

func ShowDetailPhoto(response http.ResponseWriter, request *http.Request) {
	// Get params
	vars := mux.Vars(request)
	id, err := strconv.ParseInt(vars["photoId"], 10, 64)
	if err != nil {
		helpers.Response(response, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var photo entity.Photo

	if err := config.DB.Joins("User").First(&photo, id).First(&photo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(response, http.StatusNotFound, "Photo Not Found", nil)
			return
		}

		helpers.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.Response(response, http.StatusOK, "Success Get Detail Photo", photo)
}

func UpdatePhoto(response http.ResponseWriter, request *http.Request) {
	// Get params
	vars := mux.Vars(request)
	id, err := strconv.ParseInt(vars["photoId"], 10, 64)
	if err != nil {
		helpers.Response(response, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Check photoId in database
	var photo entity.Photo
	if err := config.DB.First(&photo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(response, http.StatusNotFound, "Photo Not Found", nil)
			return
		}

		helpers.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Check user access
	var user entity.User
	userInfo := request.Context().Value("userinfo").(*helpers.JWTClaims)
	if err := config.DB.First(&user, userInfo.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(response, http.StatusNotFound, "User Not Found", nil)
			return
		}

		helpers.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if photo.UserID != userInfo.ID {
		helpers.Response(response, http.StatusUnauthorized, "User Can't Access Edit This Photo ", nil)
		return
	}

	// Take input from json
	var photoRequest validation.PhotoRequestPut
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&photoRequest); err != nil {
		helpers.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	defer request.Body.Close()

	validate := validator.New()
	if err := validate.Struct(photoRequest); err != nil {
		helpers.Response(response, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if photoRequest.Title != "" {
		photo.Title = photoRequest.Title
	}
	if photoRequest.Caption != "" {
		photo.Caption = photoRequest.Caption
	}
	if photoRequest.PhotoUrl != "" {
		photo.PhotoUrl = photoRequest.PhotoUrl
	}

	if err := config.DB.Where("id = ?", id).Updates(&photo).Error; err != nil {
		helpers.Response(response, http.StatusBadRequest, err.Error(), nil)
		return
	}

	photoResponses := entity.PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
	}

	helpers.Response(response, http.StatusOK, "Success Update Photo", photoResponses)
}

func DeletePhoto(response http.ResponseWriter, request *http.Request) {
	// Get params
	vars := mux.Vars(request)
	id, err := strconv.ParseInt(vars["photoId"], 10, 64)
	if err != nil {
		helpers.Response(response, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Check photoId in database
	var photo entity.Photo
	if err := config.DB.First(&photo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(response, http.StatusNotFound, "Photo Not Found", nil)
			return
		}

		helpers.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Check User access
	var user entity.User
	userInfo := request.Context().Value("userinfo").(*helpers.JWTClaims)
	if err := config.DB.First(&user, userInfo.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(response, http.StatusNotFound, "User Not Found", nil)
			return
		}

		helpers.Response(response, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if photo.UserID != userInfo.ID {
		helpers.Response(response, http.StatusUnauthorized, "Can't Access Delete This Photo ", nil)
		return
	}

	// Delete photo in database
	res := config.DB.Delete(&photo, id)
	if res.Error != nil {
		helpers.Response(response, http.StatusInternalServerError, res.Error.Error(), nil)
		return
	}
	if res.RowsAffected == 0 {
		helpers.Response(response, http.StatusNotFound, "Photo Not Found", nil)
		return
	}

	photoResponses := entity.PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
	}

	helpers.Response(response, http.StatusOK, "Success Delete Photo", photoResponses)
}
