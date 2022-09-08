package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 拓扑图上线条的的信息
type TraceTopologyLineInfo struct {

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 耗时
	TimeUsed *int64 `json:"time_used,omitempty"`

	// 参数信息，比如调用的url信息等
	Argument *string `json:"argument,omitempty"`

	// event的id
	EventId *string `json:"event_id,omitempty"`
}

func (o TraceTopologyLineInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TraceTopologyLineInfo struct{}"
	}

	return strings.Join([]string{"TraceTopologyLineInfo", string(data)}, " ")
}
