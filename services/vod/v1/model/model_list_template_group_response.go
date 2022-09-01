package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTemplateGroupResponse struct {

	// 模板组信息<br/>
	TemplateGroupList *[]TemplateGroup `json:"template_group_list,omitempty" xml:"template_group_list"`

	// 总记录条数<br/>
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 错误码<br/>
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误描述<br/>
	ErrorMsg       *string `json:"error_msg,omitempty" xml:"error_msg"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTemplateGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateGroupResponse struct{}"
	}

	return strings.Join([]string{"ListTemplateGroupResponse", string(data)}, " ")
}
