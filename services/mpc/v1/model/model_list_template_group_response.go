package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTemplateGroupResponse struct {

	// 模板组信息列表。
	TemplateGroupList *[]TemplateGroup `json:"template_group_list,omitempty" xml:"template_group_list"`

	// 转码模板组总数
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTemplateGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateGroupResponse struct{}"
	}

	return strings.Join([]string{"ListTemplateGroupResponse", string(data)}, " ")
}
