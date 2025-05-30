package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var errInvalidValue = errors.New("value must be greater than zero")

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
func calculatePackages(w http.ResponseWriter, r *http.Request) {
	var req calculatePackagesReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := req.validate(); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

}
