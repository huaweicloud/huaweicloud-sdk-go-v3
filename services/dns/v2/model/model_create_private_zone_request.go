package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePrivateZoneRequest struct {
	Body *CreatePrivateZoneReq `json:"body,omitempty"`
}

func (o CreatePrivateZoneRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePrivateZoneRequest struct{}"
	}

	return strings.Join([]string{"CreatePrivateZoneRequest", string(data)}, " ")
}
