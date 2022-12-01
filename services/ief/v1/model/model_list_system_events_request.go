package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSystemEventsRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 系统订阅名称
	Name *string `json:"name,omitempty"`

	// 查询返回记录的数量限制
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListSystemEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemEventsRequest struct{}"
	}

	return strings.Join([]string{"ListSystemEventsRequest", string(data)}, " ")
}
