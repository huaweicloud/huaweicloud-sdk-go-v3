package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeletePublicZoneRequest struct {
	// 待删除zone的ID

	ZoneId string `json:"zone_id"`
}

func (o DeletePublicZoneRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePublicZoneRequest struct{}"
	}

	return strings.Join([]string{"DeletePublicZoneRequest", string(data)}, " ")
}
