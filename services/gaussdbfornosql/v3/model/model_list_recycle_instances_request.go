package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecycleInstancesRequest Request Object
type ListRecycleInstancesRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询个数上限值。取值范围：1~100。不传该参数时，默认查询前100条信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRecycleInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecycleInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListRecycleInstancesRequest", string(data)}, " ")
}
