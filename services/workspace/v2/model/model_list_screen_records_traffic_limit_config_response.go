package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScreenRecordsTrafficLimitConfigResponse Response Object
type ListScreenRecordsTrafficLimitConfigResponse struct {

	// 总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 录屏记录。
	Configs        *[]ScreenRecordsConfigResultReqConfigs `json:"configs,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ListScreenRecordsTrafficLimitConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScreenRecordsTrafficLimitConfigResponse struct{}"
	}

	return strings.Join([]string{"ListScreenRecordsTrafficLimitConfigResponse", string(data)}, " ")
}
