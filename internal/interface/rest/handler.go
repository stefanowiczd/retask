package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var errInvalidValue = errors.New("value must be greater than zero")

// HandlerPacksManager handles HTTP requests for calculating optimal Packss usage
type HandlerPacksManager struct {
	PacksManager ServicePacksManger
}

func NewHandlerPacksManager(s ServicePacksManger) *HandlerPacksManager {
	return &HandlerPacksManager{
		PacksManager: s,
	}
}

type calculatePacksReq struct {
	AmountPacks int `json:"amountPacks"`
	Small       int `json:"small"`
	Medium      int `json:"medium"`
	Large       int `json:"large"`
}

func (r *calculatePacksReq) validate() error {
	if r.AmountPacks < 0 {
		return fmt.Errorf("invalid param: amountPacks: %w", errInvalidValue)
	}

	if r.Small < 0 {
		return fmt.Errorf("checking param: small: %w", errInvalidValue)
	}

	if r.Medium < 0 {
		return fmt.Errorf("checking param: medium: %w", errInvalidValue)
	}

	if r.Large < 0 {
		return fmt.Errorf("checking param: large: %w", errInvalidValue)
	}

	return nil
}

type calculatePacksResp struct{}

// calculatePackss calculate optimum number of Packss.
func (h *HandlerPacksManager) calculatePacks(w http.ResponseWriter, r *http.Request) {
	var req calculatePacksReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := req.validate(); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, _, _, err := h.PacksManager.CalculateOptimumPacksAmount(r.Context(), req.Small, req.Medium, req.Large)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(""))
}
