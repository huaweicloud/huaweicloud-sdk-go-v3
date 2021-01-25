package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowAgencyResponse struct {
	Agency         *AgencyResult `json:"agency,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowAgencyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAgencyResponse struct{}"
	}

	return strings.Join([]string{"ShowAgencyResponse", string(data)}, " ")
}
