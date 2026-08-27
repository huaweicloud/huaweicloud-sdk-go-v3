package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateScreenRecordsTrafficLimitConfigRequestBodyConfigs struct {

	// 站点ID。
	SiteId *string `json:"site_id,omitempty"`

	// 录屏限速。
	TrafficLimit *int32 `json:"traffic_limit,omitempty"`
}

func (o UpdateScreenRecordsTrafficLimitConfigRequestBodyConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScreenRecordsTrafficLimitConfigRequestBodyConfigs struct{}"
	}

	return strings.Join([]string{"UpdateScreenRecordsTrafficLimitConfigRequestBodyConfigs", string(data)}, " ")
}
