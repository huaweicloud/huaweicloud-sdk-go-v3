package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowLogTrendNewRequest Request Object
type ShowSlowLogTrendNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 开始时间（Unix timestamp），单位：毫秒
	StartTime int64 `json:"start_time"`

	// 结束时间（Unix timestamp），单位：毫秒
	EndTime int64 `json:"end_time"`
}

func (o ShowSlowLogTrendNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogTrendNewRequest struct{}"
	}

	return strings.Join([]string{"ShowSlowLogTrendNewRequest", string(data)}, " ")
}
