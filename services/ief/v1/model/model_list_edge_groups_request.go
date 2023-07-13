package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEdgeGroupsRequest Request Object
type ListEdgeGroupsRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 边缘节点组名称
	Name *string `json:"name,omitempty"`

	// 查询返回记录的数量限制
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录
	Offset *int32 `json:"offset,omitempty"`

	// 显示的条目排列顺序，使用:分隔参考值和顺序，如sort=created_at%3Adesc表示根据创建时间逆序排列
	Sort *string `json:"sort,omitempty"`
}

func (o ListEdgeGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListEdgeGroupsRequest", string(data)}, " ")
}
