package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOidcProvidersV5Request Request Object
type ListOidcProvidersV5Request struct {

	// 每页显示的条目数量，范围为1到200条，默认为100条。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记，长度为4到400个字符，只包含字母、数字、\"+\"、\"/\"、\"=\"、\"-\"和\"_\"的字符串。
	Marker *string `json:"marker,omitempty"`
}

func (o ListOidcProvidersV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOidcProvidersV5Request struct{}"
	}

	return strings.Join([]string{"ListOidcProvidersV5Request", string(data)}, " ")
}
