package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowRecordSetByZoneRequest struct {
	ZoneId string `json:"zone_id"`
}

func (o ShowRecordSetByZoneRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowRecordSetByZoneRequest struct{}"
	}

	return strings.Join([]string{"ShowRecordSetByZoneRequest", string(data)}, " ")
}
