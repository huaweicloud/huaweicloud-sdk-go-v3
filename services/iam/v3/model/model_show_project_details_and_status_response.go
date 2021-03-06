package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowProjectDetailsAndStatusResponse struct {
	Project        *ProjectDetailsAndStatusResult `json:"project,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ShowProjectDetailsAndStatusResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProjectDetailsAndStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectDetailsAndStatusResponse", string(data)}, " ")
}
