package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListProductsRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 查询返回记录的数量限制
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录
	Offset *int32 `json:"offset,omitempty"`

	// 显示的条目排列顺序，使用:分隔参考值和顺序， 如sort=created_at%3Adesc表示根据创建时间逆序排列
	Sort *string `json:"sort,omitempty"`
}

func (o ListProductsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRequest struct{}"
	}

	return strings.Join([]string{"ListProductsRequest", string(data)}, " ")
}
