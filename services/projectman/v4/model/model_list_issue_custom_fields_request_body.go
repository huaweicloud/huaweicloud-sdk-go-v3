package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询Scrum工作项的自定义字段，custom_fields和names的相关自定义字段都返回
type ListIssueCustomFieldsRequestBody struct {

	// 自定义字段
	CustomFields *[]string `json:"custom_fields,omitempty"`

	// 自定义字段页面显示的含义
	Names *[]string `json:"names,omitempty"`
}

func (o ListIssueCustomFieldsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueCustomFieldsRequestBody struct{}"
	}

	return strings.Join([]string{"ListIssueCustomFieldsRequestBody", string(data)}, " ")
}
