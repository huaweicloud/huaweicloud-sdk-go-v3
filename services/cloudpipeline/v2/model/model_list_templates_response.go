package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTemplatesResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 页码数
	PageNumber *int32 `json:"page_number,omitempty" xml:"page_number"`

	// 每页显示数
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size"`

	// 模板数据,list类型数据
	Content        *[]TemplateView `json:"content,omitempty" xml:"content"`
	HttpStatusCode int             `json:"-"`
}

func (o ListTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListTemplatesResponse", string(data)}, " ")
}
