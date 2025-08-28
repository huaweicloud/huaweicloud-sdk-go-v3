package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyVersionsV5Request Request Object
type ListPolicyVersionsV5Request struct {

	// 身份策略ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	PolicyId string `json:"policy_id"`

	// 每页显示的条目数量，范围为1到200条，默认为100条。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记，长度为4到400个字符，只包含字母、数字、\"+\"、\"/\"、\"=\"、\"-\"和\"_\"的字符串。
	Marker *string `json:"marker,omitempty"`
}

func (o ListPolicyVersionsV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyVersionsV5Request struct{}"
	}

	return strings.Join([]string{"ListPolicyVersionsV5Request", string(data)}, " ")
}
