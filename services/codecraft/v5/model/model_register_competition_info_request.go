package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RegisterCompetitionInfoRequest struct {
	Body *RegisterInfoRequestModel `json:"body,omitempty"`
}

func (o RegisterCompetitionInfoRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RegisterCompetitionInfoRequest struct{}"
	}

	return strings.Join([]string{"RegisterCompetitionInfoRequest", string(data)}, " ")
}
