package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIssueCustomFieldsRequestBody 查询Scrum工作项的自定义字段，custom_fields和names的相关自定义字段都返回
type ListIssueCustomFieldsRequestBody struct {

	// 自定义字段
	CustomFields *[]string `json:"custom_fields,omitempty"`

	// 自定义字段页面显示的含义
	Names *[]string `json:"names,omitempty"`

	// 查询结果是否包括未使用的自定义字段，默认仅查询使用中的自定义字段，设为true时查询项目中所有自定义字段
	IncludedNotInUse *bool `json:"included_not_in_use,omitempty"`
}

func (o ListIssueCustomFieldsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueCustomFieldsRequestBody struct{}"
	}

	return strings.Join([]string{"ListIssueCustomFieldsRequestBody", string(data)}, " ")
}
