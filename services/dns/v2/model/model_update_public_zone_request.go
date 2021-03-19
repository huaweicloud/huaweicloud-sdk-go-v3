package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePublicZoneRequest struct {
	ZoneId string `json:"zone_id"`

	Body *UpdatePublicZoneInfo `json:"body,omitempty"`
}

func (o UpdatePublicZoneRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePublicZoneRequest struct{}"
	}

	return strings.Join([]string{"UpdatePublicZoneRequest", string(data)}, " ")
}
