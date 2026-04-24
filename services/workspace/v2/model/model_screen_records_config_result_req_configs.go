package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScreenRecordsConfigResultReqConfigs struct {

	// 站点ID。
	SiteId *string `json:"site_id,omitempty"`

	// 录屏限速。
	TrafficLimit *int32 `json:"traffic_limit,omitempty"`
}

func (o ScreenRecordsConfigResultReqConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScreenRecordsConfigResultReqConfigs struct{}"
	}

	return strings.Join([]string{"ScreenRecordsConfigResultReqConfigs", string(data)}, " ")
}
