package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowPublicZoneRequest struct {
	// 待查询zone的ID。

	ZoneId string `json:"zone_id"`
}

func (o ShowPublicZoneRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPublicZoneRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicZoneRequest", string(data)}, " ")
}
