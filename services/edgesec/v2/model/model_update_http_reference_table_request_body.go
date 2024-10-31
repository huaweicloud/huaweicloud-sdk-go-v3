package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateHttpReferenceTableRequestBody struct {

	// 引用表名称
	Name string `json:"name"`

	// 引用表类型，可选值为：url、params、ip、cookie、referer、user-agent、header、response_code、response_header、response_body。
	Type string `json:"type"`

	// 引用表的值
	Values []string `json:"values"`

	// 引用表描述，最长128字符
	Description *string `json:"description,omitempty"`
}

func (o UpdateHttpReferenceTableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpReferenceTableRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHttpReferenceTableRequestBody", string(data)}, " ")
}
