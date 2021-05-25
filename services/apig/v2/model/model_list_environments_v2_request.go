package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListEnvironmentsV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 环境名称

	Name *string `json:"name,omitempty"`
	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEnvironmentsV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEnvironmentsV2Request struct{}"
	}

	return strings.Join([]string{"ListEnvironmentsV2Request", string(data)}, " ")
}
