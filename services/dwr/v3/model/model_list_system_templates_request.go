package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSystemTemplatesRequest struct {

	// 模板前缀。
	Prefix *string `json:"prefix,omitempty"`

	// 算子模板的分类。
	Category *string `json:"category,omitempty"`

	// 查询的起始位置。start大于等于1，最大1000，不设置则取默认值1。
	Offset *int32 `json:"offset,omitempty"`

	// 请求返回的最大记录条数。limit取值最小1，最大1000，不设置则取默认值10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSystemTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListSystemTemplatesRequest", string(data)}, " ")
}
