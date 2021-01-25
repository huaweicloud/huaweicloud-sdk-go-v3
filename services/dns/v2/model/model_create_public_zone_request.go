package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePublicZoneRequest struct {
	Body *CreatePublicZoneReq `json:"body,omitempty"`
}

func (o CreatePublicZoneRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePublicZoneRequest struct{}"
	}

	return strings.Join([]string{"CreatePublicZoneRequest", string(data)}, " ")
}
