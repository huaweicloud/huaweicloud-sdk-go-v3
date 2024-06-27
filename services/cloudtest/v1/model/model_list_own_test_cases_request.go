package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOwnTestCasesRequest Request Object
type ListOwnTestCasesRequest struct {

	// 当前页数
	PageNo int32 `json:"page_no"`

	// 每页条数
	PageSize int32 `json:"page_size"`

	// 排序字段
	SortField *string `json:"sort_field,omitempty"`

	// 排序方式
	SortType *string `json:"sort_type,omitempty"`

	// 关键字查询，用例名或编号
	Keyword *string `json:"keyword,omitempty"`
}

func (o ListOwnTestCasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOwnTestCasesRequest struct{}"
	}

	return strings.Join([]string{"ListOwnTestCasesRequest", string(data)}, " ")
}
