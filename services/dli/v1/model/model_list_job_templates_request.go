package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobTemplatesRequest Request Object
type ListJobTemplatesRequest struct {

	// 类型。
	Type *string `json:"type,omitempty"`

	// 模板名过滤关键字，模糊匹配，获取模板名含有该关键字的所有模板。
	Keyword *string `json:"keyword,omitempty"`

	// 每页显示的最大结果行数，范围: [1, 100]。默认值为：50。
	PageSize *int32 `json:"page-size,omitempty"`

	// 当前页码，默认为第一页。
	CurrentPage *int32 `json:"current-page,omitempty"`
}

func (o ListJobTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListJobTemplatesRequest", string(data)}, " ")
}
