package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePrivateZoneStatusRequest Request Object
type UpdatePrivateZoneStatusRequest struct {

	// 域名ID。
	ZoneId string `json:"zone_id"`

	Body *UpdateZoneStatusRequestBody `json:"body,omitempty"`
}

func (o UpdatePrivateZoneStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateZoneStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrivateZoneStatusRequest", string(data)}, " ")
}
