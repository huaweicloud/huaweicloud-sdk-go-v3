package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeletePrivateZoneRequest struct {
	ZoneId string `json:"zone_id"`
}

func (o DeletePrivateZoneRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePrivateZoneRequest struct{}"
	}

	return strings.Join([]string{"DeletePrivateZoneRequest", string(data)}, " ")
}
