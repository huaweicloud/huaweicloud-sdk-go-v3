package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDnssecConfigRequest Request Object
type ShowDnssecConfigRequest struct {

	// 公网域名的ID。
	ZoneId string `json:"zone_id"`
}

func (o ShowDnssecConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDnssecConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowDnssecConfigRequest", string(data)}, " ")
}
