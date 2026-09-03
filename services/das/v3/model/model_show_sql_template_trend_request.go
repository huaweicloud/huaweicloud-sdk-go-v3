package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlTemplateTrendRequest Request Object
type ShowSqlTemplateTrendRequest struct {

	// 实例ID，实例的唯一标识
	InstanceId string `json:"instance_id"`

	// 节点ID，实例节点的唯一标识
	NodeId *string `json:"node_id,omitempty"`

	// 开始时间，Unix timestamp，单位：毫秒
	StartAt int64 `json:"start_at"`

	// 结束时间，Unix timestamp，单位：毫秒
	EndAt int64 `json:"end_at"`

	// 聚合毫秒数
	IntervalMillis *int64 `json:"interval_millis,omitempty"`
}

func (o ShowSqlTemplateTrendRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlTemplateTrendRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlTemplateTrendRequest", string(data)}, " ")
}
