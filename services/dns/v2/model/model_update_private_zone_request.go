package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePrivateZoneRequest Request Object
type UpdatePrivateZoneRequest struct {

	// 域名ID。
	ZoneId string `json:"zone_id"`

	Body *UpdatePrivateZoneInfoReq `json:"body,omitempty"`
}

func (o UpdatePrivateZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateZoneRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrivateZoneRequest", string(data)}, " ")
}
