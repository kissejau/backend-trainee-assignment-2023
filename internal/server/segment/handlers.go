package segment

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/kissejau/backend-trainee-assignment-2023/internal/server/handlers"
	"github.com/kissejau/backend-trainee-assignment-2023/pkg/response"
)

const (
	segmentUrl  = "/api/v1/segment"
	segmentsUrl = "/api/v1/segments"
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
	r.HandlerFunc("POST", segmentUrl, h.CreateSegment)
	r.HandlerFunc("GET", segmentUrl, h.GetSegmentBySlug)
	r.HandlerFunc("GET", segmentsUrl, h.GetSegments)
	r.HandlerFunc("PATCH", segmentUrl, h.UpdateSegment)
	r.HandlerFunc("DELETE", segmentUrl, h.DeleteSegment)
}

func (h *handler) CreateSegment(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	var segment Segment
	err = json.Unmarshal(data, &segment)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	err = h.r.Create(segment)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	response.Respond(w, http.StatusAccepted, data)
}

func (h *handler) GetSegmentBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.Header.Get("slug")
	if len(slug) == 0 {
		response.Respond(w, http.StatusBadRequest, []byte("incorrect header slug"))
		return
	}

	segment, err := h.r.Get(slug)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	data, err := json.Marshal(segment)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	response.Respond(w, http.StatusAccepted, data)
}

func (h *handler) GetSegments(w http.ResponseWriter, r *http.Request) {
	segments, err := h.r.List()
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

func (h *handler) UpdateSegment(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	var segment Segment
	err = json.Unmarshal(data, &segment)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	err = h.r.Update(segment)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}
	response.Respond(w, http.StatusAccepted, []byte("segment was updated"))
}

func (h *handler) DeleteSegment(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")
	_, err := strconv.Atoi(id)
	if len(id) == 0 || err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	err = h.r.Delete(id)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	response.Respond(w, http.StatusAccepted, []byte(fmt.Sprintf("segment with id=%v was delted", id)))
}
