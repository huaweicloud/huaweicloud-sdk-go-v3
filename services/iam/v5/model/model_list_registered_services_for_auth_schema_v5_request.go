package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRegisteredServicesForAuthSchemaV5Request Request Object
type ListRegisteredServicesForAuthSchemaV5Request struct {

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记，长度为4到400个字符，只包含字母、数字、\"+\"、\"/\"、\"=\"、\"-\"和\"_\"的字符串。
	Marker *string `json:"marker,omitempty"`
}

func (o ListRegisteredServicesForAuthSchemaV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRegisteredServicesForAuthSchemaV5Request struct{}"
	}

	return strings.Join([]string{"ListRegisteredServicesForAuthSchemaV5Request", string(data)}, " ")
}
