package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSignatureKeysBindedToApiV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// API的编号
	ApiId string `json:"api_id"`

	// 签名密钥的编号
	SignId *string `json:"sign_id,omitempty"`

	// 签名密钥的名称
	SignName *string `json:"sign_name,omitempty"`

	// 环境编号
	EnvId *string `json:"env_id,omitempty"`
}

func (o ListSignatureKeysBindedToApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSignatureKeysBindedToApiV2Request struct{}"
	}

	return strings.Join([]string{"ListSignatureKeysBindedToApiV2Request", string(data)}, " ")
}
