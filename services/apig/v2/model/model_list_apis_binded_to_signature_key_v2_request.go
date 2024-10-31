package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApisBindedToSignatureKeyV2Request Request Object
type ListApisBindedToSignatureKeyV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
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

	// API标签，该参数可指定多个，多个不同的参数值为或关系；不指定或为空时，表示不筛选标签；指定为#no_tags#时，表示筛选无标签API。
	Tags *string `json:"tags,omitempty"`
}

func (o ListApisBindedToSignatureKeyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisBindedToSignatureKeyV2Request struct{}"
	}

	return strings.Join([]string{"ListApisBindedToSignatureKeyV2Request", string(data)}, " ")
}
