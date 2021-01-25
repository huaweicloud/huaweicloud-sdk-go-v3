package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowPublicZoneNameServerRequest struct {
	ZoneId string `json:"zone_id"`
}

func (o ShowPublicZoneNameServerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPublicZoneNameServerRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicZoneNameServerRequest", string(data)}, " ")
}
