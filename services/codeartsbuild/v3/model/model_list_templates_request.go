package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTemplatesRequest Request Object
type ListTemplatesRequest struct {

	// 检索的模板的名字模糊查询
	Name *string `json:"name,omitempty"`

	// 分页页码， 表示从此页开始查询
	Page *string `json:"page,omitempty"`

	// 每页显示的条目数量，page_size小于等于100
	PageSize *string `json:"page_size,omitempty"`
}

func (o ListTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListTemplatesRequest", string(data)}, " ")
}
