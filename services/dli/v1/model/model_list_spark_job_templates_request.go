package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSparkJobTemplatesRequest Request Object
type ListSparkJobTemplatesRequest struct {

	// 类型，需要且必须传入spark作为值。
	Type *string `json:"type,omitempty"`

	// 模板名过滤关键字，模糊匹配，获取模板名含有该关键字的所有模板。
	Keyword *string `json:"keyword,omitempty"`

	// 每页显示的最大结果行数，范围: [1, 1000]。默认值为：50。
	PageSize *int32 `json:"page-size,omitempty"`

	// 当前页码，默认为第一页。
	CurrentPage *int32 `json:"current-page,omitempty"`
}

func (o ListSparkJobTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSparkJobTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListSparkJobTemplatesRequest", string(data)}, " ")
}
