package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var errInvalidValue = errors.New("value must be greater than zero")

// HandlerPackageManager handles HTTP requests for calculating optimal packages usage
type HandlerPackageManager struct {
	packageManager ServicePackageManger
}

func NewHandlerPackageManager(s ServicePackageManger) *HandlerPackageManager {
	return &HandlerPackageManager{
		packageManager: s,
	}
}

type calculatePackagesReq struct {
	Small  int `json:"small"`
	Medium int `json:"medium"`
	Large  int `json:"large"`
}

func (r *calculatePackagesReq) validate() error {
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

type calculatePackagesResp struct{}

// calculatePackages calculate optimum number of packages.
func (h *HandlerPackageManager) calculatePackages(w http.ResponseWriter, r *http.Request) {
	var req calculatePackagesReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := req.validate(); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, _, _, err := h.packageManager.CalculateOptimumPackagesNumber(r.Context(), req.Small, req.Medium, req.Large)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(""))
}
