package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublicZoneNameServerRequest Request Object
type ShowPublicZoneNameServerRequest struct {

	// 域名ID。
	ZoneId string `json:"zone_id"`
}

func (o ShowPublicZoneNameServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicZoneNameServerRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicZoneNameServerRequest", string(data)}, " ")
}
