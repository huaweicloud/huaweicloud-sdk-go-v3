package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePrivateZoneRequest struct {
	ZoneId string `json:"zone_id"`

	Body *UpdatePrivateZoneInfoReq `json:"body,omitempty"`
}

func (o UpdatePrivateZoneRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePrivateZoneRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrivateZoneRequest", string(data)}, " ")
}
