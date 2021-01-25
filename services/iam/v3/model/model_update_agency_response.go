package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateAgencyResponse struct {
	Agency         *AgencyResult `json:"agency,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateAgencyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAgencyResponse struct{}"
	}

	return strings.Join([]string{"UpdateAgencyResponse", string(data)}, " ")
}
