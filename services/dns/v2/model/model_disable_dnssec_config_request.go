package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableDnssecConfigRequest Request Object
type DisableDnssecConfigRequest struct {

	// 公网域名的ID。
	ZoneId string `json:"zone_id"`
}

func (o DisableDnssecConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableDnssecConfigRequest struct{}"
	}

	return strings.Join([]string{"DisableDnssecConfigRequest", string(data)}, " ")
}
