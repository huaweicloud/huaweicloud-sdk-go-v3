package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMfaDevicesV5Request Request Object
type ListMfaDevicesV5Request struct {

	// IAM用户ID。
	UserId *string `json:"user_id,omitempty"`

	// 每页显示的条目数量，范围为1到200条，默认为100条。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记，长度为4到400个字符，只包含字母、数字、\"+\"、\"/\"、\"=\"、\"-\"和\"_\"的字符串。
	Marker *string `json:"marker,omitempty"`
}

func (o ListMfaDevicesV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMfaDevicesV5Request struct{}"
	}

	return strings.Join([]string{"ListMfaDevicesV5Request", string(data)}, " ")
}
