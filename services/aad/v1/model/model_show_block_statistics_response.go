package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBlockStatisticsResponse Response Object
type ShowBlockStatisticsResponse struct {

	// 总解封次数
	TotalUnblockingTimes *int32 `json:"total_unblocking_times,omitempty"`

	// 人工解封次数
	ManualUnblockingTimes *int32 `json:"manual_unblocking_times,omitempty"`

	// 自动解封次数
	AutomaticUnblockingTimes *int32 `json:"automatic_unblocking_times,omitempty"`

	// 当前封堵Ip数
	CurrentBlockedIpNumbers *int32 `json:"current_blocked_ip_numbers,omitempty"`
	HttpStatusCode          int    `json:"-"`
}

func (o ShowBlockStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlockStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowBlockStatisticsResponse", string(data)}, " ")
}
