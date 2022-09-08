package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRedislogRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 返回日志的类型，当前仅支持Redis运行日志，类型为run
	LogType string `json:"log_type"`
}

func (o ListRedislogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRedislogRequest struct{}"
	}

	return strings.Join([]string{"ListRedislogRequest", string(data)}, " ")
}
