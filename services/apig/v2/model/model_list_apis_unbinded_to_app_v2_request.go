package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListApisUnbindedToAppV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 应用id

	AppId string `json:"app_id"`
	// 环境id

	EnvId string `json:"env_id"`
	// API分组编号

	GroupId *string `json:"group_id,omitempty"`
	// API编号

	ApiId *string `json:"api_id,omitempty"`
	// API名称

	ApiName *string `json:"api_name,omitempty"`
	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListApisUnbindedToAppV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisUnbindedToAppV2Request struct{}"
	}

	return strings.Join([]string{"ListApisUnbindedToAppV2Request", string(data)}, " ")
}
