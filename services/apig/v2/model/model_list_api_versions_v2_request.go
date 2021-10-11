package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListApiVersionsV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// API的编号

	ApiId string `json:"api_id"`
	// 环境的编号

	EnvId *string `json:"env_id,omitempty"`
	// 环境的名称

	EnvName *string `json:"env_name,omitempty"`
	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListApiVersionsV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApiVersionsV2Request struct{}"
	}

	return strings.Join([]string{"ListApiVersionsV2Request", string(data)}, " ")
}
