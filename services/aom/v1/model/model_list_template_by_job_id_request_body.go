package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTemplateByJobIdRequestBody 根据作业id分页查询方案集合
type ListTemplateByJobIdRequestBody struct {

	// 方案名称
	Name *string `json:"name,omitempty"`

	// page_num为正整数
	PageNum *int32 `json:"page_num,omitempty"`

	// 每页显示的条数
	PageSize *int32 `json:"page_size,omitempty"`

	// 需要排序的字段(默认为更新时间),支持字段有name，create_time，update_time。
	OrderByColumn string `json:"order_by_column"`

	// 排序规则(默认降序) 传入升序或降序，升序：ASC，降序：DESC。
	SortOrder *string `json:"sort_order,omitempty"`
}

func (o ListTemplateByJobIdRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateByJobIdRequestBody struct{}"
	}

	return strings.Join([]string{"ListTemplateByJobIdRequestBody", string(data)}, " ")
}
