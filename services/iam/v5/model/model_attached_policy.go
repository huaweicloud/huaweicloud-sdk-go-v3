package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttachedPolicy struct {

	// 身份策略名称，长度为1到128个字符，只包含字母、数字、\"_\"、\"+\"、\"=\"、\".\"、\"@\"和\"-\"的字符串。
	PolicyName string `json:"policy_name"`

	// 身份策略ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	PolicyId string `json:"policy_id"`

	// 统一资源名称。
	Urn string `json:"urn"`

	// 身份策略的附加时间。
	AttachedAt *sdktime.SdkTime `json:"attached_at"`
}

func (o AttachedPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachedPolicy struct{}"
	}

	return strings.Join([]string{"AttachedPolicy", string(data)}, " ")
}
