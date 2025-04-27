package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePublicZoneStatusRequest Request Object
type UpdatePublicZoneStatusRequest struct {

	// 域名ID。
	ZoneId string `json:"zone_id"`

	Body *UpdateZoneStatusRequestBody `json:"body,omitempty"`
}

func (o UpdatePublicZoneStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicZoneStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdatePublicZoneStatusRequest", string(data)}, " ")
}
