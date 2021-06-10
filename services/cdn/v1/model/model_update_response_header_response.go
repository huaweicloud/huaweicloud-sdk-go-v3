package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateResponseHeaderResponse struct {
	Headers        *HeaderMap `json:"headers,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o UpdateResponseHeaderResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateResponseHeaderResponse struct{}"
	}

	return strings.Join([]string{"UpdateResponseHeaderResponse", string(data)}, " ")
}
