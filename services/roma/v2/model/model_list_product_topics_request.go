package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListProductTopicsRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 产品ID

	ProductId int32 `json:"product_id"`
	// 每页显示条目数量，最大数量999，超过999后只返回999

	Limit *int32 `json:"limit,omitempty"`
	// 偏移量，表示从此偏移量开始查询， offset大于等于0

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListProductTopicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductTopicsRequest struct{}"
	}

	return strings.Join([]string{"ListProductTopicsRequest", string(data)}, " ")
}
