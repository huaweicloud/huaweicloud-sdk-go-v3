package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowPrivateZoneRequest struct {
	ZoneId string `json:"zone_id"`
}

func (o ShowPrivateZoneRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPrivateZoneRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateZoneRequest", string(data)}, " ")
}
