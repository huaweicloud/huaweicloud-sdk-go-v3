package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePublicZoneStatusRequest struct {
	// 待设置状态Zone的ID

	ZoneId string `json:"zone_id"`

	Body *UpdatePublicZoneStatus `json:"body,omitempty"`
}

func (o UpdatePublicZoneStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePublicZoneStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdatePublicZoneStatusRequest", string(data)}, " ")
}
