package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAppsBindedToApiV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// API编号

	ApiId *string `json:"api_id,omitempty"`
	// APP名称

	AppName *string `json:"app_name,omitempty"`
	// APP编号

	AppId *string `json:"app_id,omitempty"`
	// 环境编号

	EnvId *string `json:"env_id,omitempty"`
	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAppsBindedToApiV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAppsBindedToApiV2Request struct{}"
	}

	return strings.Join([]string{"ListAppsBindedToApiV2Request", string(data)}, " ")
}
