package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableDnssecConfigRequest Request Object
type EnableDnssecConfigRequest struct {

	// 公网域名的ID。
	ZoneId string `json:"zone_id"`
}

func (o EnableDnssecConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDnssecConfigRequest struct{}"
	}

	return strings.Join([]string{"EnableDnssecConfigRequest", string(data)}, " ")
}
