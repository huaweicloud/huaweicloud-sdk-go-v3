package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListValueListRequest Request Object
type ListValueListRequest struct {

	// 分页查询时，返回第几页数据。默认值为1，表示返回第1页数据。
	Page *int32 `json:"page,omitempty"`

	// 分页查询时，每页包含多少条结果。范围1-100，默认值为10，表示每页包含10条结果。
	Pagesize *int32 `json:"pagesize,omitempty"`

	// 引用表名称
	Name *string `json:"name,omitempty"`
}

func (o ListValueListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListValueListRequest struct{}"
	}

	return strings.Join([]string{"ListValueListRequest", string(data)}, " ")
}
