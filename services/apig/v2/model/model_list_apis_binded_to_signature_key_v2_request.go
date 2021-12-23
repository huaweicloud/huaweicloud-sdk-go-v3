package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListApisBindedToSignatureKeyV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
	// 签名密钥编号

	SignId string `json:"sign_id"`
	// 环境编号

	EnvId *string `json:"env_id,omitempty"`
	// API的编号

	ApiId *string `json:"api_id,omitempty"`
	// API名称

	ApiName *string `json:"api_name,omitempty"`
	// API分组编号

	GroupId *string `json:"group_id,omitempty"`
}

func (o ListApisBindedToSignatureKeyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisBindedToSignatureKeyV2Request struct{}"
	}

	return strings.Join([]string{"ListApisBindedToSignatureKeyV2Request", string(data)}, " ")
}
