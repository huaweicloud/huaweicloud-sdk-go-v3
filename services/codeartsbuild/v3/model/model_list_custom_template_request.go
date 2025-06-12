package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomTemplateRequest Request Object
type ListCustomTemplateRequest struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 过滤条件
	Filter *string `json:"filter,omitempty"`

	// 分页页码，表示从此页开始查询，page大于等于1
	Page *int32 `json:"page,omitempty"`

	// 每页显示的条目数量，page_size小于等于100
	PageSize *int32 `json:"page_size,omitempty"`

	// 构建状态过滤条件
	Tags *string `json:"tags,omitempty"`
}

func (o ListCustomTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListCustomTemplateRequest", string(data)}, " ")
}
