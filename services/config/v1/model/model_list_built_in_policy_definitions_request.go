package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBuiltInPolicyDefinitionsRequest Request Object
type ListBuiltInPolicyDefinitionsRequest struct {

	// 选择接口返回的信息的语言，默认为\"zh-cn\"中文
	XLanguage *string `json:"X-Language,omitempty"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`
}

func (o ListBuiltInPolicyDefinitionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBuiltInPolicyDefinitionsRequest struct{}"
	}

	return strings.Join([]string{"ListBuiltInPolicyDefinitionsRequest", string(data)}, " ")
}
