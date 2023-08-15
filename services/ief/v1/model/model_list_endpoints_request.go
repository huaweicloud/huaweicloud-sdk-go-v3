package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointsRequest Request Object
type ListEndpointsRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 端点名称
	Name *string `json:"name,omitempty"`

	// 端点类型 枚举值： - dis - servicebus - apigw
	Type *string `json:"type,omitempty"`

	// 端点是否共享
	IsShared *string `json:"is_shared,omitempty"`

	// 查询返回记录的数量限制
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListEndpointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointsRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointsRequest", string(data)}, " ")
}
