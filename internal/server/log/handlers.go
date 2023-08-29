package log

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/kissejau/backend-trainee-assignment-2023/internal/server/errors"
	"github.com/kissejau/backend-trainee-assignment-2023/internal/server/handlers"
	"github.com/kissejau/backend-trainee-assignment-2023/pkg/response"
)

const (
	LogsUrl = "/api/v1/logs"
	LogUrl  = "/api/v1/log"
)

type handler struct {
	r Repository
}

func NewHandler(r Repository) handlers.Handler {
	return &handler{
		r: r,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc("GET", LogsUrl, h.GetLogs)
	router.HandlerFunc("GET", LogUrl, h.GetLog)
}

func (h *handler) GetLog(w http.ResponseWriter, r *http.Request) {
	logName := r.URL.Query().Get("log")
	if len(logName) == 0 {
		response.Respond(w, http.StatusBadRequest, []byte(errors.ErrInvalidParams.Error()))
		return
	}

	file, err := GetLogFile(logName)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(errors.ErrIncorrectParams.Error()))
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=log.csv")

	_, err = io.Copy(w, file)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}
}

func (h *handler) GetLogs(w http.ResponseWriter, r *http.Request) {
	var dateDTO DateDTO
	err := json.NewDecoder(r.Body).Decode(&dateDTO)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(errors.ErrReadingBody.Error()))
		return
	}

	_, err = strconv.Atoi(dateDTO.Year)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(errors.ErrInvalidBody.Error()))
		return
	}

	_, err = strconv.Atoi(dateDTO.Month)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(errors.ErrInvalidBody.Error()))
		return
	}

	date, err := time.Parse("2006-01-02", fmt.Sprintf("%s-%s-01", dateDTO.Year, dateDTO.Month))
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}
	fmt.Println(date.String())

	logs, err := h.r.GetLogs(date, date.AddDate(0, 1, 0))
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	link, err := GenerateLogLink(logs)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	data, err := json.Marshal(link)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}
	response.Respond(w, http.StatusAccepted, data)
}
