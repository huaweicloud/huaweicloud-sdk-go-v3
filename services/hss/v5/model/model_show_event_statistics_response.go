package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEventStatisticsResponse Response Object
type ShowEventStatisticsResponse struct {

	// **参数解释**: 存在告警的服务器数量 **取值范围**: 最小值0，最大值9223372036854775807
	RiskyHostNum *int64 `json:"risky_host_num,omitempty"`

	// **参数解释**: 存在告警的容器数量 **取值范围**: 最小值0，最大值9223372036854775807
	RiskyContainerNum *int64 `json:"risky_container_num,omitempty"`

	// **参数解释**: 待处理告警事件数量 **取值范围**: 最小值0，最大值9223372036854775807
	UnhandledNum *int64 `json:"unhandled_num,omitempty"`

	// **参数解释**: 已处理告警事件数量 **取值范围**: 最小值0，最大值9223372036854775807
	ProcessedNum *int64 `json:"processed_num,omitempty"`

	// **参数解释**: 自动处置告警事件数量 **取值范围**: 最小值0，最大值9223372036854775807
	AutoBlockedNum *int64 `json:"auto_blocked_num,omitempty"`

	// **参数解释**: 已拦截IP数量 **取值范围**: 最小值0，最大值9223372036854775807
	BlockedIpNum *int64 `json:"blocked_ip_num,omitempty"`

	// **参数解释**: 已隔离文件数量 **取值范围**: 最小值0，最大值9223372036854775807
	IsolatedFileNum *int64 `json:"isolated_file_num,omitempty"`

	// **参数解释**: 是否支持ES查询 **取值范围**: - true：支持 - false：不支持
	SupportES      *bool `json:"support_ES,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowEventStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowEventStatisticsResponse", string(data)}, " ")
}
