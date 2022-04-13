package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBuiltInPolicyDefinitionsRequest struct {
	// 选择接口返回的信息的语言，默认为\"zh-cn\"中文

	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListBuiltInPolicyDefinitionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBuiltInPolicyDefinitionsRequest struct{}"
	}

	return strings.Join([]string{"ListBuiltInPolicyDefinitionsRequest", string(data)}, " ")
}
