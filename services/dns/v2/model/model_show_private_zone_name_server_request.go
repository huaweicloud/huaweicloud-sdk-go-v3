package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowPrivateZoneNameServerRequest struct {
	// 待查询内网zone的ID。

	ZoneId string `json:"zone_id"`
}

func (o ShowPrivateZoneNameServerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPrivateZoneNameServerRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateZoneNameServerRequest", string(data)}, " ")
}
