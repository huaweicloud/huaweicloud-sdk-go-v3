package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExchangesRequest Request Object
type ListExchangesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 所属Vhost名称
	Vhost string `json:"vhost"`

	// 分页查询偏移量，表示从此偏移量开始查询，offset大于等于0，默认从0开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 分页查询单页数量，取值范围0~50，默认查询10条。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListExchangesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExchangesRequest struct{}"
	}

	return strings.Join([]string{"ListExchangesRequest", string(data)}, " ")
}
