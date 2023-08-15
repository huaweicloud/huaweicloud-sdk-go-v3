package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionsRequest Request Object
type ListConnectionsRequest struct {

	// 偏移量，表示从此偏移量开始查询，偏移量不能小于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，不能小于1或大于1000
	Limit *int32 `json:"limit,omitempty"`

	// 指定查询排序
	Sort *string `json:"sort,omitempty"`

	// 指定查询的目标连接名称，精准匹配
	Name *string `json:"name,omitempty"`

	// 指定查询的目标连接名称，模糊匹配
	FuzzyName *string `json:"fuzzy_name,omitempty"`

	// 指定查询的目标连接，队列的实例id
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o ListConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListConnectionsRequest", string(data)}, " ")
}
