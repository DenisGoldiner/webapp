package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"

	"github.com/DenisGoldiner/webapp/internal"
)

type TravellerHandler struct {
	service internal.Travellers
}

func NewTravellerHandler(service internal.Travellers) TravellerHandler {
	return TravellerHandler{
		service: service,
	}
}

func (h TravellerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetTraveller(w, r)
	case http.MethodPost:
		h.CreateTraveller(w, r)
	case http.MethodDelete:
		h.DeleteTraveller(w, r)
	default:
		msg := fmt.Sprintf("method %s is not supported", r.Method)
		slog.Info(msg)
		http.Error(w, msg, http.StatusMethodNotAllowed)
	}
}

func (h TravellerHandler) GetTraveller(w http.ResponseWriter, r *http.Request) {
	slog.Info("GetTraveller")

	ctx := r.Context()

	idParam := r.URL.Query().Get("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, fmt.Sprintf("id must be a valid uuid"), http.StatusBadRequest)
		return
	}

	res, err := h.service.GetTraveller(ctx, id)
	if errors.Is(err, internal.ErrNoResource) {
		slog.Warn("request failed", "error", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		slog.Error("request failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respTraveller := Traveller{
		ID:        res.ID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
	}

	if err = json.NewEncoder(w).Encode(respTraveller); err != nil {
		slog.Error("failed to encode response", "error", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h TravellerHandler) CreateTraveller(w http.ResponseWriter, r *http.Request) {
	slog.Info("CreateTraveller")

	if r.Body == nil {
		http.Error(w, "body must not be nil", http.StatusBadRequest)
		return
	}

	var payload CreateTravellerPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, fmt.Sprintf("failed to parse the body %v", err), http.StatusBadRequest)
		return
	}

	slog.Info("CreateTravellerPayload", payload)

	h.service.CreateTraveller(r.Context(), internal.Traveller{})

	w.WriteHeader(http.StatusOK)
}

func (h TravellerHandler) DeleteTraveller(w http.ResponseWriter, r *http.Request) {
	slog.Info("DeleteTraveller")

	h.service.DeleteTraveller()

	w.WriteHeader(http.StatusOK)
}
