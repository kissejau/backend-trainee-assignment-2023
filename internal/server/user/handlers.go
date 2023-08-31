package user

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/kissejau/backend-trainee-assignment-2023/internal/server/errors"
	"github.com/kissejau/backend-trainee-assignment-2023/internal/server/handlers"
	"github.com/kissejau/backend-trainee-assignment-2023/pkg/response"
)

const (
	userUrl     = "/api/v1/user"
	usersUrl    = "/api/v1/users"
	segmentsUrl = "/segments"
)

type handler struct {
	r Repository
}

func NewHandler(r Repository) handlers.Handler {
	return &handler{
		r: r,
	}
}

func (h *handler) Register(r *httprouter.Router) {
	r.HandlerFunc("POST", userUrl, h.CreateUser)
	r.HandlerFunc("GET", userUrl, h.GetUserById)
	r.HandlerFunc("GET", usersUrl, h.GetUsers)
	r.HandlerFunc("PATCH", userUrl, h.UpdateUser)
	r.HandlerFunc("DELETE", userUrl, h.DeleteUser)

	r.HandlerFunc("GET", userUrl+segmentsUrl, h.GetUserSegments)
	r.HandlerFunc("POST", userUrl+segmentsUrl, h.SetUserSegments)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(errors.ErrReadingBody.Error()))
		return
	}

	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(errors.ErrInvalidBody.Error()))
		return
	}

	err = h.r.Create(user.Name)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	response.Respond(w, http.StatusAccepted, data)
}

func (h *handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")
	_, err := strconv.Atoi(id)
	if len(id) == 0 || err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(errors.ErrIncorrectHeader.Error()))
		return
	}

	user, err := h.r.Get(id)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	response.Respond(w, http.StatusAccepted, data)
}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.r.List()
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	data, err := json.Marshal(users)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	response.Respond(w, http.StatusAccepted, data)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(errors.ErrReadingBody.Error()))
		return
	}

	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(errors.ErrInvalidBody.Error()))
		return
	}

	err = h.r.Update(user)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}
	response.Respond(w, http.StatusAccepted, data)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")
	_, err := strconv.Atoi(id)
	if len(id) == 0 || err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(errors.ErrIncorrectHeader.Error()))
		return
	}

	err = h.r.Delete(id)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	response.Respond(w, http.StatusAccepted, []byte(""))
}

func (h *handler) GetUserSegments(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")
	_, err := strconv.Atoi(id)
	if len(id) == 0 || err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(errors.ErrIncorrectHeader.Error()))
		return
	}

	segments, err := h.r.GetSegments(id)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	data, err := json.Marshal(segments)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	response.Respond(w, http.StatusAccepted, data)
}

func (h *handler) SetUserSegments(w http.ResponseWriter, r *http.Request) {
	var setUserSegmentsDTO SetUserSegmentsDTO
	err := json.NewDecoder(r.Body).Decode(&setUserSegmentsDTO)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(errors.ErrReadingBody.Error()))
		return
	}

	if setUserSegmentsDTO.UserId == "" {
		response.Respond(w, http.StatusBadRequest, []byte(errors.ErrInvalidBody.Error()))
		return
	}

	sqlableSetUserSegmentsDTO := setUserSegmentsDTO.SqlableSetUserSegmentsDTO()
	log.Println("HANDLER NOTE", sqlableSetUserSegmentsDTO)
	err = h.r.SetSegments(sqlableSetUserSegmentsDTO)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	response.Respond(w, http.StatusAccepted, []byte(""))
}
